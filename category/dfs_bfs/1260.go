package main

import "fmt"

func main() {
	var graph [][2]int = [][2]int{}
	var vcnt int
	var lcnt int
	var start int

	fmt.Scanf("%d %d %d", &vcnt, &lcnt, &start)

	for i := 0; i < lcnt; i++ {
		var v1, v2 int
		fmt.Scanf("%d %d", &v1, &v2)
		graph = append(graph, [2]int{v1, v2})
	}

	fmt.Printf("graph : %v \n", graph)
}
