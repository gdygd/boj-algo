package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	maps [][]int
	dp   [][]int
	dir  [4][2]int = [4][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	M, N int
)

func dfs(r, c int) int {
	if r == M-1 && c == N-1 {
		return 1
	}

	if dp[r][c] != -1 {
		return dp[r][c]
	}

	dp[r][c] = 0

	for _, d := range dir {
		nr, nc := d[0]+r, d[1]+c

		if nr < 0 || nr >= M || nc < 0 || nc >= N {
			continue
		}
		if maps[r][c] > maps[nr][nc] {
			dp[r][c] += dfs(nr, nc)
		}

	}
	return dp[r][c]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &M, &N)

	maps = make([][]int, M)
	dp = make([][]int, M)
	for i := 0; i < M; i++ {
		maps[i] = make([]int, N)
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscan(reader, &maps[i][j])
			dp[i][j] = -1
		}

	}
	rst := dfs(0, 0)

	fmt.Printf("%d\n", rst)
}
