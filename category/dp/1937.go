package main

import "fmt"

var (
	Sz   int
	Maps [][]int
	Dp   [][]int
)

type Que struct {
	arr [][2]int
}

func NewQue() *Que {
	return &Que{
		arr: [][2]int{},
	}
}

func (q *Que) Push(pos [2]int) {
	q.arr = append(q.arr, pos)
}

func (q *Que) Pop() ([2]int, bool) {
	if len(q.arr) == 0 {
		return [2]int{}, false
	}

	f := q.arr[0]
	q.arr = q.arr[1:]
	return f, true
}

func (q *Que) Size() int {
	return len(q.arr)
}

var Dir [4][2]int = [4][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pandadfs(i, j int) int {
	if Dp[i][j] > 0 {
		return Dp[i][j]
	}

	Dp[i][j] = 1

	for _, d := range Dir {
		ni, nj := i+d[0], j+d[1]

		if ni < 0 || ni >= Sz || nj < 0 || nj >= Sz {
			continue
		}
		/*
			dp[i][j] = max(dp[i][j], dp[ni][nj]+1)
		*/
		if Maps[i][j] < Maps[ni][nj] {
			Dp[i][j] = max(Dp[i][j], pandadfs(ni, nj)+1)
		}

	}
	return Dp[i][j]
}

func main() {
	fmt.Scanf("%d", &Sz)
	Maps = make([][]int, Sz)
	Dp = make([][]int, Sz)

	var ans int = 0

	for i := 0; i < Sz; i++ {
		Maps[i] = make([]int, Sz)
		Dp[i] = make([]int, Sz)

		for j := 0; j < Sz; j++ {
			fmt.Scanf("%d", &Maps[i][j])
		}
	}

	for i := 0; i < Sz; i++ {
		for j := 0; j < Sz; j++ {
			v := pandadfs(i, j)
			if v > ans {
				ans = v
			}
		}
	}

	fmt.Printf("%d\n", ans)
}
