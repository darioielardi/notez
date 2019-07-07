package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   interface{} `json:"error"`
}

func SendError(w http.ResponseWriter, err interface{}, status int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	value, isErr := err.(error)

	var r *ErrorResponse

	if isErr {
		r = &ErrorResponse{
			Success: false,
			Error:   value.Error(),
		}
	} else {
		r = &ErrorResponse{
			Success: false,
			Error:   err,
		}
	}

	json.NewEncoder(w).Encode(r)
}

func SendUnauthorizedErr(w http.ResponseWriter) {
	SendError(w, "Unauthorized", 401)
}

func SendInternalErr(w http.ResponseWriter) {
	SendError(w, "Internal Server Error", 500)
}

func SendForbiddenErr(w http.ResponseWriter) {
	SendError(w, "Forbidden", 403)
}

func SendNotFoundErr(w http.ResponseWriter, name string) {
	SendError(w, name+" not found", 404)
}
