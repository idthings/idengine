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
	ds.Connect()
}

// Start starts
func Start() {

	mw := nestedMiddleware()

	http.HandleFunc("/identities/new/", mw(handlerCreateIdentity))
	http.HandleFunc("/identities/rotate/", mw(handlerRotateSecret))
	http.HandleFunc("/identities/", mw(handlerValidate))
	http.HandleFunc("/epoch/", mw(handlerEpoch))
	http.HandleFunc("/", mw(handlerDefault))

	log.Println("webserver.Start(): listening on port", webserverPort)
	log.Fatal(http.ListenAndServe(":"+webserverPort, nil))
}
