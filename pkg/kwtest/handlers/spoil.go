package handlers

import (
	"log"
	"net/http"

	"github.com/odacremolbap/kwtest/pkg/model"
)

// SpoilHandler spoils the instance
func SpoilHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for %s", r.URL.Path)

	spoiler := model.GetSpoiler()
	spoiler.Run()
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{'spoiling': True}"))
}
