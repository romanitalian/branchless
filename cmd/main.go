package main

import (
	"fmt"
	"github.com/romanitalian/branchless"
	"math"
)

func main() {
	fmt.Println(math.Min(1, 3))
	fmt.Println(branchless.Min(1, 3))

	fmt.Println(branchless.Abs(-12))
	fmt.Println(branchless.Abs(13))
	fmt.Println(branchless.Abs(0))

	fmt.Println(branchless.MinBranchless(1, 3))
	fmt.Println(branchless.MinBranchless(3, 1))
	fmt.Println(branchless.MinBranchless(2, 49))
	fmt.Println(branchless.MinBranchless(49, 2))
}
