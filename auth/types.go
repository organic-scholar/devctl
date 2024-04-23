package auth

import (
	"encoding/json"
)

type Error struct {
	ErrorMessage     string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func UnmarshalErrorResponse(data []byte) (Error, error) {
	var r Error
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Error) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (e Error) Error() string {
	return e.ErrorMessage
}
