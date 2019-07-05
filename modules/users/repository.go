package users

import (
	"github.com/jinzhu/gorm"
	
	"notez/models"
)

type User models.User

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// FindOne finds a user by id
func (ur *Repository) FindOne(id int) (*User, error) {
	
	user := &User{}
	
	if err := ur.db.First(user, id).Error; err != nil {
		return nil, err
	}
	
	return user, nil
}

func (ur *Repository) FindByAuthId(authId string) (*User, error) {
	
	user := &User{}
	
	if err := ur.db.Where("auth_id = ?", authId).First(user).Error; err != nil {
		return nil, err
	}
	
	return user, nil
}
