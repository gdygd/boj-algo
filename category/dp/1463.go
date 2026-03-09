package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var x int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &x)
	var dp [1000001]int
	for i := 0; i < 1000001; i++ {
		dp[i] = 99999999
	}
	dp[x] = 0

	for n := x; n >= 1; n-- {

		if n >= 3 && n%3 == 0 {
			dp[n/3] = min(dp[n]+1, dp[n/3])
		}
		if n >= 2 && n%2 == 0 {
			dp[n/2] = min(dp[n]+1, dp[n/2])
		}
		if n-1 > 0 {
			dp[n-1] = min(dp[n]+1, dp[n-1])
		}
	}
	fmt.Printf("%d\n", dp[1])
}
