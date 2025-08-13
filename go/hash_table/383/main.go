// 383. Ransom Note

package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {

	if len(ransomNote) > len(magazine) {
		return false
	}

	a := make(map[rune]int)

	for _, val := range magazine {
		a[val]++
	}

	for _, val := range ransomNote {

		if a[val] == 0 {
			return false
		}

		a[val]--
	}

	return true
}

func main() {

	fmt.Println(canConstruct("aca", "aac"))
}
