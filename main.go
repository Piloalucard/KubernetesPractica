package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo en GO!")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func main() {
	fmt.Println("App is running in port: 3000")
	setupRoutes()
	http.ListenAndServe("3000", nil)
}
