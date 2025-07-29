package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

// args with same type can be omited. only give type to last arg for simplicity
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
