package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetPerson(t *testing.T) {

	tt := []struct {
		testName string
		name     string
		expected string
	}{

		{
			testName: "Test with name",
			name:     "Nikita",
			expected: "Hello, Nikita!",
		},
		{
			testName: "Test with empty name",
			name:     "",
			expected: "Hello, World!",
		},
	}

	for _, tc := range tt {
		t.Run(
			tc.testName,
			func(t *testing.T) {
				assert.Equal(t, tc.expected, GreetPerson(tc.name))
			},
		)
	}

}
