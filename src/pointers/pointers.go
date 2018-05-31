package main

import "fmt"

func main() {
	x := "George"
	printValue(&x)
	fmt.Println(x)
}

func printValue(s *string) {
	// print the pointer value
	fmt.Println(s)
	// print the string value
	fmt.Println(*s)
	// change the string value
	*s = "Ringo"
}
