package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jgersain/entropy-chat-api/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(username, password, host, name string) {

	var err error
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, username, name, password)
	server.DB, err = gorm.Open("postgres", dbUri)

	if err != nil {
		fmt.Printf("No es posible conectarse a la base de datos\n")
	} else {
		fmt.Printf("Conexión a la base de datos realizada\n")
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Contact{})

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Ejecutando... http://127.0.0.1%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
