package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	writer = bufio.NewWriter(os.Stdout)
	reader = bufio.NewReader(os.Stdin)
)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()
	var n int
	var numlist []int = []int{
		// 2, 4, -10, 4, -9,
	}

	fmt.Scanf("%d", &n)
	numlist = make([]int, n)
	var compNum []int = []int{}
	var checkN map[int]int = make(map[int]int)

	// input
	for i := 0; i < n; i++ {
		var num int
		scanf("%d", &num)
		if checkN[num] < 1 {
			compNum = append(compNum, num)
		}
		// numlist = append(numlist, num)
		numlist[i] = num
		checkN[num]++

	}

	var rank map[int]int = make(map[int]int)

	for i := 0; i < len(compNum); i++ {
		for j := 0; j < len(compNum); j++ {
			if i == j {
				continue
			}

			if compNum[i] > compNum[j] {
				rank[compNum[i]]++
			}
		}
	}

	for _, num := range numlist {
		if _, ok := rank[num]; !ok {
			printf("0 ")
		} else {
			r := rank[num]
			printf("%d ", r)
		}
	}
}
