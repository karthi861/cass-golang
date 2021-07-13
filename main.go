package main

import (
	"encoding/json"
	"inventory/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type home struct {
	Message string `json:"message"`
}

func homelink(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(home{Message: "Welcome"})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homelink)
	router.HandleFunc("/allproducts", handlers.GetallProducts).Methods("GET")
	router.HandleFunc("/product/{id}", handlers.GetProduct).Methods("GET")
	router.HandleFunc("/addproduct", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/updateproduct/{id}", handlers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/deleteproduct/{id}", handlers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/createcategory", handlers.CreateCategory).Methods("POST")
	router.HandleFunc("/allcategories", handlers.Getallcategories).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
