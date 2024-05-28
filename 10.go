package main

import (
	"fmt"
	"strings"
)

func ten() {
	lines, _ := toLineArray("10.in")

	size := len(lines[0])
	maze := strings.Join(lines, "")

	up, down := -size, size
	left, right := -1, 1

	moves := map[string][]int{
		"-": []int{left, right},
		"|": []int{up, down},
		"7": []int{left, down},
		"L": []int{up, right},
		"F": []int{right, down},
		"J": []int{up, left},
		".": []int{},
		"S": []int{},
	}

	start := strings.Index(maze, "S")

	todo := make([]int, 0, len(maze))
	done := make(map[int]bool, len(maze))
	done[start] = true

	if strings.Contains("-LF", maze[start+left:start+left+1]) {
		todo = append(todo, start+left)
	}
	if strings.Contains("|7F", maze[start+up:start+up+1]) {
		todo = append(todo, start+left)
	}
	if strings.Contains("|JL", maze[start+down:start+down+1]) {
		todo = append(todo, start+left)
	}
	if strings.Contains("-7J", maze[start+right:start+right+1]) {
		todo = append(todo, start+left)
	}

	todo_idx := 0
	for todo_idx < len(todo) {
		pos := todo[todo_idx]
		todo_idx++
		if done[pos] {
			continue
		}
		done[pos] = true
		for _, move := range moves[maze[pos:pos+1]] {
			todo = append(todo, pos+move)
		}
	}

	// TODO B flemme
	fmt.Println(len(done)/2, 0)
}
