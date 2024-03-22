package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr  string
	store Store
}

func NewApiServer(addr string, store Store) *ApiServer {
	return &ApiServer{addr: addr, store: store}
}

func (s *ApiServer) Server() {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("api/v1").Subrouter()

	// registering our service

	log.Println("starting the api servive at ", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
