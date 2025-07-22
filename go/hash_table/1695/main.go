// 1695. Maximum Erasure Value

// Input: nums = [5,2,1,2,5,2,1,2,5]
// Output: 8
// Explanation: The optimal subarray here is [5,2,1] or [1,2,5].

package main

import "fmt"

func maximumUniqueSubarray(nums []int) int {

	max, s, i := 0, 0, 0
	a := make(map[int]bool)
	n := len(nums)

	for j := 0; j < n; j++ {
		for a[nums[j]] {
			a[nums[i]] = false
			s -= nums[i]
			i++
		}

		a[nums[j]] = true
		s += nums[j]
		if s > max {
			max = s
		}
	}

	return max
}

func main() {

	fmt.Println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5}))

}
