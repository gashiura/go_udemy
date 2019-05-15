package main

import (
	"fmt"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v)
	default:
		fmt.Println("I don't Know.")
	}
}

func main() {
	do(2)
	do("aaaa")
	do(true)
}
