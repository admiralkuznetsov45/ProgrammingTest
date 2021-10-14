package main

import "fmt"

func getSum(a, b int) int {

	sum := 0

	//cara pertama yang kompleks
	if a < b {
		for i := a; i <= b; i++ {
			sum += i
		}
	} else if b < a {
		for i := b; i <= a; i++ {
			sum += i
		}
	} else if a == b {
		sum = a
	}
	return sum
}

func main() {
	fmt.Println(getSum(-1, 2))
}
