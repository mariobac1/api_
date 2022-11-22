package login

import (
	"encoding/json"
	"net/http"

	"github.com/mariobac1/api_/authorization"
	"github.com/mariobac1/api_/domain/user"
	per "github.com/mariobac1/api_/handler/person"
	"github.com/mariobac1/api_/models"
)

const (
	Error   = "error"
	Message = "message"
)

type login struct {
	storage user.Storage
}

func newLogin(s user.Storage) login {
	return login{s}
}

// handler Create a person
func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := per.NewResponse(Error, "Method not permit", nil)
		per.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := models.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp := per.NewResponse(Error, "struct no valid", nil)
		per.ResponseJSON(w, http.StatusBadRequest, resp)
		return
	}
	_, valid, _ := l.storage.GetByEmail(data.Email, data.Password)

	if !valid {
		resp := per.NewResponse(Error, "user or password not valid", nil)
		per.ResponseJSON(w, http.StatusBadRequest, resp)
		return
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := per.NewResponse(Error, "We can't make a new token", nil)
		per.ResponseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	dataToken := map[string]string{"token": token}
	resp := per.NewResponse(Message, "Ok", dataToken)
	per.ResponseJSON(w, http.StatusOK, resp)
}

// handler signUp a User
func (l *login) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := per.NewResponse(Error, "Method not permit", nil)
		per.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	data := models.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := per.NewResponse(Error, "The structure is wrong", nil)
		per.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	err = l.storage.Create(&data)
	if err != nil {
		response := per.NewResponse(Error, "An issue occurs when try create a person", nil)
		per.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := per.NewResponse(Message, "Person created Ok", nil)
	per.ResponseJSON(w, http.StatusCreated, response)
}
