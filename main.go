package main

import "fmt"

func main() {
	pointer(4)
}

func pointer(a int) {

	b := &a
	fmt.Println(a)
	fmt.Println(b)
	*b /= 2
	fmt.Println(a)
}
