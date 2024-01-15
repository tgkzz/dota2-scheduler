package models

import "errors"

type ErrResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

var ErrEmptyFilter error = errors.New("filters are empty or incorrect")
