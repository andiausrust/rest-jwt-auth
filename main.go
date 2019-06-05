package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main() {

	router := mux.NewRouter()

	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleware(ProtectedEndpoint)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8888", router))
}

func signup(w http.ResponseWriter, r *http.Request){}