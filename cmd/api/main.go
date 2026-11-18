package main

import (
	"log"
	"net/http"

	"github.com/arvindreddy422/paylink/internal/httpapi"
)


func main() {
	srv := &http.Server{
		Addr: ":8080",
		Handler: httpapi.NewMux(),
	}

	log.Fatal(srv.ListenAndServe())

}