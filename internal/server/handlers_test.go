package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootHandler(t *testing.T) {

	tt := []struct {
		testName       string
		method         string
		path           string
		statusCode     int
		expectedInBody string
	}{
		{
			testName:       "Test root handler",
			method:         http.MethodGet,
			path:           "/",
			statusCode:     http.StatusOK,
			expectedInBody: "Hello world",
		},
		{
			testName:       "Test root handler with wrong method",
			method:         http.MethodPost,
			path:           "/",
			statusCode:     http.StatusMethodNotAllowed,
			expectedInBody: "Method not allowed",
		},
	}

	for _, tc := range tt {
		t.Run(
			tc.testName,
			func(t *testing.T) {

				req := httptest.NewRequest(tc.method, tc.path, nil)
				w := httptest.NewRecorder()

				handleRoot(w, req)

				resp := w.Result()
				defer resp.Body.Close()

				assert.Equal(t, tc.statusCode, resp.StatusCode)
				assert.Contains(t, w.Body.String(), tc.expectedInBody)

			},
		)
	}

}
