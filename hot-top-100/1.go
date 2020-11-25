package main

import (
	"fmt"
)

/**
两数之和
https://leetcode-cn.com/problems/two-sum/
*/

/**
执行用时：4 ms, 在所有 Go 提交中击败了96.57%的用户
内存消耗：3.3 MB, 在所有 Go 提交中击败了69.71%的用户
*/
func SolutionOne(nums []int, target int) []int {
	if len(nums) < 2 {
		panic("num length error")
	}
	numMap := make(map[int]int, len(nums))
	for index, num := range nums {
		if targetIndex, has := numMap[target-num]; has {
			return []int{targetIndex, index}
		} else {
			numMap[num] = index
		}
	}
	return nil
}

/**
执行用时：4 ms, 在所有 Go 提交中击败了96.57%的用户
内存消耗：3 MB, 在所有 Go 提交中击败了73.18%的用户
*/

func SolutionTwo(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	twoIndex := SolutionTwo(nums, target)
	fmt.Println(twoIndex)
}
