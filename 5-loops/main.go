package main

import "fmt"

// sadly go has only 1 loop (for) but it's flexible and can be built to mimic any loop
func main() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	} // prints 1 - 5 number -> more like a while loop

	for j := 0; j < 2; j++ {
		fmt.Println(j, "i am a for loop")
	}

	// above loop can be shortended with range
	for k := range 5 {
		fmt.Println(k+1, "k")
	}

	for {
		fmt.Println("I'll run only once")
		break // will become infinte without me
	} // most useless loop tho

	for i := range 10 {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i, "prints only odd numbers")
	}
}
