package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice2017/ask-server/model"
)

func helloMFWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World MZF!!!")
}

func newFaskHandler(w http.ResponseWriter, r *http.Request) {

	var murl string
	var surl string

	murl, surl, err := model.NewFask()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	//как я понимаю я должен как то передать в Encode адреса чтобы они вернулись из функции
	if err = json.NewEncoder(w).Encode(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	return
}
