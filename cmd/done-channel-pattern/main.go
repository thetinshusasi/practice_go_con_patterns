package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {

	for {
		select {
		// all you need a value so if you try to read from a close channel  , you will get a default nil value of its type
		case <-done:
			return
		default:
			fmt.Println("Doing work")
		}
	}
}

func main() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 5)
	// Reading from a closed channel succeeds immediately, returning the zero value of the underlying type.
	//  The optional second return value is true if the value received was delivered by a successful send operation to the channel,
	//  or false if it was a zero value generated because the channel is closed and empty.
	close(done)
}
