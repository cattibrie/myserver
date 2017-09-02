package myhandler

import (
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
)

var V int64

func (s *S) AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nv := atomic.AddInt64(&s.V, 1)
		fmt.Fprintf(w, "Value v = %d", nv)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func DecHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nv := atomic.AddInt64(&V, -1)
		fmt.Fprintf(w, "Value v = %d", nv)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func ResHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		nv := atomic.LoadInt64(&V)
		fmt.Fprintf(w, "Value v = %d", nv)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func SetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		sval := r.URL.Query().Get("value")
		val, err := strconv.ParseInt(sval, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid value argument: %v", err)
			return
		}
		atomic.StoreInt64(&V, val)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
