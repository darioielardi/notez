package response

import (
	"encoding/json"
	"net/http"
)

type DataResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
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
