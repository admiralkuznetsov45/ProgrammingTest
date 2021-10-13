package main

import (
	"fmt"
)

func cockroachSpeed(s float64) float64 {
	var hasil = s * 100000 / 3600

	return hasil
}

func main() {
	fmt.Println(cockroachSpeed(1.08))
}
