package book

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Author string `json:"author"`
	Title  string `json:"title"`
	Rating int    `json:"rating"`
}

func (b Book) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Author, validation.Required),
		validation.Field(&b.Title, validation.Required),
		validation.Field(&b.Rating, validation.Required, validation.Max(5)),
	)
}
