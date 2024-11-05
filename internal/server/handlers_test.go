package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mock_client "github.com/meraiku/go_fundamentals/internal/client/mocks"
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

func TestSearchHandler(t *testing.T) {

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	client := mock_client.NewMockClient(ctrl)

	api := New(client)

	tt := []struct {
		testName       string
		path           string
		query          string
		prepare        func(path, query, out string)
		statusCode     int
		expectedInBody string
	}{
		{
			testName: "Test search handler",
			path:     "/search?q=World+Wide",
			prepare: func(path, query, out string) {

				client.EXPECT().Get(fmt.Sprintf("%s?%s", path, query)).Return([]byte(out), nil)
			},
			statusCode:     http.StatusOK,
			expectedInBody: "World Wide",
		},
		{
			testName: "Test search without query parameters",
			path:     "/search",
			prepare: func(path, query, out string) {

				client.EXPECT().Get(fmt.Sprintf("%s?q=%s", path, query)).Return([]byte(out), nil)
			},
			statusCode:     http.StatusOK,
			expectedInBody: "",
		},
		{
			testName: "Test search with fail in client",
			path:     "/search",
			prepare: func(path, query, out string) {

				client.EXPECT().Get(fmt.Sprintf("%s?q=%s", path, query)).Return(nil, fmt.Errorf("Internal server error"))
			},
			statusCode:     http.StatusInternalServerError,
			expectedInBody: "Internal server error",
		},
	}

	for _, tc := range tt {
		t.Run(
			tc.testName,
			func(t *testing.T) {

				req := httptest.NewRequest(http.MethodGet, tc.path, nil)
				w := httptest.NewRecorder()

				tc.prepare(req.URL.Path, req.URL.RawQuery, tc.expectedInBody)

				api.handleSearch(w, req)

				resp := w.Result()
				defer resp.Body.Close()

				assert.Equal(t, tc.statusCode, resp.StatusCode)
				assert.Contains(t, w.Body.String(), tc.expectedInBody)
			},
		)
	}

}
