package apiserver

import (
	"fmt"
	"log"
	"net/http"
)

func (s *Server) testHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, err := s.CountProcess()
		if err != nil {
			s.Log(fmt.Sprintf("Не найдено активных процессов: %v", err))
			log.Println("count err:", err)

		}
		log.Println("count:", count)
	}
}
