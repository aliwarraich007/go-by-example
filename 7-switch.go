package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	fmt.Println(i)

	// simple switch syntax
	switch i {
	case 1:
		fmt.Println("i is", 1)
	case 2:
		fmt.Println("i is", 2)
	}

	// without the value, it can act as if else
	switch time.Now().Weekday() {
	// using commas to separate cases.
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend")
	default:
		fmt.Println("Its weekday")
	}

	// complex switch with type

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
