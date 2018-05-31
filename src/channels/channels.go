package main

import "fmt"

func process(messages chan string, quit chan struct{}) {
	defer close(quit)

	// Loop through your messages
	for m := range messages {
		// print the message with a for loop using range
		fmt.Println(m)
	}
}

func main() {
	// declare the messages channel of type string and capacity of 5
	messages := make(chan string, 5)

	// declare a signal channel
	quit := make(chan struct{})

	// launch process in a goroutine
	go process(messages, quit)

	// declare 5 fruits in a []string
	fruits := []string{
		"Apple",
		"Orange",
		"Banana",
		"Pear",
		"Passion Fruit",
	}

	// loop through the fruits and send them to the messages channel
	for _, f := range fruits {
		messages <- f
	}

	// close the messages channel
	close(messages)

	// wait for everything to finish
	i, ok := <-quit
	if !ok {
		fmt.Println("done")
	} else {
		fmt.Printf("read %d from channel\n", i)
	}

}
