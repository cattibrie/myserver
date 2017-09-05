package myhandler

import (
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
)

type Value struct {
	V int64
}

type Operat interface {
	add() int64
	dec() int64
	res() int64
	set(nv *int64)
}

type Handlers struct {
	H Operat
}

func (v *Value) add() int64 {
	return atomic.AddInt64(&v.V, 1)
}

func (v *Value) dec() int64 {
	return atomic.AddInt64(&v.V, -1)
}

func (v *Value) res() int64 {
	return atomic.LoadInt64(&v.V)
}

func (v *Value) set(nv *int64) {
	atomic.StoreInt64(&v.V, *nv)
}

func (h *Handlers) AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//nv := atomic.AddInt64(&s.V, 1)
		fmt.Fprintf(w, "Value v = %d", h.H.add())
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *Handlers) DecHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//nv := atomic.AddInt64(&V, -1)
		fmt.Fprintf(w, "Value v = %d", h.H.dec())
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *Handlers) ResHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//nv := atomic.LoadInt64(&V)
		fmt.Fprintf(w, "Value v = %d", h.H.res())
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *Handlers) SetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		sval := r.URL.Query().Get("value")
		nv, err := strconv.ParseInt(sval, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid value argument: %v", err)
			return
		}
		//atomic.StoreInt64(&V, val)
		h.H.set(&nv)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
