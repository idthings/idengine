package webserver

import (
	"flag"
	"github.com/idthings/idengine/pkg/datastore"
	"github.com/idthings/idengine/pkg/datastore/hashicorpvault"
	"github.com/idthings/idengine/pkg/datastore/redis"
	"log"
	"net/http"
	"os"
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
	} else {
		ds = &redis.Datastore{}
	}

	result := ds.Connect()
	if !result {
		log.Println("Datasore connection failed, exiting...")
		os.Exit(1)
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
