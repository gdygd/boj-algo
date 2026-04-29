package main

import (
	"fmt"
)

func main() {
	var n, s int
	var numlist []int

	fmt.Scanf("%d %d", &n, &s)
	numlist = make([]int, n)

	for i := 0; i < len(numlist); i++ {
		fmt.Scanf("%d", &numlist[i])
	}

	right := 0
	left := 0

	minCnt := 100001
	sum := numlist[0]
	for right <= left && left < n {
		if sum >= s {

			if minCnt > left-right+1 {
				minCnt = left - right + 1
			}
			sum -= numlist[right]
			right++

		} else {
			left++
			if left < n {
				sum += numlist[left]
			}

		}
	}
	if minCnt == 100001 {
		fmt.Println("0")
	} else {
		fmt.Println(minCnt)
	}
}
