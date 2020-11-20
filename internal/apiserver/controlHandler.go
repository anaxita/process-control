package apiserver

import (
	"fmt"
	"net/http"
	"os/user"
)

func (s *Server) controlHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := user.Current()
		s.Log(fmt.Sprintf("%s | %s", u.Username, u.Name))

		count, err := s.CountProcess()
		if err != nil {
			s.Log(fmt.Sprintf("ОШИБКА получения количества процессов: %v", err))
		}

		if count == 0 {
			s.Log(fmt.Sprintf("Не найдено процессов для остановки. Count: %v .", count))
			w.Write([]byte(fmt.Sprintf("Не найдено процессов для остановки. Count: %v .", count)))
			return
		}

		exitCode, err := s.StopProcess()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Не удалось запустить команду остановки: %v", err)))
			return
		}

		if exitCode != 0 {
			s.Log(fmt.Sprintf("Ошибка остановки процессов. Код завершения: %v", exitCode))
			w.Write([]byte(fmt.Sprintf("Ошибка остановки процессов. Код завершения: %v", exitCode)))
			return
		}

		newcount, err := s.CountProcess()
		if err != nil {
			s.Log(fmt.Sprintf("ОШИБКА получения количества процессов: %v", err))
		}

		s.Log(fmt.Sprintf("Найдено: %v | Остановлено: %v | Не остановлено: %v | Код завершения: %v", count, count-newcount, newcount, exitCode))
		w.Write([]byte(fmt.Sprintf("Найдено: %v | Остановлено: %v | Не остановлено: %v", count, count-newcount, newcount)))
	}
}
