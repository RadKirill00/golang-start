package main

import "fmt"

type Point struct {
	X int
	Y int
	T string
}

func (p Point) method() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
}

func main() {
	str := Point{
		X: 2,
		Y: 3,
		T: "Hello, GO",
	}
	str.method()
	fmt.Println(str)
}
