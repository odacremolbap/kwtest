package handlers

import (
	"log"
	"net/http"

	"github.com/odacremolbap/kwtest/pkg/model"
)

// HealHandler heals the instance
func HealHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for %s", r.URL.Path)

	spoiler := model.GetSpoiler()
	spoiler.Stop()
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{'spoiling': False}"))
}
