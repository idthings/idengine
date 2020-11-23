package webserver

import (
	"log"
	"net/http"
)

const webserverPort = "8000"

// Start starts
func Start() {

	mw := nestedMiddleware()

	http.HandleFunc("/", mw(handlerDefault))

	http.HandleFunc("/identities/new/", mw(handlerCreateIdentity))
	http.HandleFunc("/identities/", mw(handlerValidate))

	log.Println("webserver.Start(): listening on port", webserverPort)
	log.Fatal(http.ListenAndServe(":"+webserverPort, nil))
}
