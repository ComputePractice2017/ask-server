package api

import (
	"log"
	"net/http"
	"os"

	"github.com/ComputePractice2017/ask-server/model"
	"github.com/gorilla/handlers"
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

	r.HandleFunc("/fask", newFaskHandler).Methods("POST")
	r.HandleFunc("/fask", firstOptionsHandler).Methods("OPTIONS")

	r.HandleFunc("/fask/{guid}", getMFaskHandler).Methods("GET")
	r.HandleFunc("/fask/{guid}", secondOptionsHandler).Methods("OPTIONS")

	r.HandleFunc("/fask/{guid}/{guid1}", getSFaskHandler).Methods("GET")
	r.HandleFunc("/fask/{guid}/{guid1}", threeOptionsHandler).Methods("OPTIONS")

	r.HandleFunc("/fask/{guid}/{guid1}/question/{id}", newAnswerHandler).Methods("POST")
	r.HandleFunc("/fask/{guid}/{guid1}/question/{id}", fourOptionsHandler).Methods("OPTIONS")

	r.HandleFunc("/fask/{guid}/question", newAskHandler).Methods("POST")
	r.HandleFunc("/fask/{guid}/question", fiveOptionsHandler).Methods("OPTIONS")

	log.Println("Running the server on port 8000...")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(r))

}
