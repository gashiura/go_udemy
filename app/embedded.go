package main

import (
	"fmt"
)

type Vertex struct{
	X,Y int
}

func (v Vertex) Area() int {
	return v.X * v.Y
}

type Vertex3D struct{
	Vertex
	z int
}

func (v Vertex3D) Area() int {
	return v.X * v.Y * v.z
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	v := New(2, 4, 6)
	fmt.Println(v.Area())
}
