package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, x int
	var numlist []int
	fmt.Scanln(&n)
	numlist = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numlist[i])
	}
	fmt.Scanln(&x)

	sort.Slice(numlist, func(i, j int) bool {
		return numlist[i] < numlist[j]
	})

	var cnt int = 0

	p1, p2 := 0, len(numlist)-1
	for {
		if p1 >= p2 {
			break
		}
		v1 := numlist[p1]
		v2 := numlist[p2]

		/*
			v1 + v2 > x
				v2 --
			else
				v1 ++
		*/
		if v1+v2 == x {
			cnt++
			p1++
		} else if v1+v2 > x {
			p2--
		} else {
			p1++
		}
	}

	fmt.Printf("%d\n", cnt)
}
