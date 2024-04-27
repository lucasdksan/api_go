package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	params := mux.Vars(r)
	publication_id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connection_db()

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer db.Close()

	repository := repositories.New_repository_publication(db)
	publication, err := repository.Get_for_id(publication_id)

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, publication)
}

func Get_Publications(w http.ResponseWriter, r *http.Request) {
	user_id_token, err := authentication.Extract_user_id(r)

	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	db, err := db.Connection_db()

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer db.Close()

	repository := repositories.New_repository_publication(db)
	publications, err := repository.Get(user_id_token)

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, publications)
}

func Update_Publication(w http.ResponseWriter, r *http.Request) {
	user_id_token, err := authentication.Extract_user_id(r)
	params := mux.Vars(r)

	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	publication_id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	body_request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := db.Connection_db()

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer db.Close()

	repository := repositories.New_repository_publication(db)
	publication_in_db, err := repository.Get_for_id(publication_id)

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if publication_in_db.AuthorID != user_id_token {
		responses.ERR(w, http.StatusForbidden, errors.New("não é possível atualizar um post que não seja sua"))
		return
	}

	var publication models.Publications

	if err = json.Unmarshal(body_request, &publication); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	if err = publication.Prepare(); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.Update(publication_id, publication); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func Delete_Publication(w http.ResponseWriter, r *http.Request) {

}
