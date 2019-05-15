package main

import (
	"fmt"
	// "time"
)

func goroutine(s []int, c chan int) {
	sum := 0
	for _, i := range s {
		sum += i
	}
	c <-sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go goroutine(s, c)
	x := <-c
	fmt.Println(x)
}
