package validate

import (
	"fmt"
	"github.com/idthings/idengine/internal/data"
	"log"
	"net/http"
	"strings"
)

const (
	digestHeaderName    = "X-idThings-Digest"
	digestHeaderFields  = 4
	digestTypeField     = 0
	digestField         = 1
	digestEpochField    = 2
	digestDataField     = 3
	maxDigestAgeMinutes = 5
)

// Digest validates an input digest with content
func Digest(store FetchSecretInterface, r *http.Request) (int, string) {

	// get the requesting device id
	id, err := extractGUID(r.URL.Path)
	if err != nil {
		return http.StatusNotFound, "Not Found"
	}

	// ensure digest header is in the request
	digestString := r.Header.Get(digestHeaderName)
	if len(digestString) == 0 {
		log.Println("validate.Digest(): digest header string empty")
		return http.StatusBadRequest, "Bad Request: missing header"
	}

	// ensure digest header has required number of values
	values := strings.SplitN(digestString, ",", digestHeaderFields)
	if len(values) != digestHeaderFields {
		return http.StatusBadRequest, "Bad Request: invalid header"
	}

	// attempt to fetch the secret for this id
	secret, err := store.FetchSecret(id)
	if err != nil {
		log.Println("validate.Digest():", err.Error())
		return http.StatusInternalServerError, "Internal error"
	}

	var computedDigest string

	switch values[digestTypeField] {
	case "HMAC-SHA256":
		stringToSign := fmt.Sprintf("HMAC-SHA256,%s,%s,%s", id, values[digestEpochField], values[digestDataField])
		computedDigest = computeHMACSHA256(secret, values[digestEpochField], stringToSign)
	default:
		log.Println("validate.Digest(): missing digest type in header")
		return http.StatusUnauthorized, "Unauthorized"
	}

	if !timeWithinMinutes(values[digestEpochField], maxDigestAgeMinutes) {
		return http.StatusUnauthorized, "Digest Expired"
	}
	if computedDigest != values[digestField] {
		return http.StatusUnauthorized, "Digest Incorrect"
	}
	return http.StatusOK, "OK"

}

// computeHMACSHA256 computes a digest in an IoT friendly way, intended for
// small input data lengths. So we generate a large signing key.
func computeHMACSHA256(secret string, timestamp string, input string) string {

	signingKey := data.GenerateDigest(secret, timestamp)
	digest := data.GenerateDigest(signingKey, input)

	//log.Println("validate.computeHMACSHA256(): input", input)
	//log.Println("validate.computeHMACSHA256(): signingkey", signingKey)
	//log.Println("validate.computeHMACSHA256(): digest", digest)

	return digest
}
