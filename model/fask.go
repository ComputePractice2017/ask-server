package model

import (
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

//NewFask функция создания нового опросника
func NewFask() (Faskurl, string, string, error) {

	var f Faskurl
	var murl string
	var surl string

	res, err := r.UUID().Run(session)
	if err != nil {
		return f, murl, surl, err
	}

	// получаем основной guid для адресса опросника
	var MUUID string
	err = res.One(&MUUID)
	if res != nil {
		return f, murl, surl, err
	}

	// получаем секретный guid для адресса опросника
	var SUUID string
	err = res.One(&SUUID)
	if res != nil {
		return f, murl, surl, err
	}

	// получаем id для объекта опросник
	var UUID string
	err = res.One(&UUID)
	if res != nil {
		return f, murl, surl, err
	}

	f.Murl = MUUID
	f.Surl = SUUID
	f.ID = UUID

	// производим запись в БД
	res, err = r.DB("Faskdb").Table("fasker").Insert(f).Run(session)
	if res != nil {
		return f, murl, surl, err
	}

	murl = "http://localhost:8000/fask/" + MUUID
	surl = "http://localhost:8000/fask/" + MUUID + "/" + SUUID

	return f, murl, surl, nil
}
