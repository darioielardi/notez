package notez

import (
	"net/http"
	
	"github.com/jinzhu/gorm"
	
	"notez/core"
	"notez/models"
	"notez/utils/enums"
	"notez/utils/params"
	res "notez/utils/response"
)

func GetByUser(s *core.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		user := r.Context().Value("user").(*models.User)
		
		notes, err := NewRepository(s.DB).FindByUserId(user.ID)
		if err != nil {
			res.SendError(w, err, 500)
			return
		}
		
		res.SendData(w, notes, 201)
		
	}
}

func GetOne(s *core.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		id, err := params.GetIdParam(r)
		if err != nil {
			res.SendError(w, "Invalid note id", 400)
			return
		}
		
		note, err := NewRepository(s.DB).FindOne(*id)
		
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				res.SendError(w, err, 404)
				return
			}
			
			res.SendError(w, err, 500)
			return
		}
		
		user := r.Context().Value("user").(*models.User)
		
		if user.Role != enums.Admin && note.UserID != user.ID {
			res.SendError(w, "Forbidden", 403)
			return
		}
		
		res.SendData(w, note, 200)
		
	}
}

func CreateNew(s *core.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		dto := new(CreateNoteDTO)
		
		if err := s.V.ValidateBody(dto, r); err != nil {
			res.SendError(w, err, 400)
			return
		}
		
		repo := NewRepository(s.DB)
		
		user := r.Context().Value("user").(*models.User)
		
		_, err := repo.Create(dto, user)
		if err != nil {
			res.SendError(w, err, 500)
			return
		}
		
		notes, err := repo.FindByUserId(user.ID)
		if err != nil {
			res.SendError(w, err, 500)
			return
		}
		
		res.SendData(w, notes, 201)
	}
}
