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
	r.HandleFunc("/", s.indexHandler()).Methods("GET")
	r.HandleFunc("/control", s.controlHandler()).Methods("GET")
	r.HandleFunc("/list", s.listHandler()).Methods("GET")
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))

	log.Println(fmt.Sprintf("Listening %s/control", s.config.Host))
	if err := http.ListenAndServe(s.config.Host, r); err != nil {
		return err
	}
	return nil
}

// Log ...
func (s *Server) Log(v ...interface{}) {
	log.Println(v...)

	f, err := os.OpenFile(s.config.Logfile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0766)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	f.WriteString(fmt.Sprintln(v...))
}
