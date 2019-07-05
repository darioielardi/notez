package users

import (
	"net/http"
	
	"github.com/jinzhu/gorm"
	
	"notez/core"
	"notez/utils/params"
	"notez/utils/response"
)

func FindOne(s *core.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		id, err := params.GetIdParam(r)
		if err != nil {
			response.SendError(w, "Invalid user id", 400)
			return
		}
		
		user, err := NewRepository(s.DB).FindOne(*id)
		
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				response.SendError(w, err, 404)
				return
			}
			
			response.SendError(w, err, 500)
			return
		}
		
		response.SendData(w, user, 200)
	}
}
