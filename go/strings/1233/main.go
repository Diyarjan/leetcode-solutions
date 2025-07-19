// 1233. Remove Sub-Folders from the Filesystem
package main

import (
	"fmt"
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)

	result := []string{}
	prev := folder[0]
	result = append(result, prev)

	for i := 1; i < len(folder); i++ {
		curr := folder[i]
		if !strings.HasPrefix(curr, prev+"/") {
			result = append(result, curr)
			prev = curr
		}
	}

	return result
}

func main() {
	test1 := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	test2 := []string{"/a", "/a/b/c", "/a/b/d"}

	fmt.Println(removeSubfolders(test1))
	fmt.Println(removeSubfolders(test2))
}
