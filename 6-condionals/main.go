package main

import "fmt"

// THERE IS NO TERNERAY OPERATOR IN GO!!!

func main() {

	num := 2

	if num%2 == 0 {
		n := "even"
	}

	fmt.Println(n)

}
