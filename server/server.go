package server

import (
	"fmt"
	"net/http"
)

type server struct {
	http.Handler
}

func New() *server {
	fmt.Println("--------- sTaRtInG pAtChWoRk ------------")

	router := http.NewServeMux()

	router.HandleFunc("/status/check", statusCheckHandler)
	router.HandleFunc("/", defaultHandler)

	return &server{
		Handler: router,
	}
}

func defaultHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("HeLLo WorLd"))
}

func statusCheckHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("OK"))
}
