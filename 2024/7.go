package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func seven() {
	data := toLineArray("7.in")
	a := 0
	b := 0
	for _, s := range data {
		problem := sevenToProblem(s)
		if problem.solveA() {
			a += problem.value
		}
		if problem.solveB() {
			b += problem.value
		}
	}
	fmt.Println(a, b)
}

type sevenProblem struct {
	value int
	item  []int
}

func (p *sevenProblem) solveA() bool {
	possibles := []int{p.item[0]}
	for _, i := range p.item[1:] {
		np := make([]int, 0, len(possibles)*2)
		for _, v := range possibles {
			for _, nv := range []int{v + i, v * i} {
				if nv <= p.value {
					np = append(np, nv)
				}
			}
		}
		possibles = np
	}
	return slices.Contains(possibles, p.value)
}

func (p *sevenProblem) solveB() bool {
	possibles := []int{p.item[0]}
	for _, i := range p.item[1:] {
		np := make([]int, 0, len(possibles)*2)
		for _, v := range possibles {
			for _, nv := range []int{v + i, v * i, v*int(math.Pow10(len(fmt.Sprint(i)))) + i} {
				if nv <= p.value {
					np = append(np, nv)
				}
			}
		}
		possibles = np
	}
	return slices.Contains(possibles, p.value)
}

func sevenToProblem(s string) sevenProblem {
	vals := strings.Split(s, ":")
	value, _ := strconv.Atoi(vals[0])
	return sevenProblem{
		value: value,
		item:  toIntArray(vals[1]),
	}
}
