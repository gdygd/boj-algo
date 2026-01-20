package main

import "fmt"

func main() {
	var n int
	var m int

	fmt.Scanf("%d %d\n", &n, &m)
	var cards []int = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &cards[i])
	}

	var gab int = 300000
	var rst int
	for a := 0; a < n-2; a++ {
		for b := a + 1; b < n-1; b++ {
			for c := b + 2; c < n; c++ {

				a1 := cards[a]
				b1 := cards[b]
				c1 := cards[c]

				sum := a1 + b1 + c1
				if m-sum >= 0 {
					if gab > (m - sum) {
						gab = m - sum
						rst = sum
					}
				}
			}
		}
	}

	fmt.Printf("%d\n", rst)
}
