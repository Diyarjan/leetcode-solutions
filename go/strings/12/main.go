// No 12. Integer to Roman.  Given an integer, convert it to a Roman numeral.
package main

import "fmt"

func intToRoman(num int) string {
	a := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	s := [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""

	for i, val := range a {
		for num >= val {
			num -= val
			roman += s[i]
		}
	}
	return roman
}

func main() {

	fmt.Println(intToRoman(199))

}
