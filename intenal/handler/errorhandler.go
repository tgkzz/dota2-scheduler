package handler

import (
	"log"
	"strconv"
)

func ServerError(err error, status int) {
	log.Print(err.Error() + " " + strconv.Itoa(status))

	return
}
