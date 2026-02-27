package utils

import "fmt"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("cannot divide by 0")
	}
	return a / b, nil
}
