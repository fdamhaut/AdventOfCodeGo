package main

import "fmt"

func eleven() {
	lines, _ := toLineArray("11.in")

	empty_lines := make([]bool, len(lines))
	empty_cols := make([]bool, len(lines[0]))

	galaxies := make([]int, 0, 1e5)

	for x, line := range lines {
		for y, val := range line {
			if val == '#' {
				empty_lines[x] = true
				empty_cols[y] = true
				galaxies = append(galaxies, x*len(lines[0])+y)
			}
		}
	}

	for n, v := range empty_cols {
		empty_cols[n] = !v
	}
	for n, v := range empty_lines {
		empty_lines[n] = !v
	}

	a, b := 0, 0

	for n, g1 := range galaxies {
		for _, g2 := range galaxies[n+1:] {
			a += distance(g1, g2, empty_lines, empty_cols, 2)
			b += distance(g1, g2, empty_lines, empty_cols, 1000000)
		}
	}

	fmt.Println(a, b)
}

func distance(pos1 int, pos2 int, empty_lines []bool, empty_cols []bool, empty_dist int) int {
	empty_dist -= 1
	x1, y1 := pos1/len(empty_cols), pos1%len(empty_cols)
	x2, y2 := pos2/len(empty_cols), pos2%len(empty_cols)

	x1, x2 = min(x1, x2), max(x1, x2)
	y1, y2 = min(y1, y2), max(y1, y2)

	return x2 - x1 + y2 - y1 + empty_dist*(count(empty_lines[x1:x2])+count(empty_cols[y1:y2]))
}

func count(arr []bool) int {
	res := 0
	for _, v := range arr {
		if v {
			res += 1
		}
	}
	return res
}
