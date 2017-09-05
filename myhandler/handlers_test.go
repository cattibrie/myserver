package myhandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//var targets = []string{
//"http://localhost:8080/result",
//"http://localhost:8080/add",
//"http://localhost:8080/dec",
//"http://localhost:8080/set?value=20",
//"http://localhost:8080/set?value=monkey",
//}

var tests = []string{
	http.MethodGet,
	http.MethodPost,
}

//func addTest

func TestHandlers(t *testing.T) {
	var v = Value{V: 0}
	var i Operat = &v
	var h = Handlers{H: i}
	for _, m := range tests {
		//for _, target := range targets {
		req1 := httptest.NewRequest(m, "http://localhost:8080/add", nil)
		nv1 := v.V + 1
		wAdd := httptest.NewRecorder()
		h.AddHandler(wAdd, req1)
		respAdd := wAdd.Result()
		if respAdd.StatusCode != http.StatusOK && respAdd.StatusCode != http.StatusMethodNotAllowed {
			t.Error(
				"For", req1,
				"expected", http.StatusOK, "or", http.StatusMethodNotAllowed,
				"got", respAdd.StatusCode,
			)
		} else if respAdd.StatusCode == http.StatusOK && v.V != nv1 {
			t.Error(
				"For value",
				"expected", nv1,
				"got", v.V,
			)
		}

		req2 := httptest.NewRequest(m, "http://localhost:8080/dec", nil)
		nv2 := v.V - 1
		wDec := httptest.NewRecorder()
		h.DecHandler(wDec, req2)
		respDec := wDec.Result()
		if respDec.StatusCode != http.StatusOK && respDec.StatusCode != http.StatusMethodNotAllowed {
			t.Error(
				"For", req2,
				"expected", http.StatusOK, "or", http.StatusMethodNotAllowed,
				"got", respDec.StatusCode,
			)
		} else if respDec.StatusCode == http.StatusOK && v.V != nv2 {
			t.Error(
				"For value",
				"expected", nv2,
				"got", v.V,
			)
		}

		req3 := httptest.NewRequest(m, "http://localhost:8080/result", nil)
		nv3 := v.V
		wRes := httptest.NewRecorder()
		h.ResHandler(wRes, req3)
		respRes := wRes.Result()
		if respRes.StatusCode != http.StatusOK && respRes.StatusCode != http.StatusMethodNotAllowed {
			t.Error(
				"For", req3,
				"expected", http.StatusOK, "or", http.StatusMethodNotAllowed,
				"got", respRes.StatusCode,
			)
		} else if respRes.StatusCode == http.StatusOK && v.V != nv3 {
			t.Error(
				"For value",
				"expected", nv3,
				"got", v.V,
			)
		}

		req4 := httptest.NewRequest(m, "http://localhost:8080/set?value=20", nil)
		var nv4 int64 = 20
		wSet := httptest.NewRecorder()
		h.SetHandler(wSet, req4)
		respSet := wSet.Result()
		if respSet.StatusCode != http.StatusOK && respSet.StatusCode != http.StatusMethodNotAllowed && respSet.StatusCode != http.StatusBadRequest {
			t.Error(
				"For", req4,
				"expected", http.StatusOK, "or", http.StatusMethodNotAllowed, "or",
				http.StatusBadRequest,
				"got", respSet.StatusCode,
			)
		} else if respSet.StatusCode == http.StatusOK && v.V != nv4 {
			t.Error(
				"For value",
				"expected", nv4,
				"got", v.V,
			)
		}

	}
}

//func TestAddHandler(t *testing.T)

//func TestDecHandler(t *testing.T)

//func  TestResHandler(t *testing.T)

//func TestSetHandler(t *testing.T)
