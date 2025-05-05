package main

import (
	"fmt"
	"strings"
	"time"
)

// Graphs - DFS
// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/?envType=study-plan-v2&envId=leetcode-75

// func findCircleNum(isConnected [][]int) int {
// 	res := 0
// 	n := len(isConnected)
// 	var mark func(i int, j int)
// 	mark = func(i int, j int) {
// 		if isConnected[i][j] == 2 {
// 			return
// 		}
// 		isConnected[i][j] = 2
// 		j++
// 		for ; j < n; j++ {
// 			if isConnected[i][j] == 1 {
// 				isConnected[i][j] = 2
// 				if i != j {
// 					mark(j, 0)
// 				}
// 			} else {
// 				isConnected[i][j] = 2
// 			}
// 		}
// 	}

// 	for i := 0; i < n; i++ {
// 		if isConnected[i][0] != 2 {
// 			mark(i, 0)
// 			res++
// 		}
// 	}
// 	return res
// }

func minReorder2(n int, connections [][]int) int {
	res := 0
	size := len(connections)
	var doit func(parent int, i int) int

	doit = func(parent int, i int) int {
		if i >= size {
			return size
		}
		for i < size {
			if connections[i][0] == parent {
				res += 1
				i = doit(connections[i][1], i+1)
			} else if connections[i][1] == parent {
				i = doit(connections[i][0], i+1)
			} else {
				return i
			}
		}
		return i
	}
	doit(0, 0)
	return res
}

func minReorder3(n int, connections [][]int) int {
	res := 0
	//size := len(connections)

	var doit func(parent int)
	doit = func(parent int) {
		pts := [][]int{}
		for i := 0; i < len(connections); {
			if connections[i][0] == parent || connections[i][1] == parent {
				pts = append(pts, connections[i])
				connections = append(connections[:i], connections[i+1:]...)
				continue
			}
			i++
		}

		for i := 0; i < len(pts); i++ {
			if pts[i][0] == parent {
				res += 1
				doit(pts[i][1])
			} else if pts[i][1] == parent {
				doit(pts[i][0])
			}
		}
	}

	doit(0)
	return res
}

func minReorder4(n int, connections [][]int) int {
	res := 0
	//size := len(connections)

	var doit func(parent int, n int) int
	doit = func(parent int, j int) int {
		for i := j; i < len(connections); {
			if connections[i][0] == parent || connections[i][1] == parent {
				connections[i], connections[j] = connections[j], connections[i]
				if connections[j][0] == parent {
					res += 1
					i = doit(connections[j][1], j+1)
				} else if connections[j][1] == parent {
					i = doit(connections[j][0], j+1)
				}
				j = i
				continue
			}
			i++
		}
		return j

	}

	doit(0, 0)
	return res
}

func minReorder(n int, connections [][]int) int {
	type Edge struct {
		A int
		B int
	}

	mp1 := map[Edge]bool{}

	res := 0
	F := make(map[int][]int, len(connections))
	//B := make(map[int][]int, len(connections))
	for i := 0; i < len(connections); i++ {
		mp1[Edge{A: connections[i][0], B: connections[i][1]}] = true
		//mp1[Edge{A: connections[i][1], B: connections[i][0]}] = false
		F[connections[i][0]] = append(F[connections[i][0]], connections[i][1])
		F[connections[i][1]] = append(F[connections[i][1]], connections[i][0])
		//B[connections[i][1]] = append(B[connections[i][1]], connections[i][0])
	}
	visited := map[int]bool{}

	var doit func(parent int) int
	doit = func(parent int) int {
		if _, ok := visited[parent]; ok {
			return 0
		}
		if vals, ok := F[parent]; ok {
			visited[parent] = true
			for _, val := range vals {
				if _, ok := visited[val]; ok {
					continue
				}
				if _, ok := mp1[Edge{A: parent, B: val}]; ok {
					res += 1
				}
				doit(val)
			}
		}
		return 0
	}

	doit(0)
	return res
}

func test_minReorder() {
	testCases := []struct {
		connection [][]int
		n          int
		expected   int
	}{
		{connection: [][]int{{4, 5}, {0, 1}, {1, 3}, {2, 3}, {4, 0}}, n: 6, expected: 3},
		{connection: [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}, n: 6, expected: 3},
		{connection: [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}, n: 5, expected: 2},
		{connection: [][]int{{1, 0}, {2, 0}}, n: 3, expected: 0},
	}

	// for _, tc := range testCases {
	// 	//fmt.Printf("[%v] \n", tc.n)
	// 	start := time.Now()
	// 	res := minReorder3(tc.n, tc.connection)
	// 	fmt.Printf("It took %v\n", time.Since(start))
	// 	fmt.Printf("expected [%v] = %v\n", tc.expected, res)
	// }
	// fmt.Println(strings.Repeat("-", 25))
	for _, tc := range testCases {
		//fmt.Printf("[%v] \n", tc.n)
		start := time.Now()
		res := minReorder(tc.n, tc.connection)
		fmt.Printf("It took %v\n", time.Since(start))
		fmt.Printf("expected [%v] = %v\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
}
