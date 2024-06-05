package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func sixteen() {
	lines, _ := toLineArray("16.in")

	res := make([]int, 0, len(lines)*2+len(lines[0])*2)

	for n := range len(lines) {
		res = append(res, energize(lines, fmt.Sprint(n*len(lines[0]), "R")))
		res = append(res, energize(lines, fmt.Sprint((n+1)*len(lines[0])-1, "L")))
	}

	for n := range len(lines[0]) {
		res = append(res, energize(lines, fmt.Sprint(n, "D")))
		res = append(res, energize(lines, fmt.Sprint(len(lines[0])*len(lines)-1-n, "U")))
	}

	fmt.Println(res[0], slices.Max(res))
}

type Node struct {
	val  string
	next *Node
}

func energize(lines []string, start string) int {

	pattern := strings.Join(lines, "")
	seen := make(map[string]bool, len(pattern)*2)
	energized := make([]bool, len(pattern))
	up := -len(lines[0])
	to := map[byte]map[byte]string{
		'R': {'.': "R", '-': "R", '/': "U", '\\': "D", '|': "UD"},
		'L': {'.': "L", '-': "L", '/': "D", '\\': "U", '|': "UD"},
		'U': {'.': "U", '-': "RL", '/': "R", '\\': "L", '|': "U"},
		'D': {'.': "D", '-': "RL", '/': "L", '\\': "R", '|': "D"},
	}

	dir := map[rune]int{
		'R': 1, 'L': -1, 'U': up, 'D': -up,
	}

	queue := &Node{val: start, next: nil}

	for queue != nil {
		value := queue.val
		queue = queue.next
		_, in := seen[value]
		if in {
			continue
		}

		direction := value[len(value)-1]
		pos, _ := strconv.Atoi(value[:len(value)-1])

		seen[value] = true
		energized[pos] = true

		for _, d := range to[direction][pattern[pos]] {
			x, y := pos/len(lines[0]), pos%len(lines[0])
			if x == 0 && d == 'U' {
				continue
			} else if x == len(lines)-1 && d == 'D' {
				continue
			} else if y == 0 && d == 'L' {
				continue
			} else if y == len(lines[0])-1 && d == 'R' {
				continue
			}

			delta := dir[d]

			queue = &Node{val: fmt.Sprint(pos+delta, string(d)), next: queue}
		}
	}

	res := 0
	for _, v := range energized {
		if v {
			res += 1
		}
	}
	return res
}
