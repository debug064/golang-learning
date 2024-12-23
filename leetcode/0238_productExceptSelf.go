package main

import "fmt"

func productExceptSelf(nums []int) []int {
	a := make([]int, len(nums))
	a[0] = 1
	for i := 1; i < len(nums); i++ {
		a[i] = a[i-1] * nums[i-1]
	}
	b := 1
	for i := len(nums) - 1; i >= 0; i-- {
		a[i] *= b
		b *= nums[i]
	}
	return a
}

type TestData238 struct {
	nums     []int
	expected []int
}

func testProductExceptSelf() {
	var tests = []TestData238{
		{nums: []int{1, 2, 3, 4}, expected: []int{24, 12, 8, 6}},
		{nums: []int{1, 2, 3, 4, 5}, expected: []int{120, 60, 40, 30, 24}},
	}

	for _, val := range tests {
		fmt.Println(val)
		res := productExceptSelf(val.nums)
		fmt.Printf("expected %v = %v\n", val.expected, res)
	}
}
