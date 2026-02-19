package main

import "fmt"

func main() {
	var n int
	var dp [1001]int

	fmt.Scanln(&n)

	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	for i := 4; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 10007
	}

	fmt.Printf("%d\n", dp[n])
}
