package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var numlist []int
	fmt.Scanf("%d", &n)
	numlist = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numlist[i])
	}

	sort.Slice(numlist, func(i, j int) bool {
		return numlist[i] < numlist[j]
	})

	for _, num := range numlist {
		fmt.Println(num)
	}
}
