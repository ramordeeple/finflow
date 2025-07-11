package utils

import "log"

func ErrorHandler(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
