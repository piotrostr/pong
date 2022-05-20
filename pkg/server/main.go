package server

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func RunServer() {
	log.Println("listening on 80")
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":http", nil); err != nil {
		log.Fatalln(err)
	}
}
