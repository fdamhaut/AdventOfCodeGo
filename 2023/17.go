package main

import (
	"fmt"
)

func seventeen() {
	lines, _ := toLineArray("17.in")
	a := solve17(lines, 0, 3)
	// b := solve17(lines, 3, 7)
	fmt.Println(a)
}

type SN struct {
	next  *SN
	x     int
	y     int
	dir   int
	count int
	cost  int
}

func (head *SN) add(x int, y int, dir int, count int, cost int) *SN {
	if head == nil || head.cost > cost {
		return &SN{x: x, y: y, dir: dir, cost: cost, count: count, next: head}
	}
	n := head
	for n.next != nil && n.next.cost < cost {
		n = n.next
	}
	n.next = &SN{x: x, y: y, dir: dir, cost: cost, count: count, next: n.next}
	return head
}

func solve17(lines []string, from int, to int) int {
	LEFT, RIGHT, UP, DOWN := -1, 1, -len(lines[0]), len(lines[0])

	queue := &SN{x: 0, y: 0, dir: 0, cost: int(lines[0][0]) - 48}

	for queue != nil {
		node := queue
		// fmt.Println(node, &node)
		queue = queue.next
		if node.x == len(lines)-1 && node.y == len(lines[0])-1 {
			return node.cost
		}

		var dir []int
		if node.count < from {
			dir = []int{node.dir}
		} else {
			dir = []int{LEFT, RIGHT, UP, DOWN}
		}

		for _, d := range dir {
			count := 0
			if d == node.dir {
				count = node.count + 1
			}
			if count > to {
				continue
			}
			var x, y int
			if d == UP && node.x-1 > 0 {
				x, y = node.x-1, node.y
				queue = queue.add(x, y, d, count, node.cost+int(lines[x][y])-48)
			} else if d == DOWN && node.x+1 < len(lines) {
				x, y = node.x+1, node.y
				queue = queue.add(x, y, d, count, node.cost+int(lines[x][y])-48)
			} else if d == LEFT && node.y-1 > 0 {
				x, y = node.x, node.y-1
				queue = queue.add(x, y, d, count, node.cost+int(lines[x][y])-48)
			} else if d == RIGHT && node.y+1 < len(lines[0]) {
				x, y = node.x+1, node.y+1
				queue = queue.add(x, y, d, count, node.cost+int(lines[x][y])-48)
			}
			// fmt.Println(node, &node)

		}
	}

	return 0
}
