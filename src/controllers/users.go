package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Create_User(w http.ResponseWriter, r *http.Request) {
	var user models.User
	request_body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(request_body, &user); err != nil {
		log.Fatal(err)
	}

	db_result, err := db.Connection_db()

	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.New_repository_user(db_result)
	repository.Create(user)
}

func Get_Users(w http.ResponseWriter, r *http.Request) {

}

func Get_User(w http.ResponseWriter, r *http.Request) {

}

func Update_Users(w http.ResponseWriter, r *http.Request) {

}

func Delete_Users(w http.ResponseWriter, r *http.Request) {

}
