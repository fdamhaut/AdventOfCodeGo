package main

import "fmt"

func ten() {
	carte := toLineArray("10.in")
	a := tenSolveA(carte)
	b := tenSolveB(carte)
	fmt.Println(a, b)
}

type tenPos struct {
	x   int
	y   int
	val byte
}

var tenNEXTDICT = map[byte]byte{
	'0': '1',
	'1': '2',
	'2': '3',
	'3': '4',
	'4': '5',
	'5': '6',
	'6': '7',
	'7': '8',
	'8': '9',
}

func (p *tenPos) next(carte []string) []tenPos {
	res := make([]tenPos, 0, 4)
	x, y := p.x, p.y
	if p.x > 0 && carte[x-1][y] == tenNEXTDICT[p.val] {
		res = append(res, tenPos{x: x - 1, y: y, val: tenNEXTDICT[p.val]})
	}
	if p.x < len(carte)-1 && carte[x+1][y] == tenNEXTDICT[p.val] {
		res = append(res, tenPos{x: x + 1, y: y, val: tenNEXTDICT[p.val]})
	}
	if p.y > 0 && carte[x][y-1] == tenNEXTDICT[p.val] {
		res = append(res, tenPos{x: x, y: y - 1, val: tenNEXTDICT[p.val]})
	}
	if p.y < len(carte[0])-1 && carte[x][y+1] == tenNEXTDICT[p.val] {
		res = append(res, tenPos{x: x, y: y + 1, val: tenNEXTDICT[p.val]})
	}
	return res
}

func (p *tenPos) hash(carte []string) int {
	return p.x*len(carte[0]) + p.y
}

func (p *tenPos) numPaths(carte []string) int {
	poses := []tenPos{*p}
	res := make(map[int]bool)
	for len(poses) > 0 {
		newP := make([]tenPos, 0, 4*len(poses))
		for _, p := range poses {
			for _, np := range p.next(carte) {
				if np.val == '9' {
					res[np.hash(carte)] = true
				} else {
					newP = append(newP, np)
				}
			}
		}
		poses = newP
	}
	return len(res)
}

func (p *tenPos) numPathsB(carte []string) int {
	poses := []tenPos{*p}
	res := 0
	for len(poses) > 0 {
		newP := make([]tenPos, 0, 4*len(poses))
		for _, p := range poses {
			for _, np := range p.next(carte) {
				if np.val == '9' {
					res += 1
				} else {
					newP = append(newP, np)
				}
			}
		}
		poses = newP
	}
	return res
}

func tenSolveA(carte []string) int {
	res := 0
	for _, pos := range tenStart(carte) {
		res += pos.numPaths(carte)
	}
	return res
}

func tenSolveB(carte []string) int {
	res := 0
	for _, pos := range tenStart(carte) {
		res += pos.numPathsB(carte)
	}
	return res
}

func tenStart(carte []string) []tenPos {
	res := make([]tenPos, 0, 100)
	for i, line := range carte {
		for j, r := range line {
			if r == '0' {
				res = append(res, tenPos{x: i, y: j, val: '0'})
			}
		}
	}
	return res
}
