package main

import (
	"fmt"
	"math"
	"sort"
)

// Method 1
func subsets1(nums []int) [][]int {
	sets := [][]int{{}}
	for _, num := range nums {
		tempSets := [][]int{}
		for _, s := range sets {
			dst := make([]int, len(s))
			copy(dst, s)
			tempSets = append(tempSets, append(dst, num))
		}
		sets = append(sets, tempSets...)
	}

	for _, s := range sets {
		sort.Ints(s)
	}

	return sets
}

// Method 2

// Method 3
func subsets3(nums []int) [][]int {
	sets := [][]int{}
	n := len(nums)
	for i := 0; i < int(math.Pow(2, float64(n))); i++ {
		format := "%." + fmt.Sprintf("%d", n) + "b"
		binStr := fmt.Sprintf(format, i)
		tempSets := []int{}
		for i := 0; i < len(binStr); i++ {
			if binStr[i] == '1' {
				tempSets = append(tempSets, nums[n-1-i])
			}
		}
		sets = append(sets, tempSets)
	}
	for _, s := range sets {
		sort.Ints(s)
	}
	return sets
}

func main() {
	subsets := subsets1
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
	nums = []int{0}
	fmt.Println(subsets(nums))
	nums = []int{9, 0, 3, 5, 7}
	fmt.Println(subsets(nums))

	subsets = subsets3
	fmt.Println(subsets(nums))
}
