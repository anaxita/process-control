package apiserver

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/user"
)

type jsonResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func (s *Server) controlHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		res := new(jsonResponse)

		count, err := s.CountProcess(s.config.KillProcess)
		if err != nil {
			s.Log(fmt.Sprintf("ОШИБКА получения количества процессов: %v", err))
			w.Write([]byte(fmt.Sprintf("ОШИБКА получения количества процессов: %v", err)))
			log.Println("1")

			return
		}

		if count == 0 {
			s.Log(fmt.Sprintf("Не найдено процессов для остановки. Count: %v .", count))
			res.Error = "Не найдено процессов для остановки"
			log.Println("2")

			json.NewEncoder(w).Encode(res)
			return
		}

		exitCode, err := s.StopProcess(s.config.KillProcess)
		if err != nil {
			log.Println("3")

			s.Log(fmt.Sprintf("Не удалось запустить команду остановки 1cv7s: %v", err))
			res.Error = fmt.Sprintf("Не удалось запустить команду остановки 1cv7s: %v", err)
			json.NewEncoder(w).Encode(res)
			return
		}

		if exitCode != 0 {
			log.Println("4")

			s.Log(fmt.Sprintf("Ошибка остановки процессов 1cv7s. Код завершения: %v", exitCode))
			res.Error = fmt.Sprintf("Ошибка остановки процессов 1cv7s. Код завершения: %v", exitCode)
			json.NewEncoder(w).Encode(res)
			return
		}

		s.Log(fmt.Sprintf("Все процессы остановлены, всего %v", count))
		res.Data = "Все процессы остановлены"
		json.NewEncoder(w).Encode(res)
	}
}

func (s *Server) listHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := user.Current()
		s.Log(fmt.Sprintf("Получаем список процессов, пользователь: %s | %s", u.Username, u.Name))

		res := new(jsonResponse)

		data, err := s.ProcessList(s.config.KillProcess)
		if err != nil {
			s.Log("Процессов не найдено")
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)

			return
		}

		s.Log("Процессов найдено:", len(data))
		res.Data = data
		json.NewEncoder(w).Encode(res)
	}
}

func (s *Server) indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("temaplate:", s.config.Template)
		tmpl, err := template.ParseFiles(s.config.Template)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		index := template.Must(tmpl, err)
		index.Execute(w, nil)
	}
}
