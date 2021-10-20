package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(double("Baju")) // BBaajjuu
}

func double(str string) string {
	var res []string
	strs := strings.Split(str, "")
	for _, v := range strs {
		res = append(res, strings.Repeat(v, 2))
	}
	return strings.Join(res, "")
}
