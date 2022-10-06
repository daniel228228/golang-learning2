package solver

import (
	"fmt"

	helper "github.com/daniel228228/golang-learning-helper"
)

func check(max, value, precision float64) {
	Argument(max, value, precision)
}

func Table(from, max, step, precision float64) {
	for i := from; i < max; i += step {
		value := helper.GetValueByX(i, precision)
		fmt.Printf("x=%f,f(x)=%f\n", i, value)
		check(max, value, precision)
	}
}
