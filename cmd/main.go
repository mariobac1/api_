package main

import (
	"log"
	"net/http"

	"github.com/mariobac1/api_/authorization"
	"github.com/mariobac1/api_/handler/community"
	"github.com/mariobac1/api_/handler/login"
	"github.com/mariobac1/api_/handler/person"
	"github.com/mariobac1/api_/storage/postgres"
	postCommunity "github.com/mariobac1/api_/storage/postgres/community"
	postPerson "github.com/mariobac1/api_/storage/postgres/person"
	postUser "github.com/mariobac1/api_/storage/postgres/user"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("We can't load the certificates: %v", err)
	}

	connection, _ := postgres.NewPostgresDB()

	store := postPerson.New(connection)
	commu := postCommunity.New(connection)
	usr := postUser.New(connection)
	mux := http.NewServeMux()

	if err := commu.Migrate(); err != nil {
		log.Fatalf("commu.Migrate: %v", err)
	}

	if err = store.Migrate(); err != nil {
		log.Fatalf("store.Migrate: %v", err)
	}

	if err = usr.Migrate(); err != nil {
		log.Fatalf("store.Migrate: %v", err)
	}

	person.RoutePerson(mux, store)
	community.RouteCommunity(mux, commu)
	login.RouteUser(mux, usr)

	log.Println("Server in port 8080 start")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error in server %v\n", err)
	}
}
