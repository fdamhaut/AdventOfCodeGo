package main

import (
	"os"
	"strings"
)

func main() {
	one()
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
