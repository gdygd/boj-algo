package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var dp [1001]int
	var numlist []int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	numlist = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &numlist[i])
	}

	var ans int = 0
	for i := 1; i < n; i++ {
		num1 := numlist[i]
		for j := i - 1; j >= 0; j-- {
			num2 := numlist[j]

			if num2 < num1 {

				if dp[i] > dp[j] {
					continue
				}
				dp[i] = dp[j] + 1

				if ans < dp[i] {
					ans = dp[i]
				}
			}
		}
	}

	fmt.Printf("%d \n", ans+1)
}
