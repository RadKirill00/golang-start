package main

import "fmt"

func something(callback func(int, int) int, s string) int {
	res := callback(5, 5)
	fmt.Println(s)
	return res
}

func main() {
	callbackFunc := func(n int, q int) int { return n + q }

	res := something(callbackFunc, "Hello, Kirill Radchenko")

	fmt.Println(res)
}
