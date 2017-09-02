package myhandler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var targets = []string{
	"http://localhost:8080/result",
	"http://localhost:8080/add",
	"http://localhost:8080/dec",
	"http://localhost:8080/set?value=20",
	"http://localhost:8080/set?value=monkey",
}

//type testReq struct {
//m string
//t []string
//}

var tests = []string{
	http.MethodGet,
	http.MethodPost,
}

//var tests = []http.Request{
//httptest.NewRequest(http.MethodGet, "http://localhost:8080/result", nil),
//httptest.NewRequest(http.MethodPost, "http://localhost:8080/add", nil),
//httptest.NewRequest(http.MethodPost, "http://localhost:8080/dec", nil),
//httptest.NewRequest(http.MethodPost, "http://localhost:8080/set?value=20", nil)
//}

func TestHandlers(t *testing.T) {
	for _, m := range tests {
		for _, target := range targets {
			req := httptest.NewRequest(m, target, nil)

			if strings.Contains(target, "add") {
				wAdd := httptest.NewRecorder()
				AddHandler(wAdd, req)
				respAdd := wAdd.Result()
				if respAdd.StatusCode != http.StatusOK && respAdd.StatusCode != http.StatusMethodNotAllowed {
					t.Error(
						"For", req,
						"expected", http.StatusOK, "or", http.StatusMethodNotAllowed,
						"got", respAdd.StatusCode,
					)
				}
			}

			if strings.Contains(target, "dec") {
				wDec := httptest.NewRecorder()
				DecHandler(wDec, req)
				respDec := wDec.Result()
				if respDec.StatusCode != http.StatusOK && respDec.StatusCode != http.StatusMethodNotAllowed {
					t.Error(
						"For", req,
						"expected", http.StatusOK, "or", http.StatusMethodNotAllowed,
						"got", respDec.StatusCode,
					)
				}
			}

			wRes := httptest.NewRecorder()
			ResHandler(wRes, req)
			respRes := wRes.Result()
			if respRes.StatusCode != http.StatusOK && respRes.StatusCode != http.StatusMethodNotAllowed {
				t.Error(
					"For", req,
					"expected", http.StatusOK, "or", http.StatusMethodNotAllowed,
					"got", respRes.StatusCode,
				)
			}

			wSet := httptest.NewRecorder()
			SetHandler(wSet, req)
			respSet := wSet.Result()
			if respSet.StatusCode != http.StatusOK && respSet.StatusCode != http.StatusMethodNotAllowed && respSet.StatusCode != http.StatusBadRequest {
				t.Error(
					"For", req,
					"expected", http.StatusOK, "or", http.StatusMethodNotAllowed, "or",
					http.StatusBadRequest,
					"got", respSet.StatusCode,
				)
			}
		}
	}
}

//func TestAddHandler(t *testing.T)

//func TestDecHandler(t *testing.T)

//func  TestResHandler(t *testing.T)

//func TestSetHandler(t *testing.T)
