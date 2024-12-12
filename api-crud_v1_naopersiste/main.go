package main

import (
	"api-crud/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/produto", handlers.GetProducts).Methods("GET")

	r.HandleFunc("/produtos/{id}", handlers.GetProduct).Methods("GET")

	r.HandleFunc("/produtos", handlers.CreateProduct).Methods("POST")

	r.HandleFunc("/produtos/{id}", handlers.UpdateProduct).Methods("PUT")

	r.HandleFunc("/produtos/{id}", handlers.DeleteProduct).Methods("DELETE")

	log.Println("Servidor rodando na porta 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
