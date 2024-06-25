package main

import (
	"fmt"
	"strconv"
	"strings"
)

func eighteen() {
	data, _ := toLineArray("18.in")

	dirA, lenA := make([]int, len(data)), make([]int, len(data))
	dirB, lenB := make([]int, len(data)), make([]int, len(data))

	for _, d := range data {
		da, la, db, lb := get18dirlen(d)
		fmt.Println(da, la, db, lb)
		dirA = append(dirA, da)
		lenA = append(lenA, la)
		dirB = append(dirB, db)
		lenB = append(lenB, lb)
	}

	fmt.Println(solve18(dirA, lenA), solve18(dirB, lenB))
}

func get18dirlen(line string) (int, int, int, int) {
	data_line := strings.Split(line, " ")
	var da, la, db, lb int

	if data_line[0] == "R" {
		da = 0
	} else if data_line[0] == "D" {
		da = 1
	} else if data_line[0] == "L" {
		da = 2
	} else {
		da = 3
	}

	la, _ = strconv.Atoi(data_line[1])

	db = int(data_line[2][len(data_line[2])-2]) - 48

	lb64, _ := strconv.ParseInt(data_line[2][2:len(data_line[2])-2], 16, 64)
	lb = int(lb64)

	return da, la, db, lb
}

func solve18(directions []int, lengths []int) int {
	area := 0
	x, y := 0, 0
	nx, ny := 0, 0
	for n := range len(directions) {
		if directions[n] == 0 {
			ny += lengths[n]
		} else if directions[n] == 1 {
			nx -= lengths[n]
		} else if directions[n] == 2 {
			ny -= lengths[n]
		} else {
			nx += lengths[n]
		}
		area += x*ny - y*nx + lengths[n] // boundary are only counted as 1/2 (because assume to go through middle)
		x, y = nx, ny
	}

	return area/2 + 1 // +1 is because the 4 corners are only counted for 3/4 (1/4 is formula, 1/2 is boundary above)
}
