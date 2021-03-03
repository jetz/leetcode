package main

import "fmt"

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for k := i; k <= j; k++ {
		sum += this.nums[k]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
func main() {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(na.SumRange(0, 2))
	fmt.Println(na.SumRange(2, 5))
	fmt.Println(na.SumRange(0, 5))
}
