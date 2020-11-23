package main

import (
	"github.com/idthings/idengine/internal/webserver"
	"log"
)

func main() {

	log.Println("apisvc.main()")

	webserver.Start()
}
