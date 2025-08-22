// No. 3195. Find the Minimum Area to Cover All Ones I
package main

func minimumArea(grid [][]int) int {

	n := len(grid)
	m := len(grid[0])

	minRow, minCol := -1, -1
	maxRow, maxCol := -1, -1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				if minRow == -1 {
					minRow = i
				}
				if minCol == -1 || j < minCol {
					minCol = j
				}
				maxRow = i
				if j > maxCol {
					maxCol = j
				}
			}
		}
	}

	if minRow == -1 {
		return 0
	}

	return (maxRow - minRow + 1) * (maxCol - minCol + 1)
}
