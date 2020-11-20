package apiserver

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// StopProcess ...
func (s *Server) StopProcess() (int, error) {
	c := exec.Command("powershell", "/c", "Get-Process -Name notepad | Select -Property Id | Stop-Process -Force")
	err := c.Start()

	if err != nil {
		log.Println(err)
		s.Log(fmt.Sprintf("Ошибка запуска powershell %v", err))
		return 0, err
	}

	c.Wait()
	return c.ProcessState.ExitCode(), nil
}

// CountProcess ...
func (s *Server) CountProcess() (int, error) {
	c := exec.Command("powershell", "/c", `(Get-Process -ProcessName notepad).count`)
	out, err := c.Output()
	if err != nil {
		return 0, err
	}

	split := strings.Split(string(out), ":")
	trim := strings.TrimSpace(split[0])
	count, _ := strconv.Atoi(trim)

	return count, nil
}
