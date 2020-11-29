package webserver

import (
	"github.com/idthings/idengine/pkg/datastore"
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

	http.HandleFunc("/identities/new/", handlerCreateIdentity)
	http.HandleFunc("/identities/rotate/", handlerRotateSecret)
	http.HandleFunc("/identities/", handlerValidate)
	http.HandleFunc("/epoch/", handlerEpoch)
	http.HandleFunc("/", handlerDefault)

	log.Println("webserver.Start(): listening on port", webserverPort)
	log.Fatal(http.ListenAndServe(":"+webserverPort, nil))
}
