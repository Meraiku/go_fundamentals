package tasks

import "golang.org/x/exp/constraints"

func Sum[T constraints.Ordered](nums ...T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}
