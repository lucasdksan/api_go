package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	token, _ := authentication.Create_token(user_exist_db.ID)
	fmt.Println("Token: ", token)
	w.Write([]byte("Parabens voce esta logado!"))
}
