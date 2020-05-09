package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JoeyPilla/rest-api-example/api"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "joeypilla"
	password = ""
	dbname   = "restapiexample"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	api.Connect()
	handleRequests()
}
