package main

import (
	"app/internal/app"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var BranchName = "development"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
	t := os.Getenv("APPLICATION_PORT")
	log.Println("APP PORT", t)

	config, err := app.InitConfig("api")
	if err != nil {
		log.Fatal("failed to initialize application configuration", err)
	}

	if err := app.Run(config); err != nil {
		log.Fatal("failed to run application", err)
	}
	log.Println("Version: ", BranchName)
}
