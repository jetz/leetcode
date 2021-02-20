package main

import "fmt"

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		ok := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				ok = false
				break
			}
		}
		if ok {
			return i
		}
	}
	return -1
}

func main() {
	haystack, needle := "hello", "lo"
	fmt.Println(strStr(haystack, needle))
	haystack, needle = "aaaaa", "bba"
	fmt.Println(strStr(haystack, needle))
}
