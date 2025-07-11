package utils

import "log"

func FatalErr(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
