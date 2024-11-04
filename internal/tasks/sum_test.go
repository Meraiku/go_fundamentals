package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumInt(t *testing.T) {

	tt := []struct {
		testName string
		nums     []int
		expected int
	}{
		{
			testName: "Test with empty slice",
			nums:     []int{},
			expected: 0,
		},
		{
			testName: "Test with nil slice",
			nums:     nil,
			expected: 0,
		},
		{
			testName: "Test with single value",
			nums:     []int{1},
			expected: 1,
		},
		{
			testName: "Test with ints slice",
			nums:     []int{1, 2, 3},
			expected: 6,
		},
		{
			testName: "Test with negative ints slice",
			nums:     []int{-1, -2, -3},
			expected: -6,
		},
	}

	for _, tc := range tt {
		t.Run(
			tc.testName,
			func(t *testing.T) {
				assert.Equal(t, tc.expected, Sum(tc.nums...))
			},
		)
	}
}
