package main

import (
	"fmt"
	"log"

	config "github.com/CrudOperation/configs"
	"github.com/CrudOperation/server"
)

func main() {
	log.Println("starting api server...")

	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	} else {
		log.Println("config loaded successfully.")
	}
	fmt.Printf("AppVersion: %s \n", cfg.Server.AppVersion)

	srv := server.NewServer(cfg)

	srv.ListenAndServe()

}
