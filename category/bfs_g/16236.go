package main

import (
	"fmt"
	"sort"
)

type Shark struct {
	level   int
	pos     [2]int
	path    int
	fishlv  int
	fishlv2 int

	cnt int
}

type Que2 struct {
	arr []Shark
}

func NewQue2() *Que2 {
	return &Que2{
		arr: []Shark{},
	}
}

func (q *Que2) Push(pos Shark) {
	q.arr = append(q.arr, pos)
}

func (q *Que2) Front() (Shark, bool) {
	if len(q.arr) == 0 {
		return Shark{}, false
	}
	return q.arr[0], true
}

func (q *Que2) Pop() (Shark, bool) {
	if len(q.arr) == 0 {
		return Shark{}, false
	}

	f := q.arr[0]
	q.arr = q.arr[1:]
	return f, true
}

func (q *Que2) Size() int {
	return len(q.arr)
}

func (q *Que2) Clear() {
	for q.Size() > 0 {
		q.Pop()
	}
}

func bfs2(q *Que2, maps [][]int, n int) {
	// var pathcnt int

	var dir [4][2]int = [4][2]int{
		{-1, 0}, {1, 0}, {0, 1}, {0, -1},
	}

	var feedlist []Shark = []Shark{}
	s, ok := q.Front()
	if !ok {
		return
	}

	var pathcnt int = 0
	var visit [20][20]int = [20][20]int{}
	var feedcnt int = 0

	for {
		feedlist = []Shark{}
		visit = [20][20]int{}
		for q.Size() > 0 {

			fish, ok := q.Pop()

			if !ok {
				break
			}

			if fish.fishlv > 0 && s.level > fish.fishlv {

				feedlist = append(feedlist, fish)
				maps[fish.pos[0]][fish.pos[1]] = 0

				q.Clear()
				visit = [20][20]int{}
				// 최초 shark 위치
				q.Push(s)

				continue
			}

			for _, d := range dir {
				ni, nj := fish.pos[0]+d[0], fish.pos[1]+d[1]

				if ni < 0 || ni >= n || nj < 0 || nj >= n {
					continue
				}

				if visit[ni][nj] == 1 {
					continue
				}

				flv := maps[ni][nj]
				if s.level < flv {
					continue
				}

				visit[ni][nj] = 1

				q.Push(Shark{
					level:  s.level,
					pos:    [2]int{ni, nj},
					fishlv: flv,
					path:   fish.path + 1,
				})
			}

		}

		if len(feedlist) == 0 {
			break
		}

		sort.Slice(feedlist, func(i, j int) bool {
			if feedlist[i].path != feedlist[j].path {
				return feedlist[i].path < feedlist[j].path
			}

			if feedlist[i].pos[0] != feedlist[j].pos[0] {
				return feedlist[i].pos[0] < feedlist[j].pos[0]
			}

			return feedlist[i].pos[1] < feedlist[j].pos[1]
		})

		// recover
		for i := 1; i < len(feedlist); i++ {
			flv := feedlist[i].fishlv
			pos := feedlist[i].pos
			maps[pos[0]][pos[1]] = flv
		}

		sh := feedlist[0]
		feedcnt++
		pathcnt += sh.path

		// 먹은fish
		maps[sh.pos[0]][sh.pos[1]] = 0

		// fmt.Printf("[feed] s[lv:%d, path:%d] f[lv:%d, pos:%v] \n", sh.level, sh.path, sh.fishlv, sh.pos)

		if feedcnt >= sh.level {
			sh.level += 1
			feedcnt = 0
		}

		// init shark pos
		q.Push(Shark{
			level: sh.level,
			pos:   sh.pos,
			path:  0,
		})
		s = Shark{
			level: sh.level,
			pos:   sh.pos,
			path:  0,
		}
	}

	fmt.Printf("%d \n", pathcnt)
}

func main() {
	var n int
	var maps [][]int
	q := NewQue2()

	fmt.Scanf("%d\n", &n)
	maps = make([][]int, n)

	for i := 0; i < n; i++ {
		maps[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &maps[i][j])

			if maps[i][j] == 9 {
				q.Push(Shark{level: 2, pos: [2]int{i, j}})
				maps[i][j] = 0
			}
		}
	}

	bfs2(q, maps, n)
}
