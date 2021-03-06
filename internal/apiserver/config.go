package apiserver

import (
	"encoding/json"
	"log"
	"os"
)

// Config ..
type Config struct {
	Host        string `json:"host"`
	Logfile     string `json:"logfile"`
	Template    string `json:"template"`
	KillProcess string `json:"kill"`
}

// NewConfig ...
func NewConfig(configPath string) (*Config, error) {

	config := &Config{}
	log.Println("пришёл конфиг:", configPath)

	f, err := os.OpenFile(configPath, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("failed to open config file")
		return nil, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(config)

	if err != nil {
		log.Fatal("failed to decode config file")
		return nil, err
	}
	return config, nil
}
