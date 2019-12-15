package controllers

import (
	"github.com/jgersain/entropy-chat-api/api/middlewares"
)

func (server *Server) initializeRoutes() {
	//Login endpoint
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).
		Methods("POST")
	//Register user endpoint
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).
		Methods("POST")
	//Update profile endpoint
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(
		middlewares.SetMiddlewareAuthentication(server.UpdateUser))).Methods("PUT")
	//Contacts endpoints
	server.Router.HandleFunc("/contacts", middlewares.SetMiddlewareJSON(server.CreateContact)).
		Methods("POST")
	server.Router.HandleFunc("/contacts", middlewares.SetMiddlewareJSON(
		middlewares.SetMiddlewareAuthentication(server.GetContactsUser))).Queries("user_id", "{id}").
		Methods("GET")
}
