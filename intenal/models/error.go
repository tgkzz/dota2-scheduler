package models

type ErrResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
