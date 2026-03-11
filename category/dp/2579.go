package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &n)
	var stairs [301]int
	var dp [301]int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &stairs[i])
	}

	dp[0] = stairs[0]
	dp[1] = stairs[1]
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]) + stairs[i]
	}
	fmt.Printf("%d\n", dp[n-1])
}
