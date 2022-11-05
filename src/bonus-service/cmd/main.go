package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bmstu-rsoi/rsoi-2022-lab2-microservices-mRrvz/src/bonus-service/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")
	r := handlers.Router()
	log.Println("Server is listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
