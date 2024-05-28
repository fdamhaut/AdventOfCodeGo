package main

import "fmt"

func eight() {
	lines, _ := toLineArray("8.in")

	operations := op2int(lines[0])
	paths := l2map8(lines[2:])

	a := solve8A(operations, paths)
	b := solve8B(operations, paths)

	fmt.Println(a, b)
}

func op2int(op string) []int {
	res := make([]int, 0, len(op))

	for _, v := range op {
		if v == 'L' {
			res = append(res, 0)
		} else if v == 'R' {
			res = append(res, 1)
		}
	}

	return res
}

func l2map8(ls []string) map[string][]string {
	res := make(map[string][]string, len(ls))

	for _, line := range ls {
		res[line[0:3]] = []string{line[7:10], line[12:15]}
	}

	return res
}

func solve8A(op []int, paths map[string][]string) int {

	pos := "AAA"
	op_idx := 0
	res := 0

	for pos != "ZZZ" {
		pos = paths[pos][op[op_idx]]
		op_idx = (op_idx + 1) % len(op)
		res++
	}

	return res
}

func solve8B(op []int, paths map[string][]string) int {

	// This code assumes that all loops start a "start" and end at the first pos[2] == 'Z' location
	// it's true in the context of this exercise but not always
	// Code would be much harder otherwise (smallest occurance in all ['startup'+n*loop for n in N])

	loops := make([]int, 0, len(paths))

	for k := range paths {
		if k[2] == 'A' {
			loops = append(loops, getLoop(k, op, paths))
		}
	}

	res := loops[0]

	for _, v := range loops {
		res = ppcm(res, v)
	}

	return res
}

func ppcm(a int, b int) int {
	return a * b / pgcd(a, b)
}

func pgcd(a int, b int) int {
	if b > a {
		return pgcd(b, a)
	}
	if a%b == 0 {
		return b
	}
	return pgcd(b, a%b)
}

func getLoop(start string, op []int, paths map[string][]string) int {

	pos := paths[start][op[0]]
	op_idx := 1
	res := 1

	seen := make(map[string]bool, len(paths))
	seen[pos] = true

	for pos[2] != 'Z' {
		pos = paths[pos][op[op_idx]]
		op_idx = (op_idx + 1) % len(op)
		res++
	}

	return res

}
