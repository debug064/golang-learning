package main

import (
	"fmt"
	"strings"
)

// Monotonic Stack
// https://leetcode.com/problems/online-stock-span/description/?envType=study-plan-v2&envId=leetcode-75
type Pair struct {
	Price int
	Span  int
}

type StockSpanner struct {
	mstack []Pair
	// or like this stack [][2]int
}

func Constructor() StockSpanner {
	s := StockSpanner{}
	return s
}

func (ss *StockSpanner) Next(price int) int {
	res := 1
	for len(ss.mstack) > 0 && price >= ss.mstack[len(ss.mstack)-1].Price {
		res += ss.mstack[len(ss.mstack)-1].Span
		ss.mstack = ss.mstack[:len(ss.mstack)-1]
	}
	ss.mstack = append(ss.mstack, Pair{price, res})
	return res
}

func testStockSpanner() {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{nums: []int{28, 14, 28, 35, 46, 53, 66, 80, 87, 88}, expected: []int{1, 1, 3, 4, 5, 6, 7, 8, 9, 10}},
		{nums: []int{100, 80, 60, 70, 60, 75, 85}, expected: []int{1, 1, 1, 2, 1, 4, 6}},
	}

	for _, tc := range testCases {
		obj := Constructor()
		fmt.Printf("%v \n", tc.nums)
		for i, price := range tc.nums {
			res := obj.Next(price)
			fmt.Printf("expected %d = %d\n", tc.expected[i], res)
		}
	}
	fmt.Println(strings.Repeat("*", 25))

}
