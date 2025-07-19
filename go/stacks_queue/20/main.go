// 20. Valid Parentheses
package main

import "fmt"

func isValid(s string) bool {

	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	validMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}

	for _, val := range s {
		if val == '(' || val == '{' || val == '[' {
			stack = append(stack, val)
		} else if val == ')' || val == '}' || val == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != validMap[val] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}


func main() {
	test1 := "{}[]({})"
	fmt.Println(isValid(test1))
}
