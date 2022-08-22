package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func ejemplon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my REST API of EDD, hopefully you enjoy it! :)")
}

func consulta_1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Consulta 1 :3")
}

func consulta_2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Consulta 2 :3")
}
func consulta_3(w http.ResponseWriter, r *http.Request) {

}
func consulta_4(w http.ResponseWriter, r *http.Request) {

}
func consulta_5(w http.ResponseWriter, r *http.Request) {

}
func consulta_6(w http.ResponseWriter, r *http.Request) {

}
func consulta_7(w http.ResponseWriter, r *http.Request) {

}
func consulta_8(w http.ResponseWriter, r *http.Request) {

}
func consulta_9(w http.ResponseWriter, r *http.Request) {

}
func consulta_10(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("Holis")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", ejemplon).Methods("GET")
	router.HandleFunc("/consulta_1", consulta_1).Methods("GET")
	router.HandleFunc("/consulta_2", consulta_2).Methods("GET")
	router.HandleFunc("/consulta_3", consulta_3).Methods("GET")
	router.HandleFunc("/consulta_4", consulta_4).Methods("GET")
	router.HandleFunc("/consulta_5", consulta_5).Methods("GET")
	router.HandleFunc("/consulta_6", consulta_6).Methods("GET")
	router.HandleFunc("/consulta_7", consulta_7).Methods("GET")
	router.HandleFunc("/consulta_8", consulta_8).Methods("GET")
	router.HandleFunc("/consulta_9", consulta_9).Methods("GET")
	router.HandleFunc("/consulta_10", consulta_10).Methods("GET")

	header := handlers.AllowedHeaders([]string{"X-Requested-with", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(header, methods, origins)(router)))
	fmt.Println("Adios")

}
