package main

import (
	"flag"
	"log"

	"github.com/anaxita/process-control/internal/apiserver"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "C:\\Program Files\\Killer1c77\\config.json", "path to json config file")
}
func main() {
	flag.Parse()

	config, err := apiserver.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	server := apiserver.NewServer(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
