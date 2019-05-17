package main

import (
	"fmt"
)

type Animal interface {
	Bite()
}

type Dog struct {
	name string
}

func (dog Dog) Bite() {
	fmt.Println(dog.name + "は噛みます。")
}

func New(name string) *Dog {
	return &Dog{name}
}

func Eat(animal Animal) {
	animal.Bite()
	fmt.Println("飲み込みました。")
}

func main() {
	dog := New("チロ")
	Eat(dog)
}
