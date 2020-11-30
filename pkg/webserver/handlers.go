package webserver

import (
	"github.com/idthings/idengine/pkg/tasks/create"
	"github.com/idthings/idengine/pkg/tasks/validate"
	"log"
	"net/http"
	"strconv"
	"time"
)

var ()

func handlerDefault(w http.ResponseWriter, r *http.Request) {

	log.Println("webserver.handlerDefault(): request ", r.Method, r.URL.Path)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
}

func handlerCreateIdentity(w http.ResponseWriter, r *http.Request) {

	log.Println("webserver.handlerCreateIdentity(): request ", r.Method, r.URL.Path)

	status, response := create.Identity(r.Context(), &ds, r)

	w.WriteHeader(status)
	w.Write([]byte(response))
	w.Write([]byte("\n"))
}

func handlerRotateSecret(w http.ResponseWriter, r *http.Request) {

	log.Println("webserver.handlerRotateSecret(): request ", r.Method, r.URL.Path)

	status, response := create.RotateSecret(r.Context(), &ds, r)

	w.WriteHeader(status)
	w.Write([]byte(response))
	w.Write([]byte("\n"))
}

func handlerValidate(w http.ResponseWriter, r *http.Request) {

	// defaults
	status := http.StatusUnauthorized
	response := "Unauthorized\n"

	log.Println("webserver.handlerValidate(): request ", r.Method, r.URL.Path)

	// a password validation
	if r.Header.Get("X-idThings-Password") != "" {
		status, response = validate.Secret(r.Context(), &ds, r)
	} else if r.Header.Get("X-idThings-Digest") != "" {
		status, response = validate.Digest(r.Context(), &ds, r)
	}

	w.WriteHeader(status)
	w.Write([]byte(response))
	w.Write([]byte("\n"))
}

// returns unix timestamp, intended for electronic devices with no
// real-time-clock that need to create HMAC digests.
// returns:
//   {1606515898312}
func handlerEpoch(w http.ResponseWriter, r *http.Request) {

	t := time.Now().UnixNano() / 1e6 // convert to milliseconds
	tString := strconv.FormatInt(t, 10)
	response := "{" + tString + "}\n"

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
	w.Write([]byte("\n"))
}
