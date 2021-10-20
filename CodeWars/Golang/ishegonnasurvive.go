package main

import "fmt"

func Hero(bullets, dragons int) bool {
	// your code
	if bullets == dragons*2 || bullets >= dragons*2 {
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Println(Hero(1, 2))
	fmt.Println(Hero(5, 2))
}
