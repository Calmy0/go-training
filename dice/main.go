package main

import (
	"dice/output"
	m "math/rand"
)

const (
	n = 10000
)

func main() {
	index := make([]int, 13)    // 0..12
	sums := make([]int, 13)     // 0..12
	percents := make([]int, 13) // 0..12

	total := 0

	var dice1, dice2 int

	for i := 0; i < n; i++ {
		dice1 = m.Intn(6)
		dice2 = m.Intn(6)
		sums[dice1+dice2+2]++
	}

	for i := 2; i < 13; i++ {
		total += sums[i]
	}
	for i := 2; i < 13; i++ {
		index[i] = i
		percents[i] = sums[i] * 100 / total
	}
	output.PrintRow(index[2:])
	output.PrintRow(sums[2:])
	output.PrintRow(percents[2:])
}
