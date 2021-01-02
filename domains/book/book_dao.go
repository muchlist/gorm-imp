package book

import (
	"github.com/muchlist/gorm-imp/database"
)

var (
	BookDao bookDaoInterface = &bookDao{}
)

type bookDao struct{}

type bookDaoInterface interface {
	GetBooks() ([]Book, error)
	CreateBook(data Book) (Book, error)
	GetBookByID(id int) (Book, error)
	DeleteBook(id int) (int64, error)
	UpdateBook(data Book) (Book, error)
}

func (b *bookDao) GetBooks() ([]Book, error) {
	db := database.DbConn
	var books []Book
	db.Find(&books)

	return books, nil
}

func (b *bookDao) CreateBook(data Book) (Book, error) {
	db := database.DbConn
	var book Book = data
	result := db.Create(&book)

	// book.ID             // returns inserted data's primary key
	// result.Error        // returns error
	// result.RowsAffected // returns inserted records count

	return book, result.Error
}

func (b *bookDao) GetBookByID(id int) (Book, error) {
	db := database.DbConn
	var book Book
	db.First(&book, id)

	return book, nil
}


func (b *bookDao) DeleteBook(id int) (int64, error) {
	db := database.DbConn
	var book Book
	deletedRow := db.Delete(&book, id).RowsAffected
	return deletedRow, nil
}

func (b *bookDao) UpdateBook(data Book) (Book, error) {
	db := database.DbConn
	var book Book
	db.First(&book, data.ID)

	book.Author = data.Author
	book.Title = data.Title
	book.Rating = data.Rating

	db.Save(book)

	return book, nil
}