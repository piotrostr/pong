package server

import (
	"log"
	"net/http"
)

func RunServer() {
	log.Println("listening on 80")
	if err := http.ListenAndServe(":http", nil); err != nil {
		log.Fatalln(err)
	}
}
