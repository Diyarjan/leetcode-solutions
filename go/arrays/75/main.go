//  75. Sort Colors
package main

import "fmt"

func sortColors(nums []int) {
	var red, white, blue int
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			red++
		case 1:
			white++
		default:
			blue++
		}
	}
	for i := 0; i < red; i++ {
		nums[i] = 0
	}
	for i := red; i < red+white; i++ {
		nums[i] = 1
	}
	for i := red + white; i < red+white+blue; i++ {
		nums[i] = 2
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
