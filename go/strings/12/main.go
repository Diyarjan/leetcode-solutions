// 9. Palindrome Number
package main

import "fmt"

func isPalindrome(x int) bool {
	// s := strconv.Itoa(x)
	// for i := 0; i <= int(len(s)/2); i++ {
	// 	if s[i] != s[len(s)-1-i] {
	// 		return false
	// 	}
	// }
	// return true

	n := x
	m := 0
	for n > 0 {
		m = m*10 + (n % 10)
		n = int(n / 10)
	}
	return m == x && x >= 0
}

func main() {

	fmt.Println(isPalindrome(1131))

}
