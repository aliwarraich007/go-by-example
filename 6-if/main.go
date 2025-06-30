package main

import "fmt"

// THERE IS NO TERNERAY OPERATOR IN GO!!!

func main() {

	num := 2
	if num%2 == 0 {
		fmt.Println("is even")
	} else {
		fmt.Println("is odd")
	}

	// var declared here will be scoped to full block
	if k := 8; k < 0 {
		fmt.Println("k is negative")
	} else if k > 10 {
		fmt.Println("k is greater than 10")
	} else {
		fmt.Println(k, "K is a postive but less than 10")
	}

}
