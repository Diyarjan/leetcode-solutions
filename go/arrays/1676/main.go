// 1672. Richest Customer Wealth
package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	max, sum := 0, 0
	for _, i := range accounts {
		sum = 0
		for _, j := range i {
			sum += j
		}

		if sum > max {
			max = sum
		}
	}
	return max
}

// func maximumWealth(accounts [][]int) int {
//     max, sum := 0, 0

//     for i := 0; i < len(accounts); i++ {
//         sum = 0
//         for j := 0; j < len(accounts[i]); j++ {
//             sum += accounts[i][j]
//         }
//         if sum > max {
//             max = sum
//         }
//     }

//     return max
// }

func main() {
	fmt.Println(maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}))
}
