package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	maps [][]int
	n    int
	dp   [][]int
	dir  [3][2]int = [3][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
	}
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)

	maps = make([][]int, n)
	dp = make([][]int, n)
	for i := 0; i < n; i++ {
		maps[i] = make([]int, i+1)
		dp[i] = make([]int, i+1)

		for j := 0; j < i+1; j++ {
			fmt.Fscan(reader, &maps[i][j])
			dp[i][j] = -1
		}
	}
	dp[0][0] = maps[0][0]

	for r := 1; r < n; r++ {
		for c := 0; c <= r; c++ {
			if c == 0 {
				// 0
				dp[r][0] = dp[r-1][0] + maps[r][c]
			} else if c == r {
				// c-1
				dp[r][c] = dp[r-1][c-1] + maps[r][c]
			} else {
				// c-1, c+1
				dp[r][c] = max(dp[r-1][c-1], dp[r-1][c]) + maps[r][c]
			}
		}
	}
	var rst int = -1
	for _, d := range dp[n-1] {
		if rst < d {
			rst = d
		}
	}
	fmt.Printf("%d\n", rst)
}
