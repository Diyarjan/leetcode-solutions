// 1717. Maximum Score From Removing Substrings
package main

import "fmt"

func maximumGain(s string, x int, y int) int {
	var t string

	total, sum := 0, 0

	if x > y {
		t, sum = removeSubString(s, 'a', 'b', x)
		total += sum
		_, sum = removeSubString(t, 'b', 'a', y)
		total += sum
	} else {
		t, sum = removeSubString(s, 'b', 'a', y)
		total += sum
		_, sum = removeSubString(t, 'a', 'b', x)
		total += sum
	}

	return total
}

func removeSubString(s string, first, second byte, score int) (string, int) {
	result := make([]byte, 0)
	sum := 0

	for i := 0; i < len(s); i++ {
		if len(result) > 0 && result[len(result)-1] == first && s[i] == second {
			result = result[:len(result)-1]
			sum += score
		} else {
			result = append(result, s[i])
		}
	}

	return string(result), sum
}

func main() {
	test1 := "cdbcbbaaabab"
	test2 := "aabbaaxybbaabb"

	fmt.Println(maximumGain(test1, 4, 5))
	fmt.Println(maximumGain(test2, 5, 4))
}