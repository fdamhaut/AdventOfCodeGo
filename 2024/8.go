package main

import "fmt"

func eight() {
	carte := toLineArray("8.in")
	a := eightSolveA(carte)
	b := eightSolveB(carte)
	fmt.Println(a, b)
}

type eightPos struct {
	char rune
	x    int
	y    int
}

func (p1 *eightPos) antinode(p2 eightPos) []eightPos {
	if p1.char != p2.char {
		return []eightPos{}
	}
	return []eightPos{
		{x: p1.x*2 - p2.x, y: p1.y*2 - p2.y, char: p1.char},
		{x: p2.x*2 - p1.x, y: p2.y*2 - p1.y, char: p1.char},
	}
}

func (p1 *eightPos) antinodeB(p2 eightPos) []eightPos {
	if p1.char != p2.char {
		return []eightPos{}
	}
	eightPoses := make([]eightPos, 0, 200)
	for i := range 50 {
		eightPoses = append(eightPoses, eightPos{x: p1.x + i*(p1.x-p2.x), y: p1.y + i*(p1.y-p2.y), char: p1.char})
		eightPoses = append(eightPoses, eightPos{x: p2.x + i*(p2.x-p1.x), y: p2.y + i*(p2.y-p1.y), char: p1.char})
	}
	return eightPoses
}

func (p *eightPos) hash(carte []string) int {
	return p.x*len(carte[0]) + p.y
}

func eightSolveA(carte []string) int {
	poses := eightMakePos(carte)
	res := make(map[int]bool)
	for key := range poses {
		for i, pos := range poses[key][:len(poses[key])-1] {
			for _, pos2 := range poses[key][i+1:] {
				antinodes := pos.antinode(pos2)
				for _, a := range antinodes {
					if a.x < 0 || a.y < 0 || a.x >= len(carte) || a.y >= len(carte[0]) {
						continue
					}
					res[a.hash(carte)] = true
				}
			}
		}
	}
	return len(res)
}

func eightSolveB(carte []string) int {
	poses := eightMakePos(carte)
	res := make(map[int]bool)
	for key := range poses {
		for i, pos := range poses[key][:len(poses[key])-1] {
			for _, pos2 := range poses[key][i+1:] {
				antinodes := pos.antinodeB(pos2)
				for _, a := range antinodes {
					if a.x < 0 || a.y < 0 || a.x >= len(carte) || a.y >= len(carte[0]) {
						continue
					}
					res[a.hash(carte)] = true
				}
			}
		}
	}
	return len(res)
}

func eightMakePos(carte []string) map[rune][]eightPos {
	res := make(map[rune][]eightPos)
	for i, line := range carte {
		for j, c := range line {
			if c != '.' {
				res[c] = append(res[c], eightPos{
					x:    i,
					y:    j,
					char: c,
				})
			}
		}
	}
	return res
}
