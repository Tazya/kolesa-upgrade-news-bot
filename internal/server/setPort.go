package server

import (
	"log"
	"os"
	"strconv"
)

func SetPort() string {

	port := ""

	for _, arg := range os.Args {
		if arg[:5] == "port=" {
			port = arg[5:]
		}
	}

	if port == "" {
		port = "8888"
	}

	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return port
}
