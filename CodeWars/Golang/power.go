package main

import "fmt"

func power(number, power int) int {

	sum := 1

	for i := 1; i <= power; i++ {
		sum *= number
	}

	return sum

}

func main() {
	fmt.Println(power(3, 4))
}
