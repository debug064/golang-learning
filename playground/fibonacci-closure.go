package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f0 := 0
	f1 := 1
	i := -1
	var maxInt int
	switch strconv.IntSize {
	case 64:
		maxInt = math.MaxInt64
	default:
		maxInt = math.MaxInt32
	}

	return func() int {
		i += 1
		switch i {
		case 0:
			return 0
		case 1:
			return 1
		}

		if f1 > maxInt-f0 {
			return f1
		}
		val := f0 + f1
		f0 = f1
		f1 = val
		return val
	}
}

func test_fibonacci() {
	fmt.Println(math.MaxInt32) // 2147483647
	fmt.Println(math.MaxInt64) // 9223372036854775807

	fmt.Println(runtime.GOOS, runtime.GOARCH)
	fmt.Println(strconv.IntSize)

	f := fibonacci()
	for i := 0; i < 11; i++ {
		fmt.Println(f())
	}
}
