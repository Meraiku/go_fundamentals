package google

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoogleClient(t *testing.T) {

	c := New()

	tt := []struct {
		testName   string
		path       string
		expected   string
		statusCode int
	}{
		{
			testName: "Test root path",
			path:     "/",
			expected: "",
		},
		{
			testName: "Test search without query",
			path:     "/search",
			expected: "",
		},
		{
			testName: "Test search with query",
			path:     "/search?q=hello",
			expected: "hello",
		},
		{
			testName:   "Test with random path",
			path:       "/doqnfoiqnq",
			expected:   "",
			statusCode: http.StatusNotFound,
		},
	}

	for _, tc := range tt {
		t.Run(
			tc.testName,
			func(t *testing.T) {
				body, err := c.Get(tc.path)
				if err != nil {
					assert.Contains(t, err.Error(), fmt.Sprintf("%d", tc.statusCode))
					return
				}

				assert.Contains(t, string(body), tc.expected)
			},
		)
	}
}
