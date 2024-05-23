package main

import (
	"fmt"
	"math/rand"
)

func repeatFuncGen[T any](done <-chan bool, fn func() T) (res <-chan T) {
	streamChan := make(chan T)

	go func() {
		defer close(streamChan)

		for {
			select {
			case <-done:
				return
			default:
				streamChan <- fn()
			}
		}

	}()

	return streamChan
}

func randFetcher() int {
	return rand.Intn(10)
}

func sq(data <-chan int) (res <-chan int) {

	out := make(chan int)

	go func() {
		for val := range data {
			select {
			case out <- val * val:
			}
		}
		close(out)

	}()

	return out

}

func main() {

	/// Generator code start
	done := make(chan bool)

	stream := repeatFuncGen(done, randFetcher)

	pipe2 := sq(stream)

	for val := range pipe2 {
		switch val {
		case 4:
			close(done)
		default:
			fmt.Println(val)
		}
	}

	// Generator end

}
