package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	funs := []func(){one, two, three, four, five, six, seven, eight, nine, ten, eleven, twelve,
		thirteen, fourteen, fifteen, sixteen, seventeen, eighteen, nineteen, twenty, twentyone}
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

func toIntArray(s string, sep ...string) []int {
	var res []int
	separator := " "
	if len(sep) > 0 {
		separator = sep[0]
	}
	for _, v := range strings.Split(s, separator) {
		i, err := strconv.Atoi(v)
		if err == nil {
			res = append(res, i)
		}
	}

	return res
}

func toIntMatrix(path string, sep ...string) [][]int {
	return _toIntMatrix(toLineArray(path), sep...)
}

func _toIntMatrix(strings []string, sep ...string) [][]int {
	var res [][]int = make([][]int, len(strings))
	for i, s := range strings {
		res[i] = toIntArray(s, sep...)
	}
	return res
}
