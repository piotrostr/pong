package server

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.String(), r.UserAgent)
	fmt.Fprintln(w, "pong")
}

func RunServer() {
	log.Println("listening on 80")
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":http", nil); err != nil {
		log.Fatalln(err)
	}
}
