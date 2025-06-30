package main

// importing multiple packages
import (
	"fmt"
	"math"
)

// global scope
const s string = "I am a constant"

func main() {
	fmt.Println(s)

	// constants
	// have arbitary precision
	// logical for math expressions and logics
	const n = 4000000000000
	const d = 3e20 / n
	fmt.Println(n)
	// math.Sin expects a float64
	// since this context needs a float 64, go auto gives d as a float64 according to the context
	fmt.Println(math.Sin(d))
}

// const x = 42
// var a int = x         // becomes int
// var b float64 = x     // becomes float64
// var c complex128 = x  // becomes complex128
// In Go, a number like 42 or 0.5 doesn't have a type until it's used somewhere that requires one — like calling a function (math.Sin) that demands float64.
