package webserver

import (
	"github.com/idthings/idengine/internal/datastore"
	"github.com/idthings/idengine/internal/tasks/create"
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
