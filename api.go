package main

import (
	"log"
	"net/http"
)

type ApiServer struct {
	Add string
}

func NewApiServer(add string) *ApiServer {
	return &ApiServer{Add: add}
}

func (s *ApiServer) Run() error {
	router := http.NewServeMux()

	server := http.Server{
		Addr:    s.Add,
		Handler: router,
	}
	log.Printf("server is running on post %s >>> http://localhost:%s", s.Add, s.Add)
	return server.ListenAndServe()
}
