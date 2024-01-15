package handler

import (
	"dota2scheduler/intenal/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func ServerError(w http.ResponseWriter, err error, status int) {
	logMsg := err.Error() + " " + strconv.Itoa(status)

	errResponse := models.ErrResponse{
		Error:  logMsg,
		Status: status,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(errResponse); err != nil {
		log.Printf("Error sending response %s\n", err)
	}
}
