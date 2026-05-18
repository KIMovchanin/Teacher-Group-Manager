package main

import (
	"log"
	"net/http"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/config"
	httptransport "github.com/KIMovchanin/Teacher-Group-Manager/internal/transport/http"
)

func main() {
	cfg := config.Load()
	addr := ":" + cfg.HTTPPort

	mux := httptransport.NewRouter()

	log.Println("server is running on", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("server failed to start:", err)
	}
}
