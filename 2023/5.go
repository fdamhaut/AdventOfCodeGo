package main

import (
	"fmt"
	"strconv"
	"strings"
)

func five() {
	data, _ := toLineArray("5.in")

	maps := make([][][]int, 7)

	seeds := data[0]

	data = data[2:]

	last_index := 0
	maps_index := 0

	for n, line := range data {
		if line == "" {
			maps[maps_index] = lines2map(data[last_index:n])
			maps_index++
			last_index = n + 1
		}
	}
	maps[maps_index] = lines2map(data[last_index:])

	seeds = seeds[strings.Index(seeds, ":")+2:]

	a := 999999999999999999
	b := 999999999999999999

	seeds_str := strings.Split(seeds, " ")
	seeds_int := make([]int, len(seeds_str))

	for n, seed := range seeds_str {
		s, _ := strconv.Atoi(seed)
		seeds_int[n] = s
	}

	for _, s := range seeds_int {
		for _, m := range maps {
			s = mapped(s, m)
		}
		a = min(a, s)
	}

	r := seed2ranges(seeds_int)
	for _, m := range maps {
		r = rangesmapped(r, m)
	}
	for _, v := range r {
		b = min(b, v[0])
	}

	fmt.Println(a, b)

}

func seed2ranges(seed []int) [][]int {
	var has_start bool = false
	var from int
	res := make([][]int, len(seed)/2)

	for n, s := range seed {
		if !has_start {
			from = s
			has_start = true
		} else {
			res[n/2] = []int{from, from + s}
			has_start = false
		}
	}
	return res

}

func lines2map(lines []string) [][]int {

	lines = lines[1:]

	res := make([][]int, len(lines))

	for n, line := range lines {
		vals := strings.Split(line, " ")
		dest, _ := strconv.Atoi(vals[0])
		source, _ := strconv.Atoi(vals[1])
		ran, _ := strconv.Atoi(vals[2])

		res[n] = []int{source, source + ran, dest - source}
	}
	return res
}

func mapped(val int, ma [][]int) int {
	for _, m := range ma {
		if m[0] <= val && val < m[1] {
			return val + m[2]
		}
	}
	return val
}

func rangesmapped(ranges [][]int, ma [][]int) [][]int {
	var res [][]int
	for _, val := range ranges {
		res = append(res, rangemapped(val, ma)...)
	}
	return res
}

func rangemapped(ran []int, ma [][]int) [][]int {
	from, to := ran[0], ran[1]

	var res [][]int

	for _, m := range ma {
		rmin, rmax, rmap := m[0], m[1], m[2]
		if rmin <= from && to <= rmax {
			return append(res, []int{from + rmap, to + rmap})
		} else if rmin <= from && from < rmax {
			res = append(res, []int{from + rmap, rmax + rmap})
			from = rmax
		} else if rmin <= to && to < rmax {
			res = append(res, []int{rmin + rmap, to + rmap})
			to = rmin
		}
	}

	if to > from {
		res = append(res, []int{from, to})
	}

	return res
}
