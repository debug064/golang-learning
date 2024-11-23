package main

import "fmt"

// https://leetcode.com/problems/can-place-flowers/?envType=study-plan-v2&envId=leetcode-75

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n <= 0 {
		return true
	}
	spots := 2
	for _, val := range flowerbed {
		if val == 0 {
			spots -= 1
			if spots == 0 {
				n--
				if n == 0 {
					break
				}
				spots = 2
			}
		} else {
			spots = 3
		}
	}
	return n == 0 || (n == 1 && spots == 1)
}

func canPlaceFlowers2(flowerbed []int, n int) bool {

	spots := 2
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		if flowerbed[i] == 0 {
			spots -= 1
			if spots == 0 {
				n--
				spots = 2
			}
		} else {
			spots = 3
		}
	}

	return n == 0 || (n == 1 && spots == 1)
}

type TestData struct {
	nPlace    int
	flowerbed []int
	expected  bool
}

func test_canPlaceFlowers() {
	var tests = []TestData{
		{nPlace: 1, flowerbed: []int{1, 0, 0, 0, 1}, expected: true},
		{nPlace: 4, flowerbed: []int{0, 0, 0, 0, 0}, expected: false},
	}

	for _, val := range tests {
		fmt.Println(val)
		res := canPlaceFlowers(val.flowerbed, val.nPlace)
		fmt.Printf("expected %t = %t\n", val.expected, res)
	}

	for _, val := range tests {
		fmt.Println(val)
		res := canPlaceFlowers2(val.flowerbed, val.nPlace)
		fmt.Printf("expected %t = %t\n", val.expected, res)
	}

	var nums = []int{1, 0, 0, 0, 1}
	fmt.Println(nums)
	res := canPlaceFlowers(nums, 1)
	fmt.Println("expected true =", res)

	nums = []int{0, 0, 0, 0, 0}
	fmt.Println(nums)
	res = canPlaceFlowers(nums, 4)
	fmt.Println("expected false =", res)

	nums = []int{0, 0, 0, 0, 0, 1, 0, 0}
	fmt.Println(nums)
	res = canPlaceFlowers2(nums, 0)
	fmt.Println("expected true =", res)

	nums = []int{1, 0, 0, 0, 1}
	fmt.Println(nums)
	res = canPlaceFlowers2(nums, 2)
	fmt.Println("expected false =", res)

	nums = []int{1, 0, 0, 0, 0, 0, 1}
	fmt.Println(nums)
	res = canPlaceFlowers(nums, 2)
	fmt.Println("expected true =", res)

	nums = []int{0, 0, 1, 0, 1}
	fmt.Println(nums)
	res = canPlaceFlowers(nums, 1)
	fmt.Println("expected true =", res)

	nums = []int{1, 0, 1, 0, 0}
	fmt.Println(nums)
	res = canPlaceFlowers(nums, 1)
	fmt.Println("expected true =", res)
}
