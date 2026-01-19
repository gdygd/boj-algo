package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	arr []int
}

func NewQue() *Queue {
	return &Queue{
		arr: []int{},
	}
}

func (q *Queue) Push(x int) {
	q.arr = append(q.arr, x)
}

func (q *Queue) Pop() int {
	if len(q.arr) == 0 {
		return -1
	}
	x := q.arr[0]
	q.arr = q.arr[1:len(q.arr)]
	return x
}

func (q *Queue) Size() int {
	return len(q.arr)
}

func (q *Queue) Empty() int {
	if len(q.arr) == 0 {
		return 1
	}
	return 0
}

func (q *Queue) Front() int {
	if len(q.arr) == 0 {
		return -1
	}
	return q.arr[0]
}

func (q *Queue) Back() int {
	if len(q.arr) == 0 {
		return -1
	}

	return q.arr[len(q.arr)-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	q := NewQue()
	var n int
	fmt.Fscanln(reader, &n)

	var rst []int = []int{}
	for i := 0; i < n; i++ {
		var cmd string
		var x int

		fmt.Fscanln(reader, &cmd, &x)

		if cmd == "push" {
			q.Push(x)
		} else if cmd == "pop" {
			rst = append(rst, q.Pop())
		} else if cmd == "size" {
			rst = append(rst, q.Size())
		} else if cmd == "empty" {
			rst = append(rst, q.Empty())
		} else if cmd == "front" {
			rst = append(rst, q.Front())
		} else if cmd == "back" {
			rst = append(rst, q.Back())
		}
	}

	for _, x := range rst {
		fmt.Fprintln(writer, x)
	}
}
