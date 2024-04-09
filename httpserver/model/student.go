package model

import "fmt"

type Student struct {
	FirstName string
	LastName  string
	ID        int64
}

func (s Student) String() string {
	return fmt.Sprintf("FirstName: %s, LastName: %s, ID: %d", s.FirstName, s.LastName, s.ID)
}
