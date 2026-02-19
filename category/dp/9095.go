package main

import "fmt"

var cnt int

func main() {
	var tcase, n int
	fmt.Scanln(&tcase)

	var rst []int = []int{}
	for i := 0; i < tcase; i++ {
		fmt.Scanln(&n)

		var dp [11]int
		dp[1] = 1
		dp[2] = 2
		dp[3] = 4

		for i := 4; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
		}

		rst = append(rst, dp[n])
	}
	for _, r := range rst {
		fmt.Println(r)
	}
}
