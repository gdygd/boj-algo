package main

import "fmt"

func main() {
	var n int
	var rgb [][3]int
	var dp [][3]int

	fmt.Scanln(&n)
	rgb = make([][3]int, n)
	dp = make([][3]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scanf("%d", &rgb[i][j])
		}
	}

	const (
		r = 0
		g = 1
		b = 2
	)
	dp[0][r] = rgb[0][r]
	dp[0][g] = rgb[0][g]
	dp[0][b] = rgb[0][b]
	for i := 1; i < len(rgb); i++ {
		dp[i][r] = min(dp[i-1][g], dp[i-1][b]) + rgb[i][r]
		dp[i][g] = min(dp[i-1][r], dp[i-1][b]) + rgb[i][g]
		dp[i][b] = min(dp[i-1][r], dp[i-1][g]) + rgb[i][b]
	}

	fmt.Printf("%d\n", min(min(dp[n-1][r], dp[n-1][g]), dp[n-1][b]))
}
