package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQueueRoutes(t *testing.T) {
	r := createRouter()

	tt := []struct {
		name    string
		method  string
		uri     string
		payload []byte
	}{
		{"get index page", "GET", "/", nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var payload io.Reader
			if tc.payload != nil {
				payload = bytes.NewBuffer(tc.payload)
			}
			req, err := http.NewRequest(tc.method, tc.uri, payload)
			if err != nil {
				t.Errorf("Get failed with error %d.", err)
			}

			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			if resp.Code != 200 {
				t.Errorf("/ failed with error code %d.", resp.Code)
			}

			if resp.Header().Get("Content-Type") != "text/html; charset=utf-8" {
				t.Errorf("/ failed with incorrect headers: %v", resp.Header().Get("Content-Type"))
			}
		})
	}
}
