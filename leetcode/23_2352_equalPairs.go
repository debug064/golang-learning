package main

import "fmt"

// Hash Map / Set
// https://leetcode.com/problems/equal-row-and-column-pairs/description/?envType=study-plan-v2&envId=leetcode-75

func equalPairs(grid [][]int) int {
	hash := func(ints []int) uint64 {
		var hash uint64 = 0
		for _, v := range ints {
			hash = (hash << 5) ^ (hash >> 27) ^ uint64(v)
		}
		return hash
	}

	m := make(map[uint64]int)
	for r := 0; r < len(grid); r++ {
		h := hash(grid[r])
		m[h] += 1
	}

	count := 0
	for c := 0; c < len(grid); c++ {
		a := make([]int, len(grid))
		for r := 0; r < len(grid); r++ {
			a[r] = grid[r][c]
		}
		h := hash(a)
		val, ok := m[h]
		if ok {
			count += val
		}
	}

	return count
}

func testEqualPairs() {
	testCases := []struct {
		nums     [][]int
		expected int
	}{
		{nums: [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}, expected: 1},
		{nums: [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}, expected: 3},
	}

	for _, tc := range testCases {
		fmt.Printf("%v \n", tc.nums)
		res := equalPairs(tc.nums)
		fmt.Printf("expected %d = %d\n", tc.expected, res)
	}
}
