package main

import "fmt"

type Point struct {
	Y int
	X int
}

func main() {
	oneMao := map[string]Point{
		"a": {1, 2},
	}
	fmt.Println(oneMao)

}
