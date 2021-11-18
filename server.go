package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/ws", WsHandler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Fatal("App running...")
	}
}
