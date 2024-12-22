package main

import (
	"fmt"
	"strings"
)

func six() {
	data := toLineArray("6.in")
	a := sixA(data)
	b := sixB(data)
	fmt.Println(a, b)
}

func sixA(carte []string) int {
	pos := sixStart(carte)
	poses := make(map[int]bool)
	poses[pos.hash(carte)] = true
	for pos.next(carte) {
		poses[pos.hash(carte)] = true
	}
	return len(poses)
}

func sixB(carte []string) int {
	res := 0
	pos := sixStart(carte)
	poses := make(map[int]bool)
	for pos.next(carte) {
		poses[pos.hash(carte)] = true
	}
	pos = sixStart(carte)
	delete(poses, pos.hash(carte))
	for pos := range poses {
		x, y := pos%len(carte[0]), pos/len(carte[0])
		carte[x] = replaceAtIndex(carte[x], '#', y)
		if sixLoop(carte) {
			res += 1
		}
		carte[x] = replaceAtIndex(carte[x], '.', y)
	}

	return res
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func sixLoop(carte []string) bool {
	pos := sixStart(carte)
	poses := make(map[int]bool)

	poses[pos.hashB(carte)] = true
	for pos.next(carte) {
		if poses[pos.hashB(carte)] {
			return true
		}
		poses[pos.hashB(carte)] = true
	}
	return false
}

type pos struct {
	x   int
	y   int
	dir int
}

func (p *pos) hash(carte []string) int {
	return p.x + p.y*len(carte[0])
}

func (p *pos) hashB(carte []string) int {
	return p.x + p.y*len(carte[0]) + p.dir*len(carte[0])*len(carte)
}

func (p *pos) next(carte []string) bool {
	nx := p.x
	ny := p.y
	if p.dir == 0 {
		nx -= 1
	} else if p.dir == 1 {
		ny += 1
	} else if p.dir == 2 {
		nx += 1
	} else if p.dir == 3 {
		ny -= 1
	}
	if nx < 0 || nx >= len(carte) {
		return false
	} else if ny < 0 || ny >= len(carte[0]) {
		return false
	}
	if carte[nx][ny] == '#' {
		p.dir = (p.dir + 1) % 4
		return p.next(carte)
	} else {
		p.x = nx
		p.y = ny
	}
	return true
}

func sixStart(carte []string) pos {
	for i, line := range carte {
		j := strings.Index(line, "^")
		if j > 0 {
			return pos{
				x:   i,
				y:   j,
				dir: 0,
			}
		}
	}
	return pos{}
}
