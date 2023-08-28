package main

import (
	"fmt"

	"github.com/adzsx/challenges/leetcode/problems"
)

var (
	nums   = []int{2, 7, 11, 15}
	target = 9
)

func main() {
	result := problems.TwoSum(nums, target)
	fmt.Println(result)
}
