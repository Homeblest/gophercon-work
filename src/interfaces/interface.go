package main

import "fmt"

type user struct {
	ID    int
	First string
	Last  string
}

// Task:
// Satisfy the stringer interface (https://golang.org/pkg/fmt/#Stringer)
// so that the user struct will print
// user <ID> is <First> <Last>
//
// example:
//      user 1 is Rob Pike

func (u user) String() string {
	return fmt.Sprintf("user %d is %s %s", u.ID, u.First, u.Last)
}

func main() {
	u := user{ID: 1, First: "Rob", Last: "Pike"}
	fmt.Println(u)
}
