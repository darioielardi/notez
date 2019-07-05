package auth

import (
	"net/http"
	
	"notez/core"
	"notez/models"
	"notez/utils/response"
)

func GetMe(_ *core.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		user := r.Context().Value("user").(*models.User)
		
		response.SendData(w, user, 200)
	}
}
