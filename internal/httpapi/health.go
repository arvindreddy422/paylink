package httpapi

import "net/http"

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthz) 
	return mux
}

func healthz(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte("ok"))
}