package main

import (
	"fmt"
	"slices"
)

func one() {
	data := toIntMatrix("1.in")
	l1 := make([]int, len(data))
	l2 := make([]int, len(data))
	for i, v := range data {
		l1[i] = v[0]
		l2[i] = v[1]
	}
	a := oneA(l1, l2)
	b := oneB(l1, l2)
	fmt.Println(a, b)
}

func oneA(l1 []int, l2 []int) int {
	slices.Sort(l1)
	slices.Sort(l2)
	res := 0
	for i := range l1 {
		val := l1[i] - l2[i]
		if val > 0 {
			res += val
		} else {
			res -= val
		}
	}
	return res
}

func count(l []int) map[int]int {
	res := make(map[int]int)
	for _, v := range l {
		res[v] += 1
	}
	return res
}

func oneB(l1 []int, l2 []int) int {
	count1 := count(l1)
	count2 := count(l2)
	res := 0
	for k, _ := range count1 {
		res += k * count1[k] * count2[k]
	}
	return res
}
