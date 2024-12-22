package main

import (
	"fmt"
)

func two() {
	in := toIntMatrix("2.in")
	a := twoA(in)
	b := twoB(in)
	fmt.Println(a, b)
}

func twoA(in [][]int) int {
	res := 0
	for _, line := range in {
		res += twoInc(line)
		res += twoDec(line)
	}
	return res
}

func twoB(in [][]int) int {
	res := 0
	for _, line := range in {
		res += twoBInc(line) | twoBDec(line)
	}
	return res
}

func twoInc(line []int) int {
	old := line[0]
	for _, v := range line[1:] {
		if v <= old || v > old+3 {
			return 0
		}
		old = v
	}
	return 1
}

func twoBInc(line []int) int {
	if twoInc(line[1:]) > 0 {
		return 1
	}
	old := line[0]
	for i, v := range line[1:] {
		if v <= old || v > old+3 {
			l1, l2 := make([]int, len(line)), make([]int, len(line))
			copy(l1, line)
			copy(l2, line)
			return twoInc(append(l1[:i], l1[i+1:]...)) | twoInc(append(l2[:i+1], l2[i+2:]...))
		} else {
			old = v
		}
	}
	return 1
}

func twoDec(line []int) int {
	old := line[0]
	for _, v := range line[1:] {
		if v >= old || v < old-3 {
			return 0
		}
		old = v
	}
	return 1
}

func twoBDec(line []int) int {
	if twoDec(line[1:]) > 0 {
		return 1
	}
	old := line[0]
	for i, v := range line[1:] {
		if v >= old || v < old-3 {
			l1, l2 := make([]int, len(line)), make([]int, len(line))
			copy(l1, line)
			copy(l2, line)
			return twoDec(append(l1[:i], l1[i+1:]...)) | twoDec(append(l2[:i+1], l2[i+2:]...))
		} else {
			old = v
		}
	}
	return 1
}
