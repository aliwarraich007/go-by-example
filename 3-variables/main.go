package main

import "fmt"

func main() {
	// initializing a variable
	var a = "String here"
	fmt.Println(a)
	// defining multiple variables here
	var b, c int = 1, 2
	fmt.Println(b, c)
	// go will initialize it by default of 0 since it was not init by us.
	// go auto inferes the type from initalization;
	var d int
	fmt.Println(d)
	// short hand to decalre variable with a value;
	e := "inferred auto"

	fmt.Println(e)
}
