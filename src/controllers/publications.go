package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Create_Publication(w http.ResponseWriter, r *http.Request) {
	user_id, err := authentication.Extract_user_id(r)

	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	body_request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publications

	if err = json.Unmarshal(body_request, &publication); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthorID = user_id

	if err = publication.Prepare(); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	db_result, err := db.Connection_db()

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	defer db_result.Close()

	repository := repositories.New_repository_publication(db_result)
	publication_id, err := repository.Create(publication)

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publication_id)
}

func Get_Publication(w http.ResponseWriter, r *http.Request) {

}
func Update_Publication(w http.ResponseWriter, r *http.Request) {

}
func Delete_Publication(w http.ResponseWriter, r *http.Request) {

}
