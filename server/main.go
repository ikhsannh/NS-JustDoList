package main

import (
	"github.com/nashkispace/NS-JustDoList/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Server is running on port 4000")

	log.Fatal(http.ListenAndServe(":4000", r))
}
