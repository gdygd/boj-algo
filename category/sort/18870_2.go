package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d", &n)

	var numlist []int = []int{}
	var compNum []int = []int{}
	var checkN map[int]bool = make(map[int]bool)

	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(reader, &num)
		numlist = append(numlist, num)

		if !checkN[num] {
			compNum = append(compNum, num)
			checkN[num] = true
		}
	}
	sort.Ints(compNum)

	for i := 0; i < len(numlist); i++ {
		rank := sort.SearchInts(compNum, numlist[i])
		fmt.Fprint(writer, rank, " ")
	}
}
