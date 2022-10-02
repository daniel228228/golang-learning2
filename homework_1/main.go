package main

import (
	"homework_1/solver"
)

const (
	precision float64 = 0.0001
	step      float64 = 0.5
	max       float64 = 5
)

func main() {
	solver.Table(0, max, step, precision)
}
