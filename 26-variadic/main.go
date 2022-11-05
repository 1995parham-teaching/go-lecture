package main

import "fmt"

func sums(nums ...int) {
	fmt.Printf("these are your numbers: %v\n", nums)

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sums(1, 2, 3)
	sums(2, 3, 4)

	nums := []int{1, 2, 3, 4}
	sums(nums...)
}
