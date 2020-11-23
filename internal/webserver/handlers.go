package webserver

import (
	"github.com/idthings/idengine/internal/datastore"
	"github.com/idthings/idengine/internal/tasks/create"
	"github.com/idthings/idengine/internal/tasks/validate"
	"log"
	"net/http"
)

var (
	ds datastore.Datastore
)

func handlerDefault(w http.ResponseWriter, r *http.Request) {

	log.Println("webserver.handlerDefault(): request ", r.Method, r.URL.Path)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func handlerCreateIdentity(w http.ResponseWriter, r *http.Request) {

	log.Println("webserver.handlerCreateIdentity(): request ", r.Method, r.URL.Path)

	status, response := create.Identity(ds, r)

	w.WriteHeader(status)
	w.Write([]byte(response))
}

func handlerValidate(w http.ResponseWriter, r *http.Request) {

	// defaults
	status := http.StatusUnauthorized
	response := "Unauthorized\n"

	log.Println("webserver.handlerValidate(): request ", r.Method, r.URL.Path)

	// a password validation
	if r.Header.Get("X-idThings-Password") != "" {
		status, response = validate.Secret(ds, r)

	}

	w.WriteHeader(status)
	w.Write([]byte(response))
}
