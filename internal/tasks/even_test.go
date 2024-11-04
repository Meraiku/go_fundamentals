package tasks

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestEvenWithQuick(t *testing.T) {

	f := func(num int) bool {
		checked := IsEven(num)
		return checked == (num%2 == 0)
	}

	err := quick.Check(f, nil)

	require.NoError(t, err)
}
