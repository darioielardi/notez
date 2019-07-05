package notez

import (
	v "github.com/go-ozzo/ozzo-validation"
)

type CreateNoteDTO struct {
	Title string `json:"title"`
}

func (dto *CreateNoteDTO) Validate() error {
	return v.ValidateStruct(
		dto,
		v.Field(&dto.Title, v.Required, v.Length(1, 100)),
	)
}
