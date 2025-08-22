// 2264. Largest 3-Same-Digit Number in String

package main

import (
	"fmt"
	"strings"
)

func largestGoodInteger(num string) string {
	max := ""

	for i := 0; i < len(num)-2; i++ {

		if num[i] == num[i+1] && num[i] == num[i+2] {
			curr := num[i : i+3]
			if curr > max {
				max = curr
			}
			i += 2
		}
	}

	return max
}

func largestGoodInteger_v2(num string) string {
	patterns := []string{"999", "888", "777", "666", "555", "444", "333", "222", "111", "000"}

	for _, p := range patterns {
		if strings.Contains(num, p) {
			return p
		}
	}
	return ""
}

func main() {
	fmt.Println(largestGoodInteger("2300019"))
	fmt.Println(largestGoodInteger_v2("6777133339"))
}
