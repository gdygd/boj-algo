package main

import (
	"fmt"
	"sort"
)

var visit [25][25]int

type Queue struct {
	arr [][2]int
}

func NewQueue() *Queue {
	return &Queue{
		arr: [][2]int{},
	}
}

func (q *Queue) Push(pos [2]int) {
	q.arr = append(q.arr, pos)
}

func (q *Queue) Pop() ([2]int, bool) {
	if len(q.arr) == 0 {
		return [2]int{}, false
	}
	f := q.arr[0]
	q.arr = q.arr[1:]
	return f, true
}

func (q *Queue) Size() int {
	return len(q.arr)
}

func bfs1(q *Queue, n int, maps [][]rune) int {
	var dir [4][2]int = [4][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	var cnt int

	for q.Size() > 0 {
		pos, ok := q.Pop()
		if !ok {
			break
		}

		// check '1'
		for _, d := range dir {
			ni, nj := pos[0]+d[0], pos[1]+d[1]

			if ni < 0 || ni >= n || nj < 0 || nj >= n {
				continue
			}

			if visit[ni][nj] == 0 && maps[ni][nj] == '1' {
				visit[ni][nj] = 1

				q.Push([2]int{ni, nj})
				cnt++
			}
		}

	}
	return cnt
}

func main() {
	q := NewQueue()
	var n int
	var maps [][]rune

	fmt.Scanln(&n)
	maps = make([][]rune, n)

	for i := 0; i < n; i++ {
		maps[i] = make([]rune, n)
		var input string
		fmt.Scanln(&input)
		maps[i] = []rune(input)
	}

	var compCnt int
	var rst []int = []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if maps[i][j] == '1' && visit[i][j] == 0 {
				q.Push([2]int{i, j})
				visit[i][j] = 1
				compCnt++
				cnt := bfs1(q, n, maps)
				rst = append(rst, cnt+1)
			}
		}
	}

	fmt.Println(compCnt)
	sort.Slice(rst, func(i, j int) bool {
		return rst[i] < rst[j]
	})
	for _, r := range rst {
		fmt.Println(r)
	}
}
