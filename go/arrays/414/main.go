// 414. Third Maximum Number
package main

import "fmt"

func thirdMax(nums []int) int {
	min := -1 * (1 << 62)
	max1 := min
	max2 := min
	max3 := min

	for i := 0; i < len(nums); i++ {
		if nums[i] > max1 {
			max3 = max2
			max2 = max1
			max1 = nums[i]
		} else if nums[i] > max2 && nums[i] != max1 {
			max3 = max2
			max2 = nums[i]
		} else if nums[i] > max3 && nums[i] != max1 && nums[i] != max2 {
			max3 = nums[i]
		}
	}
	if max3 != min {
		return max3
	} else {
		return max1
	}
}

func main() {
	fmt.Println(thirdMax([]int{4, 52, 15, 20}))
}