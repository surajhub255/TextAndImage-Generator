package main

import (
	"back/services"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/generate", services.GenerateImage).Methods("GET")
	router.HandleFunc("/generate-text", services.HandleGenerateText).Methods("POST")

	http.Handle("/", router)

	port := 8080
	fmt.Printf("Server running on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port),
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTION"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(router))

}
