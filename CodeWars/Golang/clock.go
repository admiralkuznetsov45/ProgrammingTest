package main

import "fmt"

func Past(h, m, s int) int {
	// your code here
	hh := h * 3600000
	hm := m * 60000
	hs := s * 1000

	test := hh + hm + hs
	return test
}

func main() {
	fmt.Println(Past(0, 1, 1))
}
