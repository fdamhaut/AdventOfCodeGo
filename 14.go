package main

import (
	"fmt"
	"strings"
)

func fourteen() {
	lines, _ := toLineArray("14.in")

	a := solve14A(lines)
	b := solve14B(lines)

	fmt.Println(a, b)
}

func solve14A(pattern []string) int {
	return count14(transpose(pattern))
}

func count14(pattern []string) int {
	total := 0
	for _, s := range pattern {
		idx := 0
		for n, r := range s {
			switch r {
			case 'O':
				idx++
				total += len(s) - idx + 1
			case '#':
				idx = n + 1
			}
		}
	}
	return total
}

func roll(pattern []string) []string {
	for x, s := range pattern {
		idx := 0
		ns := make([]rune, len(s))
		for n := range ns {
			ns[n] = '.'
		}
		for n, r := range s {
			switch r {
			case 'O':
				ns[idx] = 'O'
				idx++
			case '#':
				ns[n] = '#'
				idx = n + 1
			}
		}
		pattern[x] = string(ns)
	}
	return pattern
}

func rotate(pattern []string) []string {
	res := make([]string, len(pattern[0]))

	for _, line := range pattern {
		for y, char := range line {
			res[len(res)-1-y] += string(char)
		}
	}
	return res
}

func solve14B(pattern []string) int {
	pattern = transpose(pattern)

	loops := make(map[string]int)

	new := ""
	n := 0

	for n < 1e9 {
		pattern = roll(pattern)
		pattern = rotate(pattern)
		pattern = roll(pattern)
		pattern = rotate(pattern)
		pattern = roll(pattern)
		pattern = rotate(pattern)
		pattern = roll(pattern)
		pattern = rotate(pattern)
		new = strings.Join(pattern, "")
		val, ok := loops[new]
		if ok {
			loop := n - val
			for n < 1e9 {
				n += loop
			}
			n -= loop
		}
		loops[new] = n
		n++
	}
	return count14(pattern)
}

func printp14(pattern []string) {
	for _, s := range pattern {
		fmt.Println(s)
	}
	fmt.Println("=========================================")
}
