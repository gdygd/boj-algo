package main

import "fmt"

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

func bfstmt(q *Que, farm [][]int, m, n int) int {
	var max int = 0
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

			if farm[ni][nj] == -1 || farm[ni][nj] > 0 {
				continue
			}

			next := farm[pos[0]][pos[1]] + 1
			farm[ni][nj] = next

			if max < next {
				max = next
			}

			q.Push([2]int{ni, nj})

		}

	}

	if max > 1 {
		max = max - 1
	}
	for i := 0; i < n; i++ {
		isbreak := false
		for j := 0; j < m; j++ {
			if farm[i][j] == 0 {
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
	q := NewQue()
	var m, n int
	fmt.Scanf("%d %d\n", &m, &n)

	var farm [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		farm[i] = make([]int, m)

		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &farm[i][j])

			if farm[i][j] == 1 {
				q.Push([2]int{i, j})
			}
		}
	}

	rst := bfstmt(q, farm, m, n)
	fmt.Printf("%d\n", rst)
}
