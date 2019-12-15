package controllers

import (
	"github.com/jgersain/entropy-chat-api/api/middlewares"
)

func (server *Server) initializeRoutes() {
	//Login endpoint
	server.Router.HandleFunc("/api/login", middlewares.SetMiddlewareJSON(server.Login)).
		Methods("POST")

	//User endpoints
	server.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(server.CreateUser)).
		Methods("POST")
	server.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(
		middlewares.SetMiddlewareAuthentication(server.UpdateUser))).Methods("PUT")

	//Contacts endpoints
	server.Router.HandleFunc("/api/contacts", middlewares.SetMiddlewareJSON(server.CreateContact)).
		Methods("POST")
	// /api/contacts?user_id=[id]
	server.Router.HandleFunc("/api/contacts", middlewares.SetMiddlewareJSON(
		middlewares.SetMiddlewareAuthentication(server.GetContactsUser))).
		Queries("user_id", "{id}").
		Methods("GET")
	// /api/contacts/10?user_id=[id]
	server.Router.HandleFunc("/api/contacts/{id}", middlewares.SetMiddlewareJSON(
		middlewares.SetMiddlewareAuthentication(server.GetContactUser))).
		Queries("user_id", "{user_id}").
		Methods("GET")
}
