package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	var dp [1000001]int
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	for i := 4; i <= n; i++ {
		dp[i] = (dp[i-2] + dp[i-1]) % 15746
	}

	fmt.Printf("%d\n", dp[n])
}
