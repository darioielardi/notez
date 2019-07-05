package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Success bool        `json:"success""`
	Error   interface{} `json:"error"`
}

type DataResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
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

func SendData(w http.ResponseWriter, data interface{}, status int) {
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	
	r := &DataResponse{
		Success: true,
		Data:    data,
	}
	
	json.NewEncoder(w).Encode(r)
}
