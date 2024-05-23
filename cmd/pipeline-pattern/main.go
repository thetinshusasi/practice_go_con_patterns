package main

import (
	"fmt"
)

func mul(nums []int) (data <-chan int) {

	out := make(chan int)

	go func() {
		for _, val := range nums {
			select {
			case out <- val * 2:
			}
		}
		close(out)

	}()

	return out

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
	nums := []int{
		1, 2, 3, 4, 5, 6,
	}

	pipe1Chan := mul(nums)

	pipe2Chan := sq(pipe1Chan)

	for data := range pipe2Chan {
		fmt.Println(data)
	}
}
