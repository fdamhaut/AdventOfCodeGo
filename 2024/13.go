package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func thirteen() {
	data := toLineArray("13.in")
	problems := thirteenParse(data)

	a := 0
	ab := 0
	b := 0
	for _, p := range problems {
		a += p.solveA()
		ab += p.solveB()
		p.xp += 10000000000000
		p.yp += 10000000000000
		b += p.solveB()
	}
	fmt.Println(a, ab, b)
}

type thirteenProblem struct {
	xa int
	ya int

	xb int
	yb int

	xp int
	yp int
}

func thirteenParse(data []string) []thirteenProblem {
	res := make([]thirteenProblem, 0, len(data)/4)
	for i := 0; i < len(data); i += 4 {
		a := strings.Split(data[i][10:], ", ")
		xa, _ := strconv.Atoi(a[0][2:])
		ya, _ := strconv.Atoi(a[1][2:])

		b := strings.Split(data[i+1][10:], ", ")
		xb, _ := strconv.Atoi(b[0][2:])
		yb, _ := strconv.Atoi(b[1][2:])

		prize := strings.Split(data[i+2][7:], ", ")
		xp, _ := strconv.Atoi(prize[0][2:])
		yp, _ := strconv.Atoi(prize[1][2:])
		res = append(res, thirteenProblem{xa: xa, ya: ya, xb: xb, yb: yb, xp: xp, yp: yp})
	}
	return res
}

func (p *thirteenProblem) solveA() int {
	best := 500
	for a := range 101 {
		b := (p.xp - a*p.xa) / p.xb
		if b > 100 {
			continue
		}
		if a*p.xa+b*p.xb != p.xp || a*p.ya+b*p.yb != p.yp {
			continue
		}
		best = min(best, 3*a+b)
	}
	if best < 500 {
		return best
	}
	return 0
}

func (p *thirteenProblem) solveB() int {

	ad := float64(p.ya) / float64(p.xa)

	divisor := float64(p.xb)*ad - float64(p.yb)

	if divisor == 0 {
		fmt.Println("TODO ?")
	}

	b := (float64(p.xp)*ad - float64(p.yp)) / divisor
	a := (float64(p.xp) - float64(p.xb)*b) / float64(p.xa)

	ia := int(math.Round(a))
	ib := int(math.Round(b))

	if ((ia*p.ya + ib*p.yb) != p.yp) || ((ia*p.xa + ib*p.xb) != p.xp) {
		return 0
	}

	return ia*3 + ib
}
