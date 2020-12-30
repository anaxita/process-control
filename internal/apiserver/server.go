package apiserver

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))
	r.HandleFunc("/", s.indexHandler()).Methods("GET")
	r.HandleFunc("/control", s.controlHandler()).Methods("GET")
	r.HandleFunc("/list", s.listHandler()).Methods("GET")

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

	_, err = f.WriteString(fmt.Sprintln(time.Now().Format("02.01.2006 15:04:05"), v))
	if err != nil {
		log.Println(err)
		return
	}
}
