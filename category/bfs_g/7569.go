package main

import "fmt"

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

var visit [][]int

func bfstmt2(q *Que1, boxes [][]int, m, n, h int) int {
	var dir [4][2]int = [4][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	var max int = 0

	for q.Size() > 0 {
		pos, ok := q.Pop()

		if !ok {
			break
		}

		// 좌우앞뒤
		for _, d := range dir {
			ni, nj := pos[0]+d[0], pos[1]+d[1]

			if ni < 0 || ni >= (n*h) || nj < 0 || nj >= m {
				continue
			}

			// 층 경계 체크 (현재층 다음층 넘어가는 경계)
			if ni/n != pos[0]/n {
				continue
			}

			if boxes[ni][nj] != 0 {
				continue
			}

			if visit[ni][nj] == 1 {
				continue
			}
			// cur
			curnext := boxes[pos[0]][pos[1]] + 1
			boxes[ni][nj] = curnext

			q.Push([2]int{ni, nj})
			visit[ni][nj] = 1

			if max < curnext {
				max = curnext
			}
		}

		curnext := boxes[pos[0]][pos[1]] + 1

		i := pos[0]
		j := pos[1]

		// down
		if i-n >= 0 {
			if boxes[i-n][j] == 0 && visit[i-n][j] == 0 {

				visit[i-n][j] = 1
				boxes[i-n][j] = curnext
				q.Push([2]int{i - n, j})

				if max < curnext {
					max = curnext
				}
			}
		}

		// up

		if i+n < n*h {
			if boxes[i+n][j] == 0 && visit[i+n][j] == 0 {
				visit[i+n][j] = 1
				boxes[i+n][j] = curnext
				q.Push([2]int{i + n, j})

				if max < curnext {
					max = curnext
				}
			}
		}

	}

	if max > 1 {
		max = max - 1
	}
	for i := 0; i < n*h; i++ {
		isbreak := false
		for j := 0; j < m; j++ {
			if boxes[i][j] == 0 {
				isbreak = true
				max = -1
				break

			}
		}
		if isbreak {
			break
		}
	}

	return max
}

func main() {
	var m, n, h int
	q := NewQue1()
	fmt.Scanf("%d %d %d\n", &m, &n, &h)

	var boxes [][]int = make([][]int, n*h)
	visit = make([][]int, n*h)

	for i := 0; i < n*h; i++ {

		boxes[i] = make([]int, m)
		visit[i] = make([]int, m)
		for j := 0; j < m; j++ {

			fmt.Scanf("%d", &boxes[i][j])

			// cur
			if boxes[i][j] == 1 {
				visit[i][j] = 1
				q.Push([2]int{i, j})
			}

		}
	}

	rst := bfstmt2(q, boxes, m, n, h)
	fmt.Printf("%d\n", rst)
}
