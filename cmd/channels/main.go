package main

import "fmt"

func print(num int, c chan int) {

	// fmt.Printf("Num is %v \n", num)
	c <- num
}

// func main() {
// 	c := make(chan int)

// 	go print(100, c)

// 	fmt.Printf("data read from channel %v", <-c)

// }

func main() {
	c := make(chan int)
	c2 := make(chan int)

	go print(100, c)
	go print(200, c2)

	select {
	case dataFromC := <-c:
		fmt.Printf("data read from channel C %v", dataFromC)
	case dataFromC2 := <-c2:
		fmt.Printf("data read from channel C2 %v", dataFromC2)

	}

}
