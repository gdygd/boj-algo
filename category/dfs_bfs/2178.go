package main

import (
	"fmt"
)

type Que1 struct {
	arr [][2]int
}

func NewQue1() *Que1 {
	return &Que1{
		arr: [][2]int{},
	}
}

func (q *Que1) Push(pos [2]int) {
	q.arr = append(q.arr, pos)
}

func (q *Que1) Pop() ([2]int, bool) {
	if len(q.arr) == 0 {
		return [2]int{}, false
	}

	f := q.arr[0]
	q.arr = q.arr[1:]
	return f, true
}

func (q *Que1) Size() int {
	return len(q.arr)
}

var (
	visitCnt [101][101]int
	visit    [101][101]int
)

func bfs(q *Que1, maps [][]rune, n, m int) {
	var dir [4][2]int = [4][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for q.Size() > 0 {

		pos, ok := q.Pop()

		if !ok {
			break
		}

		if pos[0] == n && pos[1] == m {
			// fmt.Printf("target : %d, %d cnt : %d\n", pos[0], pos[1], visitCnt[pos[0]][pos[1]])
			fmt.Printf("%d\n", visitCnt[pos[0]][pos[1]])
		}

		for _, d := range dir {
			ni, nj := pos[0]+d[0], pos[1]+d[1]

			if ni < 0 || ni > n || nj < 0 || nj > m {
				continue
			}

			if maps[ni][nj] != '1' {
				continue
			}
			if visit[ni][nj] == 1 {
				continue
			}

			visit[ni][nj] = 1

			curcnt := visitCnt[pos[0]][pos[1]]
			if curcnt > visitCnt[ni][nj] {
				visitCnt[ni][nj] = curcnt + 1

				// fmt.Printf("\t Push[%d, %d] -> [%d, %d] cnt : %d->%d\n", pos[0], pos[1], ni, nj, curcnt, curcnt+1)
				q.Push([2]int{ni, nj})
			}

		}
	}
}

func main() {
	q := NewQue1()
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)
	var maps [][]rune = make([][]rune, n)

	for i := 0; i < n; i++ {
		maps[i] = make([]rune, m)
		var input string
		fmt.Scanln(&input)
		maps[i] = []rune(input)
	}

	visit[0][0] = 1
	visitCnt[0][0] = 1

	q.Push([2]int{0, 0})
	bfs(q, maps, n-1, m-1)
}
