package book_services

import (
	"errors"
	"github.com/muchlist/gorm-imp/domains/book"
)

var (
	BookService bookServiceInterface = &bookService{}
)

type bookService struct{}

type bookServiceInterface interface {
	GetBooks() ([]book.Book, error)
	GetBookByID(id int) (book.Book, error)
	CreateBook(data book.Book) (book.Book, error)
	DeleteBook(id int) error
	UpdateBook(data book.Book) (book.Book, error)
}

func (b *bookService) GetBooks() ([]book.Book, error) {
	return book.BookDao.GetBooks()
}

func (b *bookService) GetBookByID(id int) (book.Book, error) {
	return book.BookDao.GetBookByID(id)
}

func (b *bookService) CreateBook(data book.Book) (book.Book, error) {
	return book.BookDao.CreateBook(data)
}

func (b *bookService) DeleteBook(id int) error {
	deletedRow, _ := book.BookDao.DeleteBook(id)
	if deletedRow == 0 {
		return errors.New("Data tidak ditemukan")
	}
	return nil
}

func (b *bookService) UpdateBook(data book.Book) (book.Book, error) {
	return book.BookDao.UpdateBook(data)
}