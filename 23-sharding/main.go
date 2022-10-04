package main

import (
	"crypto/sha1"
	"fmt"
	"sync"
)

type Key interface {
	comparable
	fmt.Stringer
}

type Shard[K Key, V interface{}] struct {
	sync.RWMutex
	data map[K]V
}

type ShardedMap[K Key, V interface{}] []*Shard[K, V]

func NewShardedMap[K Key, V interface{}](nshards int) ShardedMap[K, V] {
	shards := make([]*Shard[K, V], nshards)

	for i := 0; i < nshards; i++ {
		data := make(map[K]V)
		shards[i] = &Shard[K, V]{data: data}
	}

	return shards
}

func (m ShardedMap[K, V]) GetShardIndex(key K) int {
	checksum := sha1.Sum([]byte(key.String()))
	hash := int(checksum[17])

	return hash % len(m)
}

func (m ShardedMap[K, V]) GetShard(key K) *Shard[K, V] {
	index := m.GetShardIndex(key)

	return m[index]
}

func (m ShardedMap[K, V]) Get(key K) V {
	shard := m.GetShard(key)
	shard.RLock()

	defer shard.RUnlock()

	return shard.data[key]
}

func (m ShardedMap[K, V]) Set(key K, value V) {
	shard := m.GetShard(key)
	shard.Lock()
	defer shard.Unlock()

	shard.data[key] = value
}

func (m ShardedMap[K, V]) Keys() []K {
	keys := make([]K, 0)

	var mutex sync.Mutex

	var wg sync.WaitGroup
	wg.Add(len(m))

	for _, shard := range m {
		go func(s *Shard[K, V]) {
			s.RLock()

			for key := range s.data {
				mutex.Lock()
				keys = append(keys, key)
				mutex.Unlock()
			}

			s.RUnlock()
			wg.Done()
		}(shard)
	}

	wg.Wait()

	return keys
}

type String string

func (s String) String() string {
	return string(s)
}

func main() {
	shardedMap := NewShardedMap[String, int](5)

	shardedMap.Set("Parham", 20)
	shardedMap.Set("Ahmad", 18)
	shardedMap.Set("Ali", 10)

	fmt.Println(shardedMap.Get("Parham"))
	fmt.Println(shardedMap.GetShardIndex("Parham"))

	fmt.Println(shardedMap.Get("Ahmad"))
	fmt.Println(shardedMap.GetShardIndex("Ahmad"))

	fmt.Println(shardedMap.Get("Ali"))
	fmt.Println(shardedMap.GetShardIndex("Ali"))
}
