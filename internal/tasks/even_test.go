package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEven(t *testing.T) {

	tt := []struct {
		testName string
		num      int
		expected bool
	}{
		{
			testName: "Test with even number",
			num:      2,
			expected: true,
		},
		{
			testName: "Test with odd number",
			num:      3,
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(
			tc.testName,
			func(t *testing.T) {
				assert.Equal(t, tc.expected, IsEven(tc.num))
			},
		)
	}

}
