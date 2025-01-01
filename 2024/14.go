package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fourteen() {
	robots := fourteenParse(toLineArray("14.in"))
	q_a := make(map[int]int)
	for _, r := range robots {
		q_a[r.quad(100)] += 1
	}
	a := q_a[0] * q_a[1] * q_a[2] * q_a[3]
	fmt.Println(a)
	for i := range 200 {
		// TODO
		fmt.Println("=========================================================", i, "=========================================================")
		fourteenShow(robots, i)
	}
}

var D14_ROOM_W = 101
var D14_ROOM_H = 103

type fourteenrobot struct {
	x int
	y int

	vx int
	vy int
}

func (r *fourteenrobot) quad(time int) int {
	x := (r.x + r.vx*time) % D14_ROOM_W
	y := (r.y + r.vy*time) % D14_ROOM_H

	for x < 0 {
		x += D14_ROOM_W
	}

	for y < 0 {
		y += D14_ROOM_H
	}

	if x > (D14_ROOM_W-1)/2 {
		if y > (D14_ROOM_H-1)/2 {
			return 2
		} else if y < (D14_ROOM_H-1)/2 {
			return 3
		}

	} else if x < (D14_ROOM_W-1)/2 {
		if y > (D14_ROOM_H-1)/2 {
			return 1
		} else if y < (D14_ROOM_H-1)/2 {
			return 0
		}
	}
	return -1
}

func fourteenParse(data []string) []fourteenrobot {
	res := make([]fourteenrobot, 0, len(data))
	for _, l := range data {
		pv := strings.Split(l, " v=")
		pxy := strings.Split(pv[0][2:], ",")
		x, _ := strconv.Atoi(pxy[0])
		y, _ := strconv.Atoi(pxy[1])

		vxy := strings.Split(pv[1], ",")
		vx, _ := strconv.Atoi(vxy[0])
		vy, _ := strconv.Atoi(vxy[1])

		res = append(res, fourteenrobot{x, y, vx, vy})
	}
	return res
}

func (r *fourteenrobot) pos(time int) (int, int) {
	x := (r.x + r.vx*time) % D14_ROOM_W
	y := (r.y + r.vy*time) % D14_ROOM_H

	for x < 0 {
		x += D14_ROOM_W
	}

	for y < 0 {
		y += D14_ROOM_H
	}
	return x, y
}
func fourteenShow(robots []fourteenrobot, time int) {
	pic := make([][]rune, D14_ROOM_H)

	for i := range pic {
		pic[i] = make([]rune, D14_ROOM_W)
		for j := range pic[i] {
			pic[i][j] = ' '
		}
	}

	for _, r := range robots {
		x, y := r.pos(time)
		pic[y][x] = 'X'
	}

	for _, l := range pic {
		fmt.Println(string(l))
	}
}
