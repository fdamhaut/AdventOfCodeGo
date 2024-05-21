package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func six() {
	data, _ := toLineArray("6.in")

	time := toIntArray(data[0])
	dist := toIntArray(data[1])

	a := 1
	for index := range len(time) {
		t, d := time[index], dist[index]
		b := (float64(t) - math.Sqrt(float64(t*t-4*d))) / 2
		res := (t + 1) - 2*int(math.Ceil(b+1e-5))
		a *= res
	}

	timeb := strings.Replace(data[0][strings.Index(data[0], ":")+1:], " ", "", -1)
	distb := strings.Replace(data[1][strings.Index(data[1], ":")+1:], " ", "", -1)

	t, _ := strconv.Atoi(timeb)
	d, _ := strconv.Atoi(distb)

	mid := (float64(t) - math.Sqrt(float64(t*t-4*d))) / 2
	b := (t + 1) - 2*int(math.Ceil(mid+1e-5))

	fmt.Println(a, b)
}
