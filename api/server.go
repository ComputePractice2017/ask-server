package api

import (
	"log"
	"net/http"

	"github.com/ComputePractice2017/ask-server/model"
	"github.com/gorilla/mux"
)

//Run для запуска сервера
func Run() {

	log.Println("Connecting to rethinkDB on localhost...")
	err := model.InitSession()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	r := mux.NewRouter()
	r.HandleFunc("/", helloMFWorldHandler).Methods("GET")
	r.HandleFunc("/fask/{guid}/{guid1}/question", newAskHandler).Methods("POST")
	r.HandleFunc("/fask/{guid}", getMFaskHandler).Methods("GET")
	r.HandleFunc("/fask/{guid}/{guid1}", getSFaskHandler).Methods("GET")

	log.Println("Running the server on port 8000...")
	http.ListenAndServe(":8000", r)

}
