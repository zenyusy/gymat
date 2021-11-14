package mathutil

import (
	"runtime"
)

var CNUM int = 2

func init() {
	CNUM = runtime.NumCPU()
}

func DivUp(a, b int) int {
	r := a / b
	if r*b != a {
		return r + 1
	}
	return r
}

func CalcStep(n int) int {
	return DivUp(n, CNUM)
}
