package main

import "fmt"

type Que2 struct {
	arr [][2]int
}

func NewQue2() *Que2 {
	return &Que2{
		arr: [][2]int{},
	}
}

func (q *Que2) Push(pos [2]int) {
	q.arr = append(q.arr, pos)
}

func (q *Que2) Pop() ([2]int, bool) {
	if len(q.arr) < 1 {
		return [2]int{}, false
	}

	f := q.arr[0]
	q.arr = q.arr[1:]
	return f, true
}

func (q *Que2) Size() int {
	return len(q.arr)
}

var visit [50][50]int

func bfs3(q *Que2, maps [50][50]int, n, m int) {
	var dir [4][2]int = [4][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for q.Size() > 0 {
		pos, ok := q.Pop()

		if !ok {
			break
		}

		for _, d := range dir {
			ni, nj := pos[0]+d[0], pos[1]+d[1]

			if ni < 0 || ni >= n || nj < 0 || nj >= m {
				continue
			}

			if visit[ni][nj] == 1 {
				continue
			}

			if maps[ni][nj] == 0 {
				continue
			}

			visit[ni][nj] = 1
			q.Push([2]int{ni, nj})
		}
	}
}

func main() {
	q := NewQue2()
	var m, n, cnt int
	var caseCnt int

	fmt.Scanln(&caseCnt)

	// var rst []int = []int{}

	var rst []int = []int{}
	for c := 0; c < caseCnt; c++ {
		fmt.Scanf("%d %d %d\n", &m, &n, &cnt)

		var maps [50][50]int
		visit = [50][50]int{}
		var warmcnt int

		for k := 0; k < cnt; k++ {

			var mi, ni int
			fmt.Scanf("%d %d\n", &mi, &ni)
			maps[ni][mi] = 1
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if visit[i][j] == 0 {
					if maps[i][j] == 0 {
						continue
					}
					visit[i][j] = 1

					q.Push([2]int{i, j})

					bfs3(q, maps, n, m)
					warmcnt++

				}
			}
		}

		rst = append(rst, warmcnt)

	}

	for _, r := range rst {
		fmt.Printf("%d\n", r)
	}
}
