package api

import (
	"encoding/json"
	"fmt"

	"io"
	"io/ioutil"

	"log"
	"net/http"

	"github.com/ComputePractice2017/ask-server/model"

	"github.com/gorilla/mux"

)

func helloMFWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World MZF!!!")
}




func newFaskHandler(w http.ResponseWriter, r *http.Request) {
var fasks model.Faskurl

	fasks, err := model.NewFask()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(fasks); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}
func newAnswerHandler(w http.ResponseWriter, r *http.Request) {

var answer string

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := json.Unmarshal(body, &answer); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	vars := mux.Vars(r)
	err = model.NewAnswer(vars["guid1"], vars["id"], answer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func newAskHandler(w http.ResponseWriter, r *http.Request) {

	var ask string


	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := json.Unmarshal(body, &ask); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	vars := mux.Vars(r)

	err = model.NewAsk(vars["guid"], ask
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)

  }


func getMFaskHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var fasks model.Faskurl

	fasks, err := model.GetMFask(vars["guid"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err = json.NewEncoder(w).Encode(fasks); err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func getSFaskHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var fasks model.Faskurl

	fasks, err := model.GetSFask(vars["guid1"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err = json.NewEncoder(w).Encode(fasks); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)

		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)

}
