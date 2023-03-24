package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rangerslider/API/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("SERVER ARE STARTING ......")

	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening at port 4000 ......")
}
