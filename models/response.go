package models

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   *Error      `json:"error"`
}
