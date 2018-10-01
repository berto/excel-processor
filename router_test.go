package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testBody = `
This is a multi-part message.  This line is ignored.
--test
Header1: value1
HEADER2: value2
file: baz

My value
The end.
--test
file: bigsection

never read data
--test--
ok
`

func TestRoutes(t *testing.T) {
	r := createRouter()

	tt := []struct {
		name        string
		method      string
		uri         string
		contentType string
		status      int
		payload     []byte
	}{
		{"get index page", "GET", "/", "text/html; charset=utf-8", 200, nil},
		{"post excel form", "POST", "/ship", "multipart/form-data; boundary=test", 500, []byte(testBody)},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var payload io.Reader
			if tc.payload != nil {
				payload = bytes.NewBuffer(tc.payload)
			}
			req, err := http.NewRequest(tc.method, tc.uri, payload)
			req.Header.Set("Content-Type", tc.contentType)

			if err != nil {
				t.Errorf("Get failed with error %d.", err)
			}

			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			if resp.Code != tc.status {
				t.Errorf("/ failed with incorrect status code: %v", resp.Code)
			}

		})
	}
}
