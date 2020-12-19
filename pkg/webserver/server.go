package webserver

import (
	"flag"
	"github.com/idthings/idengine/pkg/datastore"
	"github.com/idthings/idengine/pkg/datastore/hashicorpvault"
	"log"
	"net/http"
	"os"
	"time"
)

const webserverPort = "8000"

var (
	ds datastore.Interface
)

func init() {

	vaultPtr := flag.Bool("vault", false, "use Vault datastore")
	flag.Parse()

	if *vaultPtr {

		ds = &hashicorpvault.Datastore{}

		result := ds.Connect()
		if !result {
			log.Println("Vault connection failed, exiting after 1 seconds.")
			time.Sleep(1 * time.Second)
			os.Exit(1)
		}
	} else {
		ds = &datastore.Datastore{}
		ds.Connect()
	}
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
