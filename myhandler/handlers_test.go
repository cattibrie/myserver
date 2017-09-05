package myhandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func initials() (v *Value, h *Handlers) {
	v = &Value{V: 0}
	var i Operat = v
	h = &Handlers{H: i}
	return v, h
}

func checkhandlers(resp *http.Response, r *http.Request, t *testing.T, v *Value, vexp int64) {
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusMethodNotAllowed {
		t.Error("For", r,
			"expected", http.StatusOK, "or", http.StatusMethodNotAllowed,
			"got", resp.StatusCode)
	} else if resp.StatusCode == http.StatusOK && v.V != vexp {
		t.Error("For value",
			"expected", vexp,
			"got", v.V)
	}
}

func checksethandler(resp *http.Response, r *http.Request, t *testing.T, v *Value, vexp int64) {
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusMethodNotAllowed && resp.StatusCode != http.StatusBadRequest {
		t.Error("For", r,
			"expected", http.StatusOK, "or", http.StatusMethodNotAllowed, "or",
			http.StatusBadRequest, "got", resp.StatusCode)
	} else if resp.StatusCode == http.StatusOK && v.V != vexp {
		t.Error("For value",
			"expected", vexp,
			"got", v.V)
	}
}

func addhandlerTest(m string, v *Value, t *testing.T, h *Handlers) {
	req := httptest.NewRequest(m, "http://localhost:8080/add", nil)
	vexp := v.V + 1
	w := httptest.NewRecorder()
	h.AddHandler(w, req)
	resp := w.Result()
	checkhandlers(resp, req, t, v, vexp)
}

func dechandlerTest(m string, v *Value, t *testing.T, h *Handlers) {
	req := httptest.NewRequest(m, "http://localhost:8080/dec", nil)
	vexp := v.V - 1
	w := httptest.NewRecorder()
	h.DecHandler(w, req)
	resp := w.Result()
	checkhandlers(resp, req, t, v, vexp)
}

func reshandlerTest(m string, v *Value, t *testing.T, h *Handlers) {
	req := httptest.NewRequest(m, "http://localhost:8080/result", nil)
	vexp := v.V
	w := httptest.NewRecorder()
	h.ResHandler(w, req)
	resp := w.Result()
	checkhandlers(resp, req, t, v, vexp)
}

func sethandlerTest(req *http.Request, v *Value, vexp int64, t *testing.T, h *Handlers) {
	w := httptest.NewRecorder()
	h.SetHandler(w, req)
	resp := w.Result()
	checksethandler(resp, req, t, v, vexp)
}

func TestHandlersGet(t *testing.T) {
	v, h := initials()
	m := http.MethodGet
	addhandlerTest(m, v, t, h)
	dechandlerTest(m, v, t, h)
	reshandlerTest(m, v, t, h)
	setreq1 := httptest.NewRequest(m, "http://localhost:8080/set?value=20", nil)
	var vexp1 int64 = 20
	sethandlerTest(setreq1, v, vexp1, t, h)
}

func TestHandlersPOST(t *testing.T) {
	v, h := initials()
	m := http.MethodPost
	addhandlerTest(m, v, t, h)
	dechandlerTest(m, v, t, h)
	reshandlerTest(m, v, t, h)
	setreq1 := httptest.NewRequest(m, "http://localhost:8080/set?value=20", nil)
	var vexp1 int64 = 20
	sethandlerTest(setreq1, v, vexp1, t, h)
	setreq2 := httptest.NewRequest(m, "http://localhost:8080/set?value=monkey", nil)
	vexp2 := v.V
	sethandlerTest(setreq2, v, vexp2, t, h)
}
