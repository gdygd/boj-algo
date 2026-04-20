package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func sub(str string) int {
	return 0
}

func fplus(str string) int {
	var strnum string
	var sum int
	for _, ch := range str {
		if unicode.IsDigit(ch) {
			strnum += string(ch)
		} else {
			n, _ := strconv.Atoi(strnum)
			sum += n
			strnum = ""
		}
	}
	n, _ := strconv.Atoi(strnum)
	sum += n

	return sum
}

func main() {
	var expression string
	fmt.Scanf("%s", &expression)

	plus := strings.Split(expression, "-")

	var token []int = []int{}
	for _, exp := range plus {
		var isnum bool = true
		for _, ch := range exp {
			if !unicode.IsDigit(ch) {
				isnum = false
				break
			}
		}
		if isnum {
			n, _ := strconv.Atoi(exp)
			token = append(token, n)
		} else {
			n := fplus(exp)
			token = append(token, n)
		}
	}
	var rst int = token[0]
	for i := 1; i < len(token); i++ {
		rst -= token[i]
	}
	fmt.Printf("%v \n", rst)
}
