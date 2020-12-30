package apiserver

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// Win32process ...
type Win32process struct {
	ID    string
	Name  string
	Owner string
}

// StopProcess ...
func (s *Server) StopProcess(processName string) (int, error) {
	comm := fmt.Sprintf("Get-Process -Name %s | Select -Property Id | Stop-Process -Force", processName)
	c := exec.Command("powershell", "/c", comm)

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
func (s *Server) CountProcess(processName string) (int, error) {
	comm := fmt.Sprintf("(Get-Process -ProcessName %s).count", processName)

	c := exec.Command("powershell", "/c", comm)
	out, err := c.Output()
	if err != nil {
		return 0, err
	}

	trim := strings.TrimSpace(string(out))
	count, _ := strconv.Atoi(trim)

	return count, nil
}

// ProcessList ...
func (s *Server) ProcessList(processName string) ([]Win32process, error) {

	comm := fmt.Sprintf("Get-Process -ProcessName %s -IncludeUserName | Select UserName, ID", processName)
	c := exec.Command("powershell", "/c", comm)

	out, err := c.Output()
	if err != nil {
		return nil, err
	}

	sList := strings.Split(strings.TrimSpace(string(out)), "\r\n")
	newList := sList[2:]

	var processes []Win32process

	for _, v := range newList {
		sProcess := strings.Split(strings.TrimSpace(string(v)), " ")
		if len(sProcess) == 2 {
			newProcess := Win32process{
				ID:    sProcess[1],
				Name:  processName,
				Owner: sProcess[0],
			}

			processes = append(processes, newProcess)
		}
	}

	return processes, nil
}
