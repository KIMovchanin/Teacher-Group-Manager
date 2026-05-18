package http

import (
	"encoding/json"
	"log"
	nethttp "net/http"
)

func NewRouter() *nethttp.ServeMux {
	mux := nethttp.NewServeMux()

	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/students", studentsHandler)

	return mux
}

func healthHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	if r.Method != nethttp.MethodGet {
		nethttp.Error(w, "method not allowed", nethttp.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(nethttp.StatusOK)

	responce := map[string]string{"status": "ok"}

	if err := json.NewEncoder(w).Encode(responce); err != nil {
		log.Println("failed to encode health responce:", err)
	}
}

func studentsHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	if r.Method != nethttp.MethodGet {
		nethttp.Error(w, "method not allowed", nethttp.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(nethttp.StatusOK)

	responce := []map[string]any{
		{
			"id":         1,
			"first_name": "Ivan",
			"last_name":  "Petrov",
		},
		{
			"id":         2,
			"first_name": "Anna",
			"last_name":  "Sidorova",
		},
	}

	if err := json.NewEncoder(w).Encode(responce); err != nil {
		log.Println("failed to encode students response:", err)
	}
}
