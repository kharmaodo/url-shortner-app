package main

import (
	"log"
	"net/http"
	"url-shortner-app/internal/controllers"
)

func main() {

	http.HandleFunc("/", controllers.ShowIndex)
	http.HandleFunc("/shorten", controllers.Shorten)
	log.Fatal(http.ListenAndServe(":8787", nil))

}
