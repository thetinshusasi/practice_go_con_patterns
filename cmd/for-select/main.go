package main

import "fmt"

func main() {

	charChannel := make(chan string, 3)

	chars := []string{
		"a",
		"b",
		"c",
	}

	// For - Select channel pattern
	for _, val := range chars {
		select {
		case charChannel <- val:
		}
	}

	close(charChannel)
	for data := range charChannel {
		fmt.Println(data)
	}
}
