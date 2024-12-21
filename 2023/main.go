package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	eighteen()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func toLineArray(path string) ([]string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(dat), "\n"), nil
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
