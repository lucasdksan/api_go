package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	request_body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := json.Unmarshal(request_body, &user); err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db_result, err := db.Connection_db()

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	defer db_result.Close()

	repository := repositories.New_repository_user(db_result)
	user_exist_db, err := repository.Search_user(user.Email)

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(user.Password, user_exist_db.Password); err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.Create_token(user_exist_db.ID)

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	user_id := strconv.FormatUint(user_exist_db.ID, 10)

	responses.JSON(w, http.StatusOK, models.Authentication_data{ID: user_id, Token: token})
}
