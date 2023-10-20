package main

import (
	"L10WB3/handler"
	"L10WB3/repository/memorystorage"
	"log"
	"net/http"
	"os"
)

func main() {
	storage := memorystorage.NewMemoryStorage()

	h := handler.NewHandler(storage, log.New(os.Stdout, "Error: ", 0), log.New(os.Stdout, "Info: ", 0))

	http.ListenAndServe("localhost:8080", h.InitHandler())
}