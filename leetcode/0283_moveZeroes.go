package main

import "fmt"

// https://leetcode.com/problems/move-zeroes/description/?envType=study-plan-v2&envId=leetcode-75

func moveZeroes(nums []int) {
	i := 0
	j := 0
	for j < len(nums) {
		if nums[j] == 0 {
			j++
			continue
		}
		if i != j {
			nums[i], nums[j] = nums[j], nums[i]
		}
		i++
		j++
	}
	//return nums
}

func test_moveZeroes() {
	var nums = []int{0, 1, 0, 3, 12}
	fmt.Println(nums)
	moveZeroes(nums)
	fmt.Println(nums)
}
