package main

import "fmt"

func countSquares(matrix [][]int) int {
	count := 0
	n := len(matrix)
	m := len(matrix[0])

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				}

				count += dp[i][j]
			}
		}
	}

	return count
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	} else {
		if b < c {
			return b
		}
		return c
	}
}

func main() {

	fmt.Println("Test 1 : ", countSquares([][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}))
	fmt.Println("Test 1 : ", countSquares([][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}))

}
