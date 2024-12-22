package main

import (
	"fmt"
	"strconv"
	"strings"
)

func three() {
	data := toLineArray("3.in")
	mem := strings.Join(data, "")
	a := threeParser(mem, threeMethodsA())
	b := threeParser(mem, threeMethodsB())
	fmt.Println(a, b)
}

type State struct {
	result  int
	disable bool
}

type Method struct {
	name string
	exec func(string, State) (State, error)
}

func (m *Method) parse(line string, state State) (State, int) {
	mlen := len(m.name)
	if mlen > len(line) || m.name != line[:mlen] {
		return state, 0
	}

	end := strings.Index(line[mlen:], ")")
	if end <= 0 {
		return state, 0
	}

	result, err := m.exec(line[mlen+1:mlen+end], state)
	if err != nil {
		return state, 0
	}
	return result, mlen + end
}

func threeParser(line string, methods []Method) int {
	state := State{
		result:  0,
		disable: false,
	}
	adv := 0
	for idx := 0; idx < len(line); idx++ {
		for _, m := range methods {
			state, adv = m.parse(line[idx:], state)
			if adv > 0 {
				idx += adv
				break
			}
		}
	}

	return state.result
}

func threeMethodsA() []Method {
	return []Method{
		{
			name: "mul",
			exec: threeMulInt,
		},
	}
}

func threeMethodsB() []Method {
	return []Method{
		{
			name: "mul",
			exec: threeMulInt,
		},
		{
			name: "don't",
			exec: threeDont,
		},
		{
			name: "do",
			exec: threeDo,
		},
	}
}

func threeMulInt(args string, state State) (State, error) {
	strs := strings.Split(args, ",")
	ints := make([]int, 0, len(strs))
	for _, v := range strs {
		i, err := strconv.Atoi(v)
		if err != nil {
			return state, err
		}
		ints = append(ints, i)
	}
	if !state.disable {
		state.result += ints[0] * ints[1]
	}
	return state, nil
}

func threeDo(args string, state State) (State, error) {
	state.disable = false
	return state, nil
}

func threeDont(args string, state State) (State, error) {
	state.disable = true
	return state, nil
}
