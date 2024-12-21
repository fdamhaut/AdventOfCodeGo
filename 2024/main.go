package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	funs := []func(){one}
	i, err := strconv.Atoi(os.Args[1])
	check(err)
	funs[i-1]()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func toLineArray(path string) []string {
	dat, err := os.ReadFile(path)
	check(err)
	return strings.Split(string(dat), "\n")
}

func toIntArray(s string) []int {
	var res []int

	for _, v := range strings.Split(s, " ") {
		i, err := strconv.Atoi(v)
		if err == nil {
			res = append(res, i)
		}
	}

	return res
}

func toIntMatrix(path string) [][]int {
	lineArray := toLineArray(path)
	var res [][]int = make([][]int, len(lineArray))
	for i, s := range lineArray {
		res[i] = toIntArray(s)
	}
	return res
}
