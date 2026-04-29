package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var sollist []int

	fmt.Scanln(&n)
	sollist = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &sollist[i])
	}

	p1 := 0
	p2 := len(sollist) - 1

	var minV int = math.MaxInt

	var v1, v2 int
	for p1 < p2 {
		s1 := sollist[p1]
		s2 := sollist[p2]

		x := (s1 + s2)
		if x < 0 {
			x *= -1
		}

		if minV > x {
			minV = x
			v1, v2 = s1, s2
		}

		if s1+s2 > 0 {
			p2--
		} else if s1+s2 < 0 {
			p1++
		} else {
			break
		}

	}
	fmt.Println(v1, v2)
}
