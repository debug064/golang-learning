package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
)

type MonotonicStack struct {
	items []int
}

func (s *MonotonicStack) Push(n int) {

	for len(s.items) > 0 && s.items[len(s.items)-1] > n {
		s.Pop()
	}
	s.items = append(s.items, n)
}

func (s *MonotonicStack) Pop() {
	s.items = s.items[:max(0, len(s.items)-1)]
}

// Size returns the number of items in the stack
func (s *MonotonicStack) Size() int {
	return len(s.items)

} // IsEmpty returns true if the stack is empty
func (s *MonotonicStack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *MonotonicStack) Print() {
	fmt.Printf("[%v]\n", s.items)
}

func test_MonotonicStack() {
	fmt.Println(math.MaxInt32) // 2147483647
	fmt.Println(math.MaxInt64) // 9223372036854775807

	fmt.Println(runtime.GOOS, runtime.GOARCH)
	fmt.Println(strconv.IntSize)

	s := MonotonicStack{}
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Size())
	s.Pop()
	s.Print()
	nums := []int{1, 2, 3, 4, 5}
	//nums := []int{5, 4, 3, 2, 1}
	for _, n := range nums {
		s.Push(n)
	}
	s.Print()
	s = MonotonicStack{}
	nums = []int{1, 2, 2, 1, 5}
	//nums := []int{5, 4, 3, 2, 1}
	for _, n := range nums {
		s.Push(n)
	}
	s.Print()
}
