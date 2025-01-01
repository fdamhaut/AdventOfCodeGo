package main

import (
	"fmt"
)

func twelve() {
	carte := toLineArray("12.in")
	a, b := twelveSolve(carte)
	fmt.Println(a, b)
}

func twelveSolve(carte []string) (int, int) {
	region := make(map[int]int)
	perim := make(map[int]int)
	area := make(map[int]int)
	side := make(map[int]int)
	for i, line := range carte {
		for j, r := range line {
			hash := i*len(carte[0]) + j
			region[hash] = hash
			perim[hash] = 4
			area[hash] = 1
			side[hash] = 0
			if i > 0 && r == rune(carte[i-1][j]) {
				uphash := region[(i-1)*len(carte[0])+j]
				twelveMerge(region, perim, area, side, hash, uphash)
			}
			if j > 0 && r == rune(carte[i][j-1]) {
				lhash := region[i*len(carte[0])+j-1]
				twelveMerge(region, perim, area, side, hash, lhash)
			}
			if i > 0 && j > 0 {
				uplhash := region[(i-1)*len(carte[0])+j-1]
				uphash := region[(i-1)*len(carte[0])+j]
				lhash := region[i*len(carte[0])+j-1]
				if (hash == uphash) && (hash == uplhash) && (hash == lhash) {
					continue
				} else if (hash == uphash) && ((hash == uplhash) || (hash == lhash)) {
					side[hash] += 2
				} else if (hash == uplhash) && (hash == lhash) {
					side[hash] += 2
				} else if (uphash == lhash) && (uphash == uplhash) {
					side[uphash] += 2
				}
			}
		}
	}

	a := 0
	b := 0
	done := make(map[int]bool)
	for _, v := range region {
		if !done[v] {
			a += perim[v] * area[v]
			fmt.Println(4+side[v], area[v], v)
			b += (4 + side[v]) * area[v]
			done[v] = true
		}
	}
	return a, b
}

func twelveMerge(region map[int]int, perim map[int]int, area map[int]int, side map[int]int, r1 int, r2 int) {
	for k, v := range region {
		if v == r2 {
			region[k] = r1
		}
	}
	if r1 != r2 {
		perim[r1] += perim[r2]
		area[r1] += area[r2]
		side[r1] += side[r2]
	}
	perim[r1] -= 2
}
