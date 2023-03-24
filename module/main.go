package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Ishmoth Ura Nuri")
	hello()
	r := mux.NewRouter()
	r.HandleFunc("/", web)

	log.Fatal(http.ListenAndServe(":4000", r))

}
func hello() {
	fmt.Println("be the boss ")
}
func web(w http.ResponseWriter, r *http.Request) {
	// jeta diye request kore seita r ..  okey use korty parbo .. thn w er majhe likhe pass kore dibo
	w.Write([]byte("<h1>nuri don't forget your value , be the boss</h1>"))

}
