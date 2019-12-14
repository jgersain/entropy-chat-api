package api

import (
	"github.com/jgersain/entropy-chat-api/api/controllers"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var server = controllers.Server{}

func Run() {

	//load .env file
	e := godotenv.Load()
	if e != nil {
		log.Fatalf("Error al cargar el archivo .env %v", e)
	}

	server.Initialize(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	server.Run(":8080")
}
