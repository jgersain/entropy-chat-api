package controllers

import (
	"github.com/jgersain/entropy-chat-api/api/middlewares"
)

func (server *Server) initializeRoutes() {

	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).
		Methods("POST")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).
		Methods("POST")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateUser))).
		Methods("PUT")

	server.Router.HandleFunc("/contacts", middlewares.SetMiddlewareJSON(server.CreateContact)).
		Methods("POST")
}
