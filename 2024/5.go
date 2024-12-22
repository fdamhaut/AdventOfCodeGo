package main

import (
	"fmt"
	"sort"
)

func five() {
	data := toLineArray("5.in")
	ruless := data[:1176]
	rules := fiveMakeRules(ruless)
	books := _toIntMatrix(data[1177:], ",")
	a := fiveA(rules, books)
	b := fiveB(rules, books)
	fmt.Println(a, b)
}

func fiveMakeRules(rules []string) map[int]map[int]bool {
	res := make(map[int]map[int]bool)

	for _, rule := range rules {
		ri := toIntArray(rule, "|")
		if res[ri[0]] == nil {
			res[ri[0]] = make(map[int]bool)
		}
		res[ri[0]][ri[1]] = true
	}

	return res
}

func fiveA(rules map[int]map[int]bool, books [][]int) int {
	res := 0
	for _, book := range books {
		if fiveIsOrdered(rules, book) {
			res += book[len(book)/2]
		}
	}
	return res
}

func fiveB(rules map[int]map[int]bool, books [][]int) int {
	res := 0
	for _, book := range books {
		if !fiveIsOrdered(rules, book) {
			sort.Slice(book, func(i, j int) bool {
				return rules[book[i]][book[j]]
			})
			res += book[len(book)/2]
		}
	}
	return res
}

func fiveIsOrdered(rules map[int]map[int]bool, book []int) bool {
	for i, v1 := range book[:len(book)-1] {
		for _, v2 := range book[i+1:] {
			if rules[v2][v1] {
				return false
			}
		}
	}
	return true
}
