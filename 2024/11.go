package main

import (
	"fmt"
	"strconv"
)

func eleven() {
	stonei := toIntArray(toLineArray("11.in")[0])
	stones := make([]elevenStone, len(stonei))
	for i, val := range stonei {
		stones[i] = elevenVali(val)
	}

	a := 0
	b := 0
	for _, stone := range stones {
		a += step(stone, 25)
		b += step(stone, 75)
	}
	fmt.Println(a, b)
}

type elevenStone struct {
	vali int
	vals string
}

func (s *elevenStone) next() []elevenStone {
	if s.vali == 0 {
		return []elevenStone{{vali: 1, vals: "1"}}
	} else if len(s.vals)%2 == 0 {
		return []elevenStone{elevenVals(s.vals[:len(s.vals)/2]), elevenVals(s.vals[len(s.vals)/2:])}
	} else {
		return []elevenStone{elevenVali(s.vali * 2024)}
	}
}

func elevenVali(vali int) elevenStone {
	return elevenStone{vali: vali, vals: fmt.Sprint(vali)}
}

func elevenVals(vals string) elevenStone {
	vali, _ := strconv.Atoi(vals)
	return elevenStone{vali: vali, vals: fmt.Sprint(vali)}
}

var cache map[int]map[int]int = make(map[int]map[int]int)

func step(stone elevenStone, amount int) int {
	if amount == 0 {
		return 1
	}
	if cache[stone.vali] == nil {
		cache[stone.vali] = make(map[int]int)
	}
	if cache[stone.vali][amount] != 0 {
		return cache[stone.vali][amount]
	}
	res := 0
	for _, s := range stone.next() {
		res += step(s, amount-1)
	}
	cache[stone.vali][amount] = res
	return res
}
