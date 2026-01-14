package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	arr []int
}

func (s *Stack) NewStack() {
	s.arr = []int{}
}

func (s *Stack) Push(x int) {
	s.arr = append(s.arr, x)
}

func (s *Stack) Pop() int {
	if len(s.arr) == 0 {
		return -1
	}
	n := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return n
}

func (s *Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) Empty() int {
	if len(s.arr) == 0 {
		return 1
	}
	return 0
}

func (s *Stack) Top() int {
	if len(s.arr) == 0 {
		return -1
	}
	return s.arr[len(s.arr)-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	st := &Stack{}
	st.NewStack()

	var cmdCnt int
	fmt.Fscanln(reader, &cmdCnt)
	var rst []int = []int{}
	for i := 0; i < cmdCnt; i++ {
		var cmd string
		var x int
		fmt.Fscanf(reader, "%s %d\n", &cmd, &x)

		if cmd == "push" {
			st.Push(x)
		} else if cmd == "pop" {
			rst = append(rst, st.Pop())
		} else if cmd == "size" {
			rst = append(rst, st.Size())
		} else if cmd == "empty" {
			rst = append(rst, st.Empty())
		} else if cmd == "top" {
			rst = append(rst, st.Top())
		}
	}
	for _, r := range rst {
		fmt.Fprintln(writer, r)
	}
}
