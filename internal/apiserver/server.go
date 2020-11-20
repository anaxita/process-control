package apiserver

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Server ...
type Server struct {
	config *Config
}

// NewServer ...
func NewServer(config *Config) *Server {
	return &Server{
		config: config,
	}
}

// Start ...
func (s *Server) Start() error {

	r := mux.NewRouter()

	r.HandleFunc("/control", s.controlHandler()).Methods("GET")
	r.HandleFunc("/test", s.testHandler()).Methods("GET")

	log.Println(fmt.Sprintf("Listening %s/control", s.config.Host))
	if err := http.ListenAndServe(s.config.Host, r); err != nil {
		return err
	}
	return nil
}

// Log ...
func (s *Server) Log(v ...interface{}) {

	f, _ := os.OpenFile(s.config.Logfile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0766)
	defer f.Close()

	log.SetOutput(f)
	log.Println(v...)
}
