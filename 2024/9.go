package main

import "fmt"

func nine() {
	data := toLineArray("9.in")
	a := nineSolveA(data[0])
	fmt.Println(a)
}

const R2I = 48

func nineSolveA(data string) int {
	idx := 0
	reverse := len(data) - 1
	reverse_left := int(data[reverse]) - 48

	pos := 0
	res := 0

	for idx < reverse {
		full := int(data[idx]) - 48
		empty := int(data[idx+1]) - 48

		res += (idx / 2) * (full) * (2*pos + full - 1) / 2
		fmt.Println("F", pos, idx/2, full, (idx/2)*(full)*(2*pos+full-1)/2)

		pos += full

		for empty > 0 {
			amount := min(reverse_left, empty)
			res += (reverse / 2) * (amount) * (2*pos + amount - 1) / 2
			fmt.Println("R", pos, reverse/2, amount, (reverse/2)*(amount)*(2*pos+amount-1)/2)
			reverse_left -= amount
			if reverse_left <= 0 {
				reverse -= 2
				reverse_left = int(data[reverse]) - 48
			}
			empty -= amount
			pos += amount
		}
		idx += 2
		if idx == reverse {
			fmt.Println("F", pos, reverse/2, reverse_left, (reverse/2)*(reverse_left)*(2*pos+reverse_left-1)/2)
			res += (reverse / 2) * (reverse_left) * (2*pos + reverse_left - 1) / 2
		}
	}

	return res
}
