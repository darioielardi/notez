package models

// Note is the note model
type Note struct {
	Model
	Title string `json:"title"`
	
	// Relations
	UserID uint `json:"-"`
	User   User `json:"-"`
}
