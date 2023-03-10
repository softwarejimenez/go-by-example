package main

import "fmt"

func main() {
	// Declare
	s := make([]string, 3)
	fmt.Println("empty", s)

	// Get and set
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get", s[2])

	// Append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append", s)

	// copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy", c)

	// slide
	l := s[2:5]
	fmt.Println("slide", l)

	l = s[:5]
	fmt.Println("slide 2", l)

	l = s[2:]
	fmt.Println("slide 3", l)

	// Declare and initialize
	t := []string{"g", "h", "i"}
	fmt.Println("declaration", t)

	// Multidimensional
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)
}
