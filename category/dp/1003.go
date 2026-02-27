package main

import "fmt"

type Fibo struct {
	zero int
	one  int
}

func main() {
	var tcase int

	var rst [][2]int = [][2]int{}

	fmt.Scanln(&tcase)

	for t := 0; t < tcase; t++ {
		var n int
		fmt.Scanln(&n)

		var dp [41]Fibo
		dp[0] = Fibo{1, 0}
		dp[1] = Fibo{0, 1}

		for i := 2; i <= n; i++ {
			dp[i].zero = dp[i-1].zero + dp[i-2].zero
			dp[i].one = dp[i-1].one + dp[i-2].one
		}

		// fmt.Printf("%d, %d\n", dp[n].zero, dp[n].one)
		rst = append(rst, [2]int{dp[n].zero, dp[n].one})
	}

	for _, r := range rst {
		fmt.Println(r[0], r[1])
	}
}
