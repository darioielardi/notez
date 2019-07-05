package notez

import (
	"github.com/jinzhu/gorm"
	
	"notez/models"
)

type Note models.Note

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// FindOne gets a note by id
func (nr *Repository) FindOne(id int) (*Note, error) {
	
	note := &Note{}
	
	if err := nr.db.First(note, id).Error; err != nil {
		return nil, err
	}
	
	return note, nil
}

func (nr *Repository) FindByUserId(userId uint) (*[]Note, error) {
	
	notes := &[]Note{}
	
	if err := nr.db.Where("user_id = ?", userId).Find(&notes).Error; err != nil {
		return nil, err
	}
	
	return notes, nil
}

func (nr *Repository) Create(dto *CreateNoteDTO, user *models.User) (*Note, error) {
	
	note := &Note{
		Title: dto.Title,
		User:  *user,
	}
	
	if err := nr.db.Create(note).Error; err != nil {
		return nil, err
	}
	
	return note, nil
}
