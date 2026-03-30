package main

import "fmt"

type seq struct {
	len int
	sum int
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var numlist []int = make([]int, n)
	var dp []seq = make([]seq, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numlist[i])
		dp[i].sum = numlist[i]
	}

	for i := 0; i < n; i++ {
		num1 := numlist[i]

		for j := i - 1; j >= 0; j-- {
			num2 := numlist[j]

			if num1 > num2 {
				if dp[j].sum+numlist[i] > dp[i].sum {
					dp[i].sum = dp[j].sum + numlist[i]
				}
			}
		}
	}

	var max int = 0
	for i := 0; i < n; i++ {
		if max < dp[i].sum {
			max = dp[i].sum
		}
	}

	fmt.Printf("%d\n", max)
}
