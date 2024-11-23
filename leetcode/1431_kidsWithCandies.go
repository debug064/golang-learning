package main

import "fmt"

// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/?envType=study-plan-v2&envId=leetcode-75

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var res = make([]bool, len(candies))
	m := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > m {
			m = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= m {
			res[i] = true
		}
	}
	return res
}

func kidsWithCandiesLeet(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	greatest := 0
	for _, val := range candies {
		if val > greatest {
			greatest = val
		}
	}

	for idx, val := range candies {
		if val+extraCandies >= greatest {
			result[idx] = true
		} else {
			result[idx] = false
		}
	}

	return result
}

func test_kidsWithCandies() {
	var nums = []int{2, 3, 5, 1, 3}
	fmt.Println(nums)
	res := kidsWithCandies(nums, 3)
	fmt.Println(res)

	nums = []int{4, 2, 1, 1, 2}
	fmt.Println(nums)
	res = kidsWithCandies(nums, 1)
	fmt.Println(res)

	nums = []int{12, 1, 12}
	fmt.Println(nums)
	res = kidsWithCandiesLeet(nums, 10)
	fmt.Println(res)
}
