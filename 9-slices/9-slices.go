package main

import "fmt"

func main() {
	var s []string // should be nill and length should be 0 (which is falsey)
	fmt.Println(s, s == nil, len(s))

	s = make([]string, 3) // if length is not known then we should use make method
	fmt.Println(s, s == nil, len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println(cap(s), "cap after append") // slice cap increased

	// copy a slice

	var c = make([]string, len(s))
	copy(c, s)
	fmt.Println(c, "copied") // with data

	c[1] = "zzz" // will not mutate original array
	fmt.Println(s, c)

	l := s[2:5]
	fmt.Println(l)

	f := s[2:] // slices 0 1 index
	fmt.Println(f)

	k := s[:3] // should return first 3 index
	fmt.Println(k)

}
