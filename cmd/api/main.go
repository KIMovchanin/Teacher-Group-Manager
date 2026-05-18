package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)

	return mux
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	response := map[string]string{"status": "ok"}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("failed to encode health response:", err)
	}
}

func main() {
	addr := ":8080"

	mux := newRouter()

	log.Println("server is running on", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("server failed to start:", err)
	}
}
