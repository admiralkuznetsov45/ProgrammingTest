package main

import "fmt"

func paperwork(m, n int) int {
	hasil := m * n

	if m < 0 || n < 0 {
		return 0
	}

	return hasil
}

func main() {
	fmt.Println(paperwork(5, 5))
}
