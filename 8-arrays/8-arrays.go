package main

import "fmt"

func main() {

	var a [2]int
	// not initalized then go returns default value of given data type
	fmt.Println(a)

	a[0] = 100
	fmt.Println(a)

	// initalization of array
	b := [4]int{1, 2, 4, 5}
	fmt.Println(b)

	// letting compiler calculate length of array, array length will be calculated as 4
	b = [...]int{3, 4, 6, 7}
	fmt.Println("compiler cal: ", b)

	c := [5]int{100, 200, 300, 400, 500}

	c = [...]int{500, 3: 600, 700}

	fmt.Println("indexed composite literal: ", c)

	// 2d array

	var d [2][3]int
	for i := range d { // d = 0 and d = 1 since rows is 2
		fmt.Println(i, "first for") // for each iteration of first for 2nd will run 3 times because columns are 3
		for j := range d[i] {
			fmt.Println(j, "2nd for")
			d[i][j] = i + 1
		}
	}

	fmt.Println(d)

}
