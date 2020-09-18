package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	testCases := []struct {
		desc, in, out string
	}{
		{desc: "One", in: "sapan@gmail", out: "hello, sapan"},
		{desc: "Two", in: "other", out: "hello, stranger"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet,
				"http://localhost:8080/"+tC.in,
				nil)
			if err != nil {
				t.Fatalf("Request creation failes : %v", err)
			}
			resp := httptest.NewRecorder()
			handler(resp, req)

			if resp.Code != http.StatusOK {
				t.Errorf(tC.desc+"response Code : %v", resp.Code)
			}

			if !strings.Contains(resp.Body.String(), tC.out) {
				t.Errorf(tC.desc+" body [%v]", resp.Body.String())
			}
		})
	}
}
