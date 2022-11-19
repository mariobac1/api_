package person

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/mariobac1/api_/domain/person"
	"github.com/mariobac1/api_/models"
)

type handler struct {
	usecase person.Storage
}

func newHandler(usecase person.Storage) handler {
	return handler{usecase}
}

// handler Create a person
func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Method not permit", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data := models.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "The structure is wrong", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	err = h.usecase.Create(&data)
	if err != nil {
		response := newResponse(Error, "An issue occurs when try create a person", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "Person created Ok", nil)
	responseJSON(w, http.StatusCreated, response)
}

// handler GetAll persons
func (h *handler) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Method not permit", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := h.usecase.GetAll()
	if err != nil {
		response := newResponse(Error, "An issue occurs when try get all persons", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)

}

// Handler Update
func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Method not permit", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || ID < 1 {
		response := newResponse(Error, "The ID will be positive", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := models.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "The structure is wrong", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data.ID = uint(ID)
	data.UpdatedAt = time.Now()

	err = h.usecase.Update(&data)
	if errors.Is(err, models.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "This ID doesn't exist", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "An issue occurs when try update a person", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "Person update ok", nil)
	responseJSON(w, http.StatusOK, response)
	return
}

// handler byid
func (h *handler) getById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Method not permit", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || ID < 1 {
		response := newResponse(Error, "The ID will be positive", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := models.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "The structure is wrong", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data.ID = uint(ID)
	data.UpdatedAt = time.Now()

	err = h.usecase.Update(&data)
	if errors.Is(err, models.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "This ID doesn't exist", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "An issue occurs when try update a person", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "Person update ok", nil)
	responseJSON(w, http.StatusOK, response)
	return
}

// Handler Delete person
func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Method not permit", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || ID < 1 {
		response := newResponse(Error, "The ID will be positive", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = h.usecase.Delete(uint(ID))
	if errors.Is(err, models.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "This ID doesn't exist", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "An issue occurs when try update a person", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Person delete ok", nil)
	responseJSON(w, http.StatusOK, response)
	return
}
