package main

import (
	"fmt"
	"strconv"
	"strings"
)

func four() {
	data, err := toLineArray("4.in")
	check(err)

	card_amount := make([]int, len(data))

	a, b := 0, len(data)
	for n, s := range data {
		val := fourAlineval(s)
		if val >= 1 {
			a += 1 << (val - 1)
		}
		for nn := range val {
			if n+nn+1 < len(data) {
				card_amount[n+nn+1] += card_amount[n] + 1
				b += card_amount[n] + 1
			}
		}
	}
	fmt.Println(a, b)

}

func fourAlineval(s string) int {
	cn := strings.Split(s[strings.Index(s, ":")+1:], " | ")
	cards, nums := strings.Split(cn[0], " "), strings.Split(cn[1], " ")

	cards_set := map[int]bool{}

	for _, v := range cards {
		vi, err := strconv.Atoi(v)
		if err == nil {
			cards_set[vi] = true
		}
	}

	val := 0
	for _, v := range nums {
		vi, err := strconv.Atoi(v)
		if err == nil && cards_set[vi] {
			val += 1
		}
	}

	return val
}
