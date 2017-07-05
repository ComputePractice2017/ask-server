package model

import (
	"log"
	"os"
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
	dbaddress := os.Getenv("RETHINKDB_HOST")
	if dbaddress == "" {
		dbaddress = "localhost"
	}

	log.Printf("RETHINKDB_HOST: %s\n", dbaddress)
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address: dbaddress,
	})
	if err != nil {
		return err
	}

	err = CreateDBIfNotExist()
	if err != nil {
		return err
	}

	err = CreateTableIfNotExist()

	return err
}

//CreateDBIfNotExist функция создания БД если она ещще не создана
func CreateDBIfNotExist() error {
	res, err := r.DBList().Run(session)
	if err != nil {
		return err
	}

	var dbList []string
	err = res.All(&dbList)
	if err != nil {
		return err
	}

	for _, item := range dbList {
		if item == "Faskdb" {
			return nil
		}
	}

	_, err = r.DBCreate("Faskdb").Run(session)
	if err != nil {
		return err
	}

	return nil
}

// CreateTableIfNotExist функция создания таблицы в БД если она не создана
func CreateTableIfNotExist() error {
	res, err := r.DB("Faskdb").TableList().Run(session)
	if err != nil {
		return err
	}

	var tableList []string
	err = res.All(&tableList)
	if err != nil {
		return err
	}

	for _, item := range tableList {
		if item == "fasker" {
			return nil
		}
	}

	_, err = r.DB("Faskdb").TableCreate("fasker", r.TableCreateOpts{PrimaryKey: "ID"}).Run(session)

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
func NewAnswer(url string, id string, nanswer AndAs) error {

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

	f.Fasks[nid].Answer = nanswer.Answer

	_, err = r.DB("Faskdb").Table("fasker").Get(f.ID).Replace(f).Run(session)
	if err != nil {
		return err
	}

	return nil
}

//NewAsk функция для добовления нового вопроса
func NewAsk(url string, nask AndAs) error {

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

	//var nask AndAs
	//nask.Ask = ask

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
