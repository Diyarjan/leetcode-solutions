// 13. Roman to Integer
package main

import "fmt"

func romanToInt(s string) int {
	a := make([]int, 0, len(s))
	
	for _, val := range s {
		switch val {
		case 'M':
			a = append(a, 1000)
		case 'D':
			a = append(a, 500)
		case 'C':
			a = append(a, 100)
		case 'L':
			a = append(a, 50)
		case 'X':
			a = append(a, 10)
		case 'V':
			a = append(a, 5)
		case 'I':
			a = append(a, 1)
		}
	}

	a = append(a, 0)
	sum := 0
	i := 0
	for i < len(a)-1 {
		if a[i] < a[i+1] {
			sum += (a[i+1] - a[i])
			i += 2
		} else {
			sum += a[i]
			i += 1
		}
	}
	return sum
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}