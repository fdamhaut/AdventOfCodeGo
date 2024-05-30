package main

import (
	"fmt"
	"strconv"
	"strings"
)

func twelve() {
	lines, _ := toLineArray("12.in")

	a, b := 0, 0

	for _, line := range lines {
		l := strings.Split(line, " ")
		a += solve12(l[0]+".", l[1])
		five_op := strings.Repeat(l[1]+",", 5)
		five_s := strings.Repeat(l[0]+"?", 5)
		b += solve12(five_s[:len(five_s)-1]+".", five_op[:len(five_op)-1])
	}

	fmt.Println(a, b)
}

func solve12(pattern string, next string) int {

	np := strings.ReplaceAll(pattern, "..", ".")

	for pattern != np {
		pattern = np
		np = strings.ReplaceAll(pattern, "..", ".")
	}

	res := solve12rec(pattern, next, make(map[string]int))
	return res
}

func solve12rec(pattern string, next string, cache map[string]int) int {

	if len(pattern) == 0 && len(next) == 0 {
		return 1
	}

	if len(pattern) > 0 && len(next) == 0 {
		if strings.Contains(pattern, "#") {
			return 0
		} else {
			return 1
		}
	}

	if len(pattern) == 0 && len(next) > 0 {
		return 0
	}

	val, in := cache[pattern+next]

	if in {
		return val
	}

	comma := strings.Index(next, ",")
	next_int := 0
	next_next := ""

	if comma == -1 {
		next_int, _ = strconv.Atoi(next)
	} else {
		next_int, _ = strconv.Atoi(next[:comma])
		next_next = next[comma+1:]
	}

	if len(pattern) > next_int && possible12(pattern[:next_int+1]) {
		val += solve12rec(pattern[next_int+1:], next_next, cache)
	}
	if pattern[0] == '.' || pattern[0] == '?' {
		val += solve12rec(pattern[1:], next, cache)
	}

	cache[pattern+next] = val
	return val
}

func possible12(s string) bool {
	for _, v := range s[:len(s)-1] {
		if v == '.' {
			return false
		}
	}
	return s[len(s)-1] != '#'
}
