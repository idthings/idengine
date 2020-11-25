package validate

import (
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"
	"time"
)

func extractGUID(urlPath string) (string, error) {

	// extract guid from http request, the last element
	elements := strings.Split(urlPath, "/") // always returns at least one element
	id := elements[len(elements)-1]

	_, err := uuid.Parse(id)
	if err != nil {
		log.Println("validate.extractGUID(): urlPath has no valid guid")
		return "", err
	}

	return id, nil
}

func timeWithinMinutes(t string, sinceMinutes int) bool {

	inputTime, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return false
	}

	timestamp := time.Now().UnixNano() / 1e6

	diff := timestamp - inputTime

	if int(diff) < sinceMinutes*60 {
		return true
	}

	return false
}
