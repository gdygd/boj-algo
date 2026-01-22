package main

import (
	"fmt"
	"strconv"
)

func main() {
	var strN string
	fmt.Scanf("%s\n", &strN)
	inputN, _ := strconv.Atoi(strN)

	offset := len(strN) * 9

	num, _ := strconv.Atoi(strN)
	var rst int
	for i := num - offset; i < num+offset; i++ {
		strNum := strconv.Itoa(i)

		var sum int = i
		for _, digit := range strNum {
			n, _ := strconv.Atoi(string(digit))

			sum += n
		}

		if sum == inputN {
			rst = i
			break
		}
	}
	fmt.Println(rst)
}
