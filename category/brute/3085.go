package main

import "fmt"

var dir [4][2]int = [4][2]int{
	{-1, 0}, {0, 1}, {1, 0}, {0, -1},
}

func main() {
	var board [][]rune
	var n int
	var max int = -1
	fmt.Scanf("%d\n", &n)

	board = make([][]rune, n)
	for i := 0; i < n; i++ {
		board[i] = make([]rune, n)
		var line string
		fmt.Scan(&line)
		board[i] = []rune(line)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for _, d := range dir {
				dsti := i + d[0]
				dstj := j + d[1]

				if dsti < 0 || dsti >= n || dstj < 0 || dstj >= n {
					continue
				}

				source := board[i][j]
				dst := board[dsti][dstj]

				board[i][j] = dst
				board[dsti][dstj] = source

				cnt := checkBoard(board, i, j)
				if max < cnt {
					max = cnt
				}
				cnt2 := checkBoard(board, dsti, dstj)
				if max < cnt2 {
					max = cnt2
				}

				board[dsti][dstj] = dst
				board[i][j] = source

			}
		}
	}

	fmt.Printf("%d\n", max)
}

func checkBoard(board [][]rune, i, j int) int {
	var max int = -1
	var count int = 1
	// check colume
	for j := 0; j < len(board)-1; j++ {
		if board[i][j] == board[i][j+1] {
			count++
		} else {
			count = 1
		}

		if max < count {
			max = count
		}
	}

	count = 1
	// check row
	for i := 0; i < len(board)-1; i++ {
		if board[i][j] == board[i+1][j] {
			count++
		} else {
			count = 1
		}

		if max < count {
			max = count
		}
	}

	return max
}
