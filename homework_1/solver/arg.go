package solver

import (
	"fmt"

	helper "github.com/daniel228228/golang-learning-helper"
)

func Argument(max, value, precision float64) {
	x, iter := helper.GetXByValue(0, max, value, precision)

	fmt.Printf("check: for f(x) = %f x = %f (%d iterations)\n", value, x, iter)
}
