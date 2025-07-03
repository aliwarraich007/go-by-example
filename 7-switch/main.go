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

	

	switch time.Now().Weekday() {
		case time.Saturday
	}

}
