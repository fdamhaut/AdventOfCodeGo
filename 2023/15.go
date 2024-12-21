package main

import (
	"fmt"
	"strings"
)

type ll struct {
	next  *ll
	name  string
	focus int
}

func (list *ll) remove(name string) {
	for list.next != nil {
		if list.next.name == name {
			list.next = list.next.next
			break
		}
		list = list.next
	}
}

func (list *ll) put(name string, focus int) {
	for list.next != nil {
		if list.next.name == name {
			list.next.focus = focus
			return
		}
		list = list.next
	}
	list.next = &ll{next: nil, name: name, focus: focus}
}

func (list *ll) value() int {
	r, n := 0, 1
	for list.next != nil {
		list = list.next
		r += n * list.focus
		n += 1
	}
	return r
}

func fifteen() {
	lines, _ := toLineArray("15.in")
	ops := strings.Split(lines[0], ",")

	a, b := 0, 0

	boxes := make([]ll, 256)

	for n := range 256 {
		boxes[n] = ll{next: nil, name: "", focus: 0}
	}

	for _, op := range ops {
		a += hash15(op)
		do15(op, boxes)
	}

	for n := range 256 {
		b += (n + 1) * boxes[n].value()
	}

	fmt.Println(a, b)

}

func hash15(s string) int {
	r := 0
	for _, v := range []rune(s) {
		r += int(v)
		r *= 17
		r %= 256
	}
	return r
}

func do15(op string, boxes []ll) {
	if op[len(op)-1] == '-' {
		name := op[:len(op)-1]
		boxes[hash15(name)].remove(name)
	} else {
		name := op[:len(op)-2]
		focus := int(op[len(op)-1]) - 48
		boxes[hash15(name)].put(name, focus)
	}
}
