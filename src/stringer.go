package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Age int
}

func (animal Animal) String() string {
	return fmt.Sprintf("My Name is %v.", animal.Name)
}

func main() {
	animal := Animal{"Mike", 18}
	fmt.Println(animal)
}
