package model

import (
	"strconv"

	r "gopkg.in/gorethink/gorethink.v3"
)

// AndAs структура для хранения вопроса и ответа на него
type AndAs struct {
	Ask    string `json:"ask", gorethink:"ask"`
	Answer string `json:"answer", gorethink:"answer"`
}

// Faskurl структура для хранения страничек
type Faskurl struct {
	ID    string  `json:"id", gorethink:"id"`
	Murl  string  `json:"murl", gorethink:"murl"`
	Surl  string  `json:"surl", gorethink:"surl"`
	Fasks []AndAs `json:"fasks", gorethink:"fasks"`
}

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
func NewFask() (Faskurl, error) {

	var f Faskurl

	res, err := r.UUID().Run(session)
	if err != nil {
		return f, err
	}

	// получаем основной guid для адресса опросника
	var MUUID string
	err = res.One(&MUUID)
	if err != nil {
		return f, err
	}

	// получаем секретный guid для адресса опросника
	res, err = r.UUID().Run(session)
	if err != nil {
		return f, err
	}

	var SUUID string
	err = res.One(&SUUID)
	if err != nil {
		return f, err
	}

	// получаем id для объекта опросник
	res, err = r.UUID().Run(session)
	if err != nil {
		return f, err
	}

	var UUID string
	err = res.One(&UUID)
	if err != nil {
		return f, err
	}

	f.Murl = MUUID
	f.Surl = SUUID
	f.ID = UUID

	// производим запись в БД
	res, err = r.DB("Faskdb").Table("fasker").Insert(f).Run(session)
	if err != nil {
		return f, err
	}
	return f, nil
}

//NewAnswer функция для добовления ответа на определенный вопрос
func NewAnswer(url string, id string, answer string) error {

	res, err := r.DB("Faskdb").Table("fasker").Filter(map[string]interface{}{
		"Surl": url,
	}).Run(session)
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

//NewAsk функция для добовления нового вопроса
func NewAsk(url string, ask string) error {

	res, err := r.DB("Faskdb").Table("fasker").Filter(map[string]interface{}{
		"Murl": url,
	}).Run(session)
	if err != nil {
		return err
	}

	var f Faskurl
	err = res.One(&f)
	if err != nil {
		return err
	}

	var nask AndAs
	nask.Ask = ask

	f.Fasks = append(f.Fasks, nask)

	_, err = r.DB("Faskdb").Table("fasker").Get(f.ID).Replace(f).Run(session)
	if err != nil {
		return err
	}

	return nil
}

//GetMFask функция для получения вопросов и ответов на общую страницу вопросника
func GetMFask(url string) (Faskurl, error) {
	var f Faskurl

	res, err := r.DB("Faskdb").Table("fasker").Filter(map[string]interface{}{
		"Murl": url,
	}).Run(session)
	if err != nil {
		return f, err
	}

	err = res.One(&f)
	if err != nil {
		return f, err
	}

	return f, nil

}

//GetSFask функция для получения вопросов и ответов на страницу вопросника только для создавшего вопросник
func GetSFask(url string) (Faskurl, error) {
	var f Faskurl

	res, err := r.DB("Faskdb").Table("fasker").Filter(map[string]interface{}{
		"Surl": url,
	}).Run(session)
	if err != nil {
		return f, err
	}

	err = res.One(&f)
	if err != nil {
		return f, err
	}

	return f, nil

}
