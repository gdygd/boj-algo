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
	dp[1] = stairs[0] + stairs[1]
	dp[2] = max(stairs[0]+stairs[2], stairs[1]+stairs[2])
	for i := 3; i < n; i++ {
		dp[i] = max(dp[i-3]+stairs[i-1]+stairs[i], dp[i-2]+stairs[i])
	}
	fmt.Printf("%d\n", dp[n-1])
}
