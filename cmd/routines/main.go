package main

import (
	"fmt"
	"time"
)

func print(num int) {

	fmt.Printf("Num is %v \n", num)
}

func main() {

	go print(1)
	go print(2)
	go print(3)

	time.Sleep(time.Millisecond * 300)

	fmt.Println(" main finished")
}
