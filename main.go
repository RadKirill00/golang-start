package main

import "fmt"

func totalprice(initPrce int) func(int) int {
	res := initPrce
	return func(i int) int {
		res += i
		return res
	}
}

func main() {

	result := totalprice(1)
	fmt.Println(result(1))
	fmt.Println(result(1))
	fmt.Println(result(1))
	fmt.Println(result(1))

}
