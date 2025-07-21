// 1957. Delete Characters to Make Fancy String
package main

import "fmt"

func makeFancyString(s string) string {

	result := []byte{s[0]}
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			count = 1
			result = append(result, s[i])

		} else if count < 2 {
			result = append(result, s[i])
			count++
		}
	}

	return string(result)
}

func main() {

	test1 := "leeetcode"
	test2 := "aaabaaaa"
	fmt.Println(makeFancyString(test1))
	fmt.Println(makeFancyString(test2))
}
