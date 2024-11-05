package tasks

import "fmt"

var (
	ErrDivisionByZero = fmt.Errorf("division by zero")
)

func Division(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, ErrDivisionByZero
	}
	return dividend / divisor, nil
}
