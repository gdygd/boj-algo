package main

import "fmt"

type Que struct {
	arr []int
}

func NewQue() *Que {
	return &Que{
		arr: []int{},
	}
}

func (q *Que) Push(n int) {
	q.arr = append(q.arr, n)
}

func (q *Que) Pop() (int, bool) {
	if len(q.arr) == 0 {
		return 0, false
	}

	n := q.arr[0]
	q.arr = q.arr[1:]

	return n, true
}

func (q *Que) Size() int {
	return len(q.arr)
}

var graph [1001][1001]int

/*
	1 : 2, 3, 4
	2 : 1, 4
	3 : 1, 4
	4 : 1, 2

*/

var visit [1001]int

func dfs(v int) {
	fmt.Printf("%d ", v)
	for v2 := 0; v2 < 1001; v2++ {
		if graph[v][v2] == 1 {
			if visit[v2] == 0 {
				visit[v2] = 1

				dfs(v2)
			}
		}
	}
}

func bfs(start int) {
	q := NewQue()
	q.Push(start)

	visit = [1001]int{}
	visit[start] = 1

	for q.Size() > 0 {
		v, ok := q.Pop()
		if !ok {
			break
		}

		fmt.Printf("%d ", v)

		// search
		for v2 := 1; v2 < 1001; v2++ {
			if graph[v][v2] == 1 {
				if visit[v2] == 0 {
					visit[v2] = 1
					q.Push(v2)
				}
			}
		}
	}
}

func main() {
	var vcnt int
	var lcnt int
	var start int

	fmt.Scanf("%d %d %d", &vcnt, &lcnt, &start)

	for i := 0; i < lcnt; i++ {
		var v1, v2 int
		fmt.Scanf("%d %d", &v1, &v2)
		graph[v1][v2] = 1
		graph[v2][v1] = 1
	}

	visit[start] = 1
	dfs(start)
	fmt.Println()

	bfs(start)
	fmt.Println()
}
