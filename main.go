package main

import (
	"chmafu/gorest/userservice"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	// Import the missing package
)

func main() {
	restful.Add(userservice.New())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
