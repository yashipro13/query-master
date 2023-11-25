package models

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"err"`
}

func (er *Error) Error() string {
	return er.Message
}
