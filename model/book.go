package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Book struct {
	ID        int       `json:"id" gorm:"primaryKey;type:serial"`
	Name      string    `json:"name_book" gorm:"type:varchar(50)"`
	Author    string    `json:"author" gorm:"type:varchar(50)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//DeletedAt time.Time `json:"deleted_at"`
}

func (e Book) Validation() error { // custom validation
	return validation.ValidateStruct(&e,
		validation.Field(&e.Name, validation.Required),
		validation.Field(&e.Author, validation.Required))
}
