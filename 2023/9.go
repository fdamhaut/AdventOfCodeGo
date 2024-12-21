package main

import (
	"fmt"
)

func nine() {
	lines, _ := toLineArray("9.in")

	a, b := 0, 0

	for _, line := range lines {
		inta := toIntArray(line)
		a += nineA(inta)
		b += nineB(inta)
	}

	fmt.Println(a, b)
}

func nineA(vals []int) int {
	res := 0
	for {
		res += vals[len(vals)-1]
		nv := make([]int, len(vals)-1)
		stop := true
		for n := range len(vals) - 1 {
			nv[n] = vals[n+1] - vals[n]
			stop = stop && nv[n] == 0
		}
		if stop {
			break
		}
		vals = nv
	}
	return res
}

func nineB(vals []int) int {
	res := 0
	idx := 1
	for {
		res += vals[0] * idx
		idx *= -1
		nv := make([]int, len(vals)-1)
		stop := true
		for n := range len(vals) - 1 {
			nv[n] = vals[n+1] - vals[n]
			stop = stop && nv[n] == 0
		}
		if stop {
			break
		}
		vals = nv
	}
	return res
}
