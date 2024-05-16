package main

import (
	"fmt"
	"strconv"
	"strings"
)

func two() {
	data, err := toLineArray("2.in")
	check(err)

	var a int = 0
	var b int = 0

	for n, s := range data {
		if gamePossible(s) {
			a += n + 1
		}
		b += minCubePossible(s)
	}
	fmt.Println(a, b)
}

func gamePossible(s string) bool {
	max := map[string]int{"red": 12, "green": 13, "blue": 14}

	s = s[strings.Index(s, ":")+2:]
	rounds := strings.Split(s, "; ")
	for _, r := range rounds {
		colors := strings.Split(r, ", ")
		for _, color := range colors {
			c := strings.Split(color, " ")
			count, _ := strconv.Atoi(c[0])
			if count > max[c[1]] {
				return false
			}
		}
	}
	return true
}

func minCubePossible(s string) int {
	res := make(map[string]int)
	s = s[strings.Index(s, ":")+2:]
	rounds := strings.Split(s, "; ")
	for _, r := range rounds {
		colors := strings.Split(r, ", ")
		for _, color := range colors {
			c := strings.Split(color, " ")
			count, _ := strconv.Atoi(c[0])
			color = c[1]
			res[color] = max(res[color], count)
		}
	}
	power := 1
	for _, v := range res {
		power *= v
	}
	return power
}
