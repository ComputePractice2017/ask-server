package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Run для запуска сервера
func Run() {

	r := mux.NewRouter()
	r.HandleFunc("/", helloMFWorldHandler).Methods("GET")

	log.Println("Running the server on port 8000...")
	http.ListenAndServe(":8000", r)

}
