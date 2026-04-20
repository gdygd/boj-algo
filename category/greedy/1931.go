package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	times := [][2]int{}

	for i := 0; i < n; i++ {
		var st, end int
		fmt.Scanf("%d %d", &st, &end)
		times = append(times, [2]int{st, end})
	}

	// sort.Slice(times, func(i, j int) bool {
	// 	return times[i][1] < times[j][1]
	// })
	sort.Slice(times, func(i, j int) bool {
		if times[i][1] == times[j][1] {
			return times[i][0] < times[j][0]
		}
		return times[i][1] < times[j][1]
	})

	var rst [][2]int = [][2]int{}
	idx := 0

	for {
		rst = append(rst, times[idx])
		prevEndtm := times[idx][1]
		idx += 1

		var issearch bool = false
		for i := idx; i < len(times); i++ {
			sttm := times[i][0]
			if sttm >= prevEndtm {
				idx = i
				issearch = true
				break
			}
		}
		if !issearch {
			break
		}
	}
	fmt.Printf("%d \n", len(rst))
}
