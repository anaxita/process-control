package main

import (
	"flag"
	"log"

	"github.com/anaxita/process-control/internal/apiserver"
)

func main() {
	configPath := flag.String("config", "C:\\Program Files\\ProcessControl\\config.json", "path to json config file")
	flag.Parse()

	config, err := apiserver.NewConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	server := apiserver.NewServer(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
