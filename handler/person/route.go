package person

import (
	"net/http"

	"github.com/mariobac1/api_/domain/person"
	"github.com/mariobac1/api_/middleware"
)

func RoutePerson(mux *http.ServeMux, usecase person.Storage) {
	h := newHandler(usecase)

	mux.HandleFunc("/v1/persons/create", middleware.Log(middleware.Authentication(h.create)))
	mux.HandleFunc("/v1/persons/all", middleware.Log(h.getAll))
	mux.HandleFunc("/v1/persons/getby", middleware.Log(h.getById))
	mux.HandleFunc("/v1/persons/update", middleware.Log(h.update))
	mux.HandleFunc("/v1/persons/delete", middleware.Log(h.delete))
}
