package tasks

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDivision(t *testing.T) {

	tt := []struct {
		testName string
		dividend int
		divisor  int
		expected int
		err      error
	}{
		{
			testName: "Test with positive dividend and divisor",
			dividend: 10,
			divisor:  2,
			expected: 5,
		},
		{
			testName: "Test with negative dividend and positive divisor",
			dividend: -10,
			divisor:  2,
			expected: -5,
		},
		{
			testName: "Test with negative dividend and negative divisor",
			dividend: -10,
			divisor:  -2,
			expected: 5,
		},
		{
			testName: "Test with zero dividend and positive divisor",
			dividend: 0,
			divisor:  2,
			err:      ErrDivisionByZero,
		},
		{
			testName: "Test with zero dividend and zero divisor",
			dividend: 0,
			divisor:  0,
			err:      ErrDivisionByZero,
		},
		{
			testName: "Test with zero dividend and negative divisor",
			dividend: 0,
			divisor:  -2,
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(
			tc.testName,
			func(t *testing.T) {
				res, err := Division(tc.dividend, tc.divisor)

				switch err {
				case nil:
					assert.Equal(t, tc.expected, res)
				default:
					assert.Equal(t, tc.err, err)
				}

			},
		)
	}
}

func TestDivisionWithQuick(t *testing.T) {
	f := func(num1, num2 int) bool {
		checked, err := Division(num1, num2)

		switch err {
		case nil:
			return checked == (num1 / num2)
		case ErrDivisionByZero:
			return num2 == 0
		default:
			return false
		}
	}

	err := quick.Check(f, nil)

	require.NoError(t, err)
}
