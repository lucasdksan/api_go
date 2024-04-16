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
	"strings"

	"github.com/gorilla/mux"
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

	if err := user.User_init("register"); err != nil {
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
	params := mux.Vars(r)

	user_id, err := strconv.ParseUint(params["id"], 10, 64)

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

	repository := repositories.New_repository_user(db)
	user, err := repository.Get_for_id(user_id)

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func Update_User(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)
	user_id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	user_id_token, err := authentication.Extract_user_id(r)

	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	if user_id != user_id_token {
		responses.ERR(w, http.StatusForbidden, errors.New("nao eh possivel fazer essa alteracao"))
		return
	}

	request_body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = json.Unmarshal(request_body, &user); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	if err := user.User_init("edition"); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connection_db()

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer db.Close()

	repository := repositories.New_repository_user(db)

	if err := repository.Update(user_id, user); err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func Delete_User(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user_id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	user_id_token, err := authentication.Extract_user_id(r)

	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	if user_id != user_id_token {
		responses.ERR(w, http.StatusForbidden, errors.New("nao eh possivel fazer essa remocao"))
		return
	}

	db, err := db.Connection_db()

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer db.Close()

	repository := repositories.New_repository_user(db)

	if err := repository.Delete(user_id); err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func Follow_user(w http.ResponseWriter, r *http.Request) {
	follower_id, err := authentication.Extract_user_id(r)
	params := mux.Vars(r)

	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	user_id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	if follower_id == user_id {
		responses.ERR(w, http.StatusForbidden, errors.New("não é possível seguir a si proprío"))
		return
	}

	db, err := db.Connection_db()

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer db.Close()

	repository := repositories.New_repository_user(db)

	if err = repository.Follow(follower_id, user_id); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func Un_follow_user(w http.ResponseWriter, r *http.Request) {
	follower_id, err := authentication.Extract_user_id(r)
	params := mux.Vars(r)

	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	user_id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	if follower_id == user_id {
		responses.ERR(w, http.StatusForbidden, errors.New("não é possível seguir a si proprío"))
		return
	}

	db, err := db.Connection_db()

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer db.Close()

	repository := repositories.New_repository_user(db)

	if err := repository.Un_follow(user_id, follower_id); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func Get_Followers(w http.ResponseWriter, r *http.Request) {
	user_id, err := authentication.Extract_user_id(r)

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

	repository := repositories.New_repository_user(db)
	followers, err := repository.Get_followers(user_id)

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}
