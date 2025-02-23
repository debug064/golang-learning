package main

import "fmt"

// Queue
// https://leetcode.com/problems/number-of-recent-calls/description/?envType=study-plan-v2&envId=leetcode-75

type RecentCounter struct {
	ticks []int
	begin int
}

func Constructor1() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	// this.ticks = append(this.ticks, t)
	// for this.ticks[0] < t-3000 {
	// 	this.ticks = this.ticks[1:]
	// }
	// return len(this.ticks)
	this.ticks = append(this.ticks, t)
	for this.begin < len(this.ticks) {
		if this.ticks[this.begin] >= t-3000 {
			break
		}
		this.begin++
	}
	return len(this.ticks) - this.begin
}

func test_recentCalls() {
	testCases := []struct {
		calls    []int
		expected []int
	}{
		{calls: []int{1, 100, 3001, 3002}, expected: []int{1, 2, 3, 3}},
	}

	obj := Constructor1()
	for _, tc := range testCases {
		for i := 0; i < len(tc.calls); i++ {
			fmt.Printf("%d \n", tc.calls[i])
			res := obj.Ping(tc.calls[i])
			fmt.Printf("expected %d = %d\n", tc.expected[i], res)
		}
	}
}
