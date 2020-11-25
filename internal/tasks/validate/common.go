package validate

import (
	"github.com/google/uuid"
	"log"
	"strings"
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
