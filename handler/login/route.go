package login

import (
	"net/http"

	"github.com/mariobac1/api_/domain/user"
)

// Route Login
func RouteUser(mux *http.ServeMux, usecase user.Storage) {
	h := newLogin(usecase)

	mux.HandleFunc("/v1/login", h.login)
	mux.HandleFunc("/v1/sign-up", h.create)
}
