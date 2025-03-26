package main

import (
	"fmt"
	"strings"
	"time"
)

// Graphs - DFS
// https://leetcode.com/problems/keys-and-rooms/description/?envType=study-plan-v2&envId=leetcode-75

func canVisitAllRooms(rooms [][]int) bool {
	v := make([]bool, len(rooms))
	q := []int{}
	q = append(q, rooms[0]...)
	v[0] = true
	count := 1

	for len(q) > 0 {
		room := q[0]
		if !v[room] {
			v[room] = true
			count++
			q = append(q, rooms[room]...)
		}
		q = q[1:]
	}
	return count == len(rooms)
}

func test_canVisitAllRooms() {

	// timer := func(name string) func() {
	// 	start := time.Now()
	// 	return func() {
	// 		fmt.Printf("%s took %v\n", name, time.Since(start))
	// 	}
	// }
	testCases := []struct {
		rooms    [][]int
		expected bool
	}{
		{rooms: [][]int{{1}, {2}, {3}, {}}, expected: true},
		{rooms: [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, expected: false},
	}

	for _, tc := range testCases {
		//fmt.Printf("[%v] \n", tc.n)
		start := time.Now()
		res := canVisitAllRooms(tc.rooms)
		fmt.Printf("It took %v\n", time.Since(start))
		fmt.Printf("expected [%v] = %v\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
}
