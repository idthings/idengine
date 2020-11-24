package webserver

import (
	"github.com/idthings/idengine/internal/datastore"
	"log"
	"net/http"
)

const webserverPort = "8000"

var (
	ds = datastore.Datastore{}
)

func init() {
	ds.Info()
	log.Println("here")
	ds.Connect()
}

// Start starts
func Start() {

	mw := nestedMiddleware()

	http.HandleFunc("/", mw(handlerDefault))

	http.HandleFunc("/identities/new/", mw(handlerCreateIdentity))
	http.HandleFunc("/identities/", mw(handlerValidate))

	log.Println("webserver.Start(): listening on port", webserverPort)
	log.Fatal(http.ListenAndServe(":"+webserverPort, nil))
}
