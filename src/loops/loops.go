package main

import "fmt"

func main() {
	fruits := [4]string{"Banana", "Orange", "Pineapple", "Strawberry"}

	// Use a 'classic' for loop  to print out each fruit in the array, and the
	// corresponding index.
	fmt.Println("Classic for loop:")
	for i := 0; i < len(fruits); i++ {
		fmt.Println(i, fruits[i])
	}

	fmt.Println("Range for loop:")
	// Use the range keyword to loop over the same array of fruits, again printing
	// out the fruit and the corresponding index.
	for index, name := range fruits {
		fmt.Println(index, name)
	}

	fmt.Println("Range for loop skipping even ones")
	// Using the keyword `continue`, skip every other fruit (even ones)
	for index, name := range fruits {
		if index%2 == 0 {
			continue
		}
		println(index, name)
	}
}
