package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//registering our services

	tasksService := NewTasksService(s.store)
	tasksService.RegisterRoutes(router)

	log.Println("Starting server on", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}