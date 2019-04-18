package domain

import "fmt"

type User struct {
	ID   int
	Name string
}

func(u User) String() string {
	return fmt.Sprintf("{\"ID\": %v ,\"Name\": \"%v\"}", u.ID, u.Name)
}
