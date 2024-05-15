package main

import (
	"fmt"
	"strings"
	"unicode"
)

func one() {
	data, err := toLineArray("1.in")
	check(err)
	a, b := 0, 0
	m := make(map[string]int)
	for n, v := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
		m[v] = n + 1
	}
	for n, v := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"} {
		m[v] = n + 1
	}
	for _, s := range data {
		a += line2valA(s)
		b += line2valB(s, m)
	}
	fmt.Println(a, b)
}

func firstDigitA(rs []rune, reverse bool) int {
	lr := len(rs)
	var char rune

	for i := 0; i < lr; i = i + 1 {
		if reverse {
			char = rs[lr-1-i]
		} else {
			char = rs[i]
		}
		if unicode.IsDigit(char) {
			return int(char) - 48
		}
	}
	return 0
}

func line2valA(s string) int {
	rs := []rune(s)
	return firstDigitA(rs, false)*10 + firstDigitA(rs, true)
}

func firstDigitB(s string, m map[string]int) int {
	var min_index int = len(s) + 1
	var val int
	for k, v := range m {
		index := strings.Index(s, k)
		if index >= 0 && index <= min_index {
			min_index = index
			val = v
		}
	}
	return val
}

func lastDigitB(s string, m map[string]int) int {
	var min_index int
	var val int
	for k, v := range m {
		index := strings.LastIndex(s, k)
		if index >= 0 && index >= min_index {
			min_index = index
			val = v
		}
	}
	return val
}

func line2valB(s string, m map[string]int) int {
	var val int = firstDigitB(s, m)*10 + lastDigitB(s, m)
	return val
}
