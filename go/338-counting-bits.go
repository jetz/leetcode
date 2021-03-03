package main

import "fmt"

func countBits(num int) []int {
	ans := make([]int, num+1)
	for i := 0; i <= num; i++ {
		ans[i] = ans[i/2] + (i & 1)
	}
	return ans
}

func main() {
	num := 2
	fmt.Println(countBits(num))
	num = 5
	fmt.Println(countBits(num))
}
