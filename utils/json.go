package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ErrorM struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
}

type ErrorResponse struct {
	Error ErrorM `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, s interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func Error(w http.ResponseWriter, statusCode int, err ErrorResponse) {
	JSON(w, statusCode, err)
}

func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	JSON(w, statusCode, data)
}

func DecodeJSON(w http.ResponseWriter, r *http.Request, s interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(s); err != nil {
		var syntexErr *json.SyntaxError
		var unmarshalErr *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntexErr):
			return fmt.Errorf("invalid json")
		case errors.As(err, &unmarshalErr):
			return fmt.Errorf("wrong type of fiels")
		}
	}
	return nil
}
