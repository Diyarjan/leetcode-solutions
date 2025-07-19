// 217. Contains Duplicate
package main

import "fmt"

func containsDuplicate(nums []int) bool {
	a := make(map[int]bool)

	for _, val := range nums {
		if _, ok := a[val]; ok {
			return true
		}
		a[val] = true
	}

	return false
}

func main() {

	fmt.Println(containsDuplicate([]int{1, 2, 3, 3, 4}))

}
