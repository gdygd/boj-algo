package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	var coins []int = []int{}

	for i := 0; i < n; i++ {
		var coin int
		fmt.Scanf("%d", &coin)
		coins = append(coins, coin)
	}

	var cnt int = 0

	for k > 0 {

		coin := 0
		for i := 0; i < n; i++ {
			if k >= coins[i] {
				coin = coins[i]
			} else {
				break
			}
		}

		cnt += (k / coin)
		k = k % coin

	}
	fmt.Printf("%d\n", cnt)
}
