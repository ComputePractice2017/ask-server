package model

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

// AndAs структура для хранения вопроса и ответа на него
type AndAs struct {
	ID     string `json:"id"`
	Ask    string `json:"ask"`
	Answer string `json:"answer"`
}

// Faskurl структура для хранения страничек
type Faskurl struct {
	ID    string  `json:"id"`
	Murl  string  `json:"murl"`
	Surl  string  `json:"surl"`
	Fasks []AndAs `json:"fasks"`
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
