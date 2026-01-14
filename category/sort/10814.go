package main

import (
	"fmt"
	"sort"
)

/*
3
21 Junkyu
21 Dohyun
20 Sunyoung
*/
type Person struct {
	Age  int
	Name string
	Reg  int
}

func main() {
	var members []Person = []Person{
		// {21, "Dohyun", 0},
		// {21, "Junkyu", 1},

		// {20, "Sunyoung", 2},
	}

	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var mm Person
		fmt.Scanf("%d %s", &mm.Age, &mm.Name)
		mm.Reg = i
		members = append(members, mm)
	}

	sort.Slice(members, func(i, j int) bool {
		if members[i].Age == members[j].Age {
			return members[i].Reg < members[j].Reg
		} else {
			return members[i].Age < members[j].Age
		}
	})

	for _, m := range members {
		fmt.Printf("%d, %s \n", m.Age, m.Name)
	}
}
