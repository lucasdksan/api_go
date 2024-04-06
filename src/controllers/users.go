package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func Create_User(w http.ResponseWriter, r *http.Request) {
	var user models.User
	request_body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = json.Unmarshal(request_body, &user); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	if err := user.User_init(); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	db_result, err := db.Connection_db()

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	defer db_result.Close()

	repository := repositories.New_repository_user(db_result)
	user_id, err := repository.Create(user)

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	user.ID = user_id
	responses.JSON(w, http.StatusCreated, user)
}

func Get_Users(w http.ResponseWriter, r *http.Request) {
	identifier := strings.ToLower(r.URL.Query().Get("user"))
	db, err := db.Connection_db()

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer db.Close()

	repository := repositories.New_repository_user(db)
	users, err := repository.Get(identifier)

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func Get_User(w http.ResponseWriter, r *http.Request) {

}

func Update_Users(w http.ResponseWriter, r *http.Request) {

}

func Delete_Users(w http.ResponseWriter, r *http.Request) {

}
