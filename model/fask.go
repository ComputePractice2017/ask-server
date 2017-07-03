package model

import (
	"strconv"

	r "gopkg.in/gorethink/gorethink.v3"
)

// AndAs структура для хранения вопроса и ответа на него
type AndAs struct {
	ID     string `json:"id",gorethink:"id"`
	Ask    string `json:"ask",gorethink:"ask"`
	Answer string `json:"answer",gorethink:"answer"`
}

// Faskurl структура для хранения страничек
type Faskurl struct {
	ID    string  `json:"id",gorethink:"id"`
	Murl  string  `json:"murl".gorethink:"murl"`
	Surl  string  `json:"surl",gorethink:"surl"`
	Fasks []AndAs `json:"fasks",gorethink:"fasks"`
}

// Pagedata хранилище страничек вопросника
//var Pagedata []Faskurl

var session *r.Session

//InitSession активирует сессию связи с БД
func InitSession() error {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address: "localhost",
	})
	return err
}

//NewAnswer функция для добовления ответа на определенный вопрос
func NewAnswer(url string, id string, answer string) error {

	res, err := r.DB("Faskdb").Table("fasker").GetAllByIndex("surl", url).Run(session)
	if err != nil {
		return err
	}

	var f Faskurl
	err = res.One(&f)
	if err != nil {
		return err
	}

	var nid int
	nid, err = strconv.Atoi(id)
	if err != nil {
		return err
	}
	
	f.Fasks[nid].Answer = answer

	_, err = r.DB("Faskdb").Table("fasker").Get(f.ID).Replace(f).Run(session)
	if err != nil {
		return err
	}

	return nil

}
