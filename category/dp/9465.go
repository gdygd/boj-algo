package main

import (
	"bufio"
	"fmt"
	"os"
)

func maxSticker(sticker [2][]int, col int) int {
	var dp [2][]int
	dp[0] = make([]int, col)
	dp[1] = make([]int, col)

	dp[0][0] = sticker[0][0]
	dp[1][0] = sticker[1][0]

	if col > 1 {
		dp[0][1] = dp[1][0] + sticker[0][1]
		dp[1][1] = dp[0][0] + sticker[1][1]

	}

	for i := 2; i < col; i++ {
		dp[0][i] = max(dp[1][i-1], dp[1][i-2]) + sticker[0][i]
		dp[1][i] = max(dp[0][i-1], dp[0][i-2]) + sticker[1][i]
	}

	return max(dp[0][col-1], dp[1][col-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var tCnt int
	fmt.Fscan(reader, &tCnt)
	var rst []int = []int{}

	for t := 0; t < tCnt; t++ {
		var sticker [2][]int

		var col int
		fmt.Fscan(reader, &col)

		for r := 0; r < 2; r++ {
			sticker[r] = make([]int, col)
			for c := 0; c < col; c++ {
				fmt.Fscan(reader, &sticker[r][c])
			}

		}
		maxN := maxSticker(sticker, col)
		rst = append(rst, maxN)
	}
	for _, r := range rst {
		fmt.Printf("%d\n", r)
	}
}
