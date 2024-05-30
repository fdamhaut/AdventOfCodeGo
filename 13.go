package main

import "fmt"

func thirteen() {
	lines, _ := toLineArray("13.in")

	cases := make([][]string, 0, 100)

	current := make([]string, 0, 20)
	for _, l := range lines {
		if l == "" {
			cases = append(cases, current)
			current = make([]string, 0, 20)
		} else {
			current = append(current, l)
		}
	}

	a, b := 0, 0

	for _, c := range cases {
		a += solve13A(c)
		b += solve13B(c)
	}

	fmt.Println(a, b)
}

func solve13A(pattern []string) int {
	for n := range len(pattern) - 1 {
		if mirror(pattern[:n+1], pattern[n+1:]) {
			return (n + 1) * 100
		}
	}
	pattern_t := transpose(pattern)
	for n := range len(pattern_t) - 1 {
		if mirror(pattern_t[:n+1], pattern_t[n+1:]) {
			return (n + 1)
		}
	}
	return 0
}

func solve13B(pattern []string) int {
	for n := range len(pattern) - 1 {
		if dist(pattern[:n+1], pattern[n+1:]) == 1 {
			return (n + 1) * 100
		}
	}
	pattern_t := transpose(pattern)
	for n := range len(pattern_t) - 1 {
		if dist(pattern_t[:n+1], pattern_t[n+1:]) == 1 {
			return (n + 1)
		}
	}
	return 0
}

func mirror(top []string, bot []string) bool {

	for n := range min(len(top), len(bot)) {
		if top[len(top)-1-n] != bot[n] {
			return false
		}
	}
	return true
}

func dist(top []string, bot []string) int {
	d := 0
	for n := range min(len(top), len(bot)) {
		for i := range len(top[0]) {
			if top[len(top)-1-n][i] != bot[n][i] {
				d += 1
				if d > 1 {
					return 2
				}
			}
		}
	}
	return d
}

func transpose(table []string) []string {
	res := make([]string, len(table[0]))

	for _, line := range table {
		for y, char := range line {
			res[y] += string(char)
		}
	}
	return res
}
