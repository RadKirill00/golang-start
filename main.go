package main

import "fmt"

func main() {
	var res int
	a := []int{
		1,
		2,
		3,
		4,
		5,
	}

	for _, j := range a {
		res += j

	}
	fmt.Println(res)
}
