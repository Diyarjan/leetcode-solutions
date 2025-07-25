// 3487. Maximum Unique Subarray Sum After Deletion
package main

import "fmt"

func maxSum(nums []int) int {
	a := make(map[int]bool)
	sum := 0
	max := nums[0]

	for _, val := range nums {
		if val < 1 {
			if val > max {
				max = val
			}
			continue
		}
		if _, ok := a[val]; !ok {
			sum += val
			a[val] = true
		}
	}

	if sum == 0 {
		return max
	}

	return sum
}

func main() {
	test1 := []int{1, 2, 3, 4, 5}
	test2 := []int{1, 2, -1, -2, 1, 0, -1}

	fmt.Println(maxSum(test1))
	fmt.Println(maxSum(test2))
}
