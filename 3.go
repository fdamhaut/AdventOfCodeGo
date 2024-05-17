package main

import (
	"fmt"
	"unicode"
)

func three() {
	data, err := toLineArray("3.in")
	check(err)

	data_rune := make([][]rune, len(data))
	for n, s := range data {
		data_rune[n] = []rune(s)
	}

	var a int = solve3A(data_rune)
	var b int = solve3B(data_rune)

	fmt.Println(a, b)
}

func isSym(r rune) bool {
	return r != rune('.') && !unicode.IsDigit(r)
}

func adjSym(m [][]rune, x int, y int) bool {
	for _, dx := range []int{-1, 0, 1} {
		nx := x + dx
		if nx < 0 || nx >= len(m) {
			continue
		}
		for _, dy := range []int{-1, 0, 1} {
			ny := y + dy
			if ny < 0 || ny >= len(m[nx]) {
				continue
			}
			if isSym(m[nx][ny]) {
				return true
			}
		}
	}
	return false
}

func solve3A(data [][]rune) int {
	sum := 0
	for x, line := range data {
		var val int = 0
		var adj bool = false
		for y, r := range line {
			if unicode.IsDigit(r) {
				val = val*10 + (int(r) - 48)
				adj = adj || adjSym(data, x, y)
			} else {
				if adj {
					sum += val
				}
				val = 0
				adj = false
			}
		}
		if adj {
			sum += val
		}
	}
	return sum
}

func isGear(r rune) bool {
	return r == rune('*')
}

func adjGear(m [][]rune, x int, y int) [][]int {
	res := [][]int{}
	for _, dx := range []int{-1, 0, 1} {
		nx := x + dx
		if nx < 0 || nx >= len(m) {
			continue
		}
		for _, dy := range []int{-1, 0, 1} {
			ny := y + dy
			if ny < 0 || ny >= len(m[nx]) {
				continue
			}
			if isGear(m[nx][ny]) {
				res = append(res, []int{nx, ny})
			}
		}
	}
	return res
}

func removeDuplicate(array [][]int) [][]int {
	set := make(map[int]map[int]bool, len(array))
	res := [][]int{}

	for _, v := range array {
		in := set[v[0]][v[1]]
		if !in {
			res = append(res, v)
			_, exists := set[v[0]]
			if !exists {
				set[v[0]] = map[int]bool{}
			}
			set[v[0]][v[1]] = true
		}
	}
	return res
}

func addM2(m map[int]map[int][]int, x int, y int, val int) map[int]map[int][]int {
	_, exists := m[x]
	if !exists {
		m[x] = map[int][]int{}
	}
	my, exists := m[x][y]
	if !exists {
		m[x][y] = []int{val}
	} else {
		m[x][y] = append(my, val)
	}
	return m
}

func solve3B(data [][]rune) int {
	gears := map[int]map[int][]int{}
	for x, line := range data {
		var val int = 0
		adj := [][]int{}
		for y, r := range line {
			if unicode.IsDigit(r) {
				val = val*10 + (int(r) - 48)
				adj = append(adj, adjGear(data, x, y)...)
			} else {
				for _, g := range removeDuplicate(adj) {
					gx, gy := g[0], g[1]
					gears = addM2(gears, gx, gy, val)
				}
				val = 0
				adj = [][]int{}
			}
		}
		for _, g := range removeDuplicate(adj) {
			gx, gy := g[0], g[1]
			gears = addM2(gears, gx, gy, val)
		}
	}

	sum := 0
	for _, v := range gears {
		for _, g := range v {
			if len(g) == 2 {
				sum += g[0] * g[1]
			}
		}
	}
	return sum
}
