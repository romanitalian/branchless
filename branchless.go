package branchless

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func Min(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MinBranchless(a, b int64) int64 {
	n := a - b
	mask := n >> MaxInt
	return ((a + b) - ((n ^ mask) - mask)) / 2
}

func Abs(n int64) int64 {
	mask := n >> 31
	return (n ^ mask) - mask
}
