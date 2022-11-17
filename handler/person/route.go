package person

import (
	"net/http"

	"github.com/mariobac1/api_/domain/person"
)

func RoutePerson(mux *http.ServeMux, usecase person.Storage) {
	h := newHandler(usecase)

	mux.HandleFunc("/v1/persons/create", h.create)
}
