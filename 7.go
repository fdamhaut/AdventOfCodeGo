package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func seven() {
	data, err := toLineArray("7.in")
	check(err)

	a := solve7(data, getHandA)
	b := solve7(data, getHandB)
	fmt.Println(a, b)
}

func solve7(data []string, getHand func(string) Hand) int {
	cards := toCardArray(data, getHand)
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].typ == cards[j].typ {
			return cards[i].label < cards[j].label
		} else {
			return cards[i].typ < cards[j].typ
		}
	})

	a := 0
	for n, hand := range cards {
		a += (n + 1) * hand.bid
	}

	return a
}

type Hand struct {
	label string
	bid   int
	typ   int
}

func getHandA(line string) Hand {
	line_split := strings.Split(line, " ")

	bid, _ := strconv.Atoi(line_split[1])

	hand := strings.Replace(line_split[0], "A", "e", -1)
	hand = strings.Replace(hand, "K", "d", -1)
	hand = strings.Replace(hand, "Q", "c", -1)
	hand = strings.Replace(hand, "J", "b", -1)
	hand = strings.Replace(hand, "T", "a", -1)

	count := map[rune]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, '9': 0, '8': 0, '7': 0, '6': 0, '5': 0, '4': 0, '3': 0, '2': 0}

	for _, v := range []rune(hand) {
		count[v] += 1
	}

	cval := make([]int, 0, len(count))

	for _, value := range count {
		cval = append(cval, value)
	}

	slices.Sort(cval)

	if cval[12] == 5 {
		return Hand{label: hand, bid: bid, typ: 6}
	} else if cval[12] == 4 {
		return Hand{label: hand, bid: bid, typ: 5}
	} else if cval[12] == 3 && cval[11] == 2 {
		return Hand{label: hand, bid: bid, typ: 4}
	} else if cval[12] == 3 {
		return Hand{label: hand, bid: bid, typ: 3}
	} else if cval[12] == 2 && cval[11] == 2 {
		return Hand{label: hand, bid: bid, typ: 2}
	} else if cval[12] == 2 {
		return Hand{label: hand, bid: bid, typ: 1}
	} else {
		return Hand{label: hand, bid: bid, typ: 0}
	}
}

func getHandB(line string) Hand {
	line_split := strings.Split(line, " ")

	bid, _ := strconv.Atoi(line_split[1])

	hand := strings.Replace(line_split[0], "A", "e", -1)
	hand = strings.Replace(hand, "K", "d", -1)
	hand = strings.Replace(hand, "Q", "c", -1)
	hand = strings.Replace(hand, "J", "1", -1)
	hand = strings.Replace(hand, "T", "a", -1)

	count := map[rune]int{'a': 0, '1': 0, 'c': 0, 'd': 0, 'e': 0, '9': 0, '8': 0, '7': 0, '6': 0, '5': 0, '4': 0, '3': 0, '2': 0}

	for _, v := range []rune(hand) {
		count[v] += 1
	}

	cval := make([]int, 1, len(count))

	for k, value := range count {
		if k != '1' {
			cval = append(cval, value)
		}
	}

	slices.Sort(cval)

	if cval[12]+count['1'] == 5 {
		return Hand{label: hand, bid: bid, typ: 6}
	} else if cval[12]+count['1'] == 4 {
		return Hand{label: hand, bid: bid, typ: 5}
	} else if cval[12]+count['1'] == 3 && cval[11] == 2 {
		return Hand{label: hand, bid: bid, typ: 4}
	} else if cval[12]+count['1'] == 3 {
		return Hand{label: hand, bid: bid, typ: 3}
	} else if cval[12]+count['1'] == 2 && cval[11] == 2 {
		return Hand{label: hand, bid: bid, typ: 2}
	} else if cval[12]+count['1'] == 2 {
		return Hand{label: hand, bid: bid, typ: 1}
	} else {
		return Hand{label: hand, bid: bid, typ: 0}
	}
}

func toCardArray(ls []string, getHand func(string) Hand) []Hand {
	res := make([]Hand, len(ls))
	for n, l := range ls {
		res[n] = getHand(l)
	}
	return res
}
