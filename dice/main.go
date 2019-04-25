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

	for i := 0; i < n; i++ {
		sums[m.Intn(6)+m.Intn(6)+2]++ // 0..5 + 0..5 +2
	}

	for i := 2; i < 13; i++ {
		index[i] = i
		percents[i] = sums[i] * 100 / n
	}
	output.PrintRow(index[2:])
	output.PrintRow(sums[2:])
	output.PrintRow(percents[2:])
}
