package main

import (
	"file-sharing-system/handlers"
	"file-sharing-system/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
    utils.ConnectDB()
    http.HandleFunc("/register", handlers.Register)
    http.HandleFunc("/login", handlers.Login)
    http.HandleFunc("/upload", handlers.UploadFile)

    fmt.Printf("Starting server at port 8080\n")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
