package kwtest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/odacremolbap/kwtest/pkg/kwtest/handlers"
)

// Config contains server configuration
type Config struct {
	Port int
}

// Server holds a web test
type Server struct {
	config Config
}

// NewServer creates a new test server
func NewServer(config Config) (*Server, error) {
	s := &Server{
		config: config,
	}

	return s, nil
}

// Run starts the web server loop
func (s *Server) Run() error {

	s.createRoutes()

	log.Printf("Listening on port %d", s.config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), nil))
	return nil
}

func (s *Server) createRoutes() {
	http.HandleFunc("/", handlers.RootHandler)
}
