package main

import "fmt"

var (
	N, K  int
	Items [][2]int
	Dp    [][]int
)

func dfs(idx, rem int) int {
	if idx >= N {
		return 0
	}

	if Dp[idx][rem] != -1 {
		return Dp[idx][rem]
	}

	skip := dfs(idx+1, rem)
	take := 0
	if Items[idx][0] <= rem {
		take = Items[idx][1] + dfs(idx+1, rem-Items[idx][0])
	}

	Dp[idx][rem] = Max(skip, take)
	return Dp[idx][rem]
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Scanf("%d %d", &N, &K)
	Items = make([][2]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d", &Items[i][0], &Items[i][1])
	}

	Dp = make([][]int, N)
	for i := range Dp {
		Dp[i] = make([]int, K+1)
		for j := range Dp[i] {
			Dp[i][j] = -1
		}
	}

	rst := dfs(0, K)
	fmt.Printf("%d\n", rst)
}
