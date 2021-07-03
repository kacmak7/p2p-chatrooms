package utils

import (
	"log"
)

const ERROR_ID := "Error P2P@CHATROOMS"

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s %s: %s", ERROR_ID, msg, err.Error())
	}
}
