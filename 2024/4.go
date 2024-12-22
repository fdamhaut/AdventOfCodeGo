package main

import "fmt"

func four() {
	data := toLineArray("4.in")
	X := fourMask(data, 'X')
	M := fourMask(data, 'M')
	A := fourMask(data, 'A')
	S := fourMask(data, 'S')
	a := fourA(X, M, A, S)
	b := fourB(X, M, A, S)
	fmt.Println(a, b)
}

func fourA(X [][]bool, M [][]bool, A [][]bool, S [][]bool) int {
	res := 0
	// H
	res += fourAND([][][]bool{
		fourSubmask(X, 0, 3, 0, 0),
		fourSubmask(M, 1, 2, 0, 0),
		fourSubmask(A, 2, 1, 0, 0),
		fourSubmask(S, 3, 0, 0, 0),
	})
	// Hi
	res += fourAND([][][]bool{
		fourSubmask(X, 3, 0, 0, 0),
		fourSubmask(M, 2, 1, 0, 0),
		fourSubmask(A, 1, 2, 0, 0),
		fourSubmask(S, 0, 3, 0, 0),
	})
	// V
	res += fourAND([][][]bool{
		fourSubmask(X, 0, 0, 3, 0),
		fourSubmask(M, 0, 0, 2, 1),
		fourSubmask(A, 0, 0, 1, 2),
		fourSubmask(S, 0, 0, 0, 3),
	})
	// Vi
	res += fourAND([][][]bool{
		fourSubmask(X, 0, 0, 0, 3),
		fourSubmask(M, 0, 0, 1, 2),
		fourSubmask(A, 0, 0, 2, 1),
		fourSubmask(S, 0, 0, 3, 0),
	})
	// \
	res += fourAND([][][]bool{
		fourSubmask(X, 0, 3, 0, 3),
		fourSubmask(M, 1, 2, 1, 2),
		fourSubmask(A, 2, 1, 2, 1),
		fourSubmask(S, 3, 0, 3, 0),
	})
	// \i
	res += fourAND([][][]bool{
		fourSubmask(X, 3, 0, 3, 0),
		fourSubmask(M, 2, 1, 2, 1),
		fourSubmask(A, 1, 2, 1, 2),
		fourSubmask(S, 0, 3, 0, 3),
	})
	// /i
	res += fourAND([][][]bool{
		fourSubmask(X, 0, 3, 3, 0),
		fourSubmask(M, 1, 2, 2, 1),
		fourSubmask(A, 2, 1, 1, 2),
		fourSubmask(S, 3, 0, 0, 3),
	})
	// /
	res += fourAND([][][]bool{
		fourSubmask(X, 3, 0, 0, 3),
		fourSubmask(M, 2, 1, 1, 2),
		fourSubmask(A, 1, 2, 2, 1),
		fourSubmask(S, 0, 3, 3, 0),
	})
	return res
}

func fourB(X [][]bool, M [][]bool, A [][]bool, S [][]bool) int {
	res := 0
	// M.M
	// .A.
	// S.S
	res += fourAND([][][]bool{
		fourSubmask(M, 0, 2, 0, 2),
		fourSubmask(M, 2, 0, 0, 2),
		fourSubmask(A, 1, 1, 1, 1),
		fourSubmask(S, 0, 2, 2, 0),
		fourSubmask(S, 2, 0, 2, 0),
	})
	// M.S
	// .A.
	// M.S
	res += fourAND([][][]bool{
		fourSubmask(M, 0, 2, 0, 2),
		fourSubmask(M, 0, 2, 2, 0),
		fourSubmask(A, 1, 1, 1, 1),
		fourSubmask(S, 2, 0, 0, 2),
		fourSubmask(S, 2, 0, 2, 0),
	})
	// S.S
	// .A.
	// M.M
	res += fourAND([][][]bool{
		fourSubmask(S, 0, 2, 0, 2),
		fourSubmask(S, 2, 0, 0, 2),
		fourSubmask(A, 1, 1, 1, 1),
		fourSubmask(M, 0, 2, 2, 0),
		fourSubmask(M, 2, 0, 2, 0),
	})
	// S.M
	// .A.
	// S.M
	res += fourAND([][][]bool{
		fourSubmask(S, 0, 2, 0, 2),
		fourSubmask(S, 0, 2, 2, 0),
		fourSubmask(A, 1, 1, 1, 1),
		fourSubmask(M, 2, 0, 0, 2),
		fourSubmask(M, 2, 0, 2, 0),
	})
	return res
}

func fourMask(matrix []string, letter rune) [][]bool {
	res := make([][]bool, len(matrix))
	for i, line := range matrix {
		res[i] = make([]bool, len(line))
		for j, r := range line {
			res[i][j] = r == letter
		}
	}
	return res
}

func fourSubmask(mask [][]bool, left int, right int, top int, bot int) [][]bool {
	res := make([][]bool, len(mask)-top-bot)
	for i := range res {
		res[i] = mask[i+top][left : len(mask[i])-right]
	}
	return res
}

func fourAND(masks [][][]bool) int {
	res := 0
	for i := range masks[0] {
		for j := range masks[0][0] {
			pos := true
			for m := range masks {
				pos = pos && masks[m][i][j]
			}
			if pos {
				res += 1
			}
		}
	}
	return res
}
