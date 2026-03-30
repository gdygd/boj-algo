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

func pandadfs(i, j, depth int) int {
	if Dp[i][j] > 0 {
		return Dp[i][j]
	}

	/*
		Dp[i][j] = dfs(i, j)
	*/

	for _, d := range Dir {
		ni, nj := i+d[0], j+d[1]
		if ni < 0 || ni >= Sz || nj < 0 || nj >= Sz {
			continue
		}

		if Maps[i][j] < Maps[ni][nj] {
			Dp[i][j] = pandadfs(ni, nj, depth+1) + depth
		}
	}

	return Dp[i][j]
}

func main() {
	fmt.Scanf("%d", &Sz)
	Maps = make([][]int, Sz)
	Dp = make([][]int, Sz)

	for i := 0; i < Sz; i++ {
		Maps[i] = make([]int, Sz)
		Dp[i] = make([]int, Sz)

		for j := 0; j < Sz; j++ {
			fmt.Scanf("%d", &Maps[i][j])
		}
		fmt.Printf("%v \n", Maps[i])
	}

	for i := 0; i < Sz; i++ {
		for j := 0; j < Sz; j++ {
			pandadfs(i, j, 0)
		}
	}

	for _, d := range Dp {
		fmt.Printf("%v \n", d)
	}
}
