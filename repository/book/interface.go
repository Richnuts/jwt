package book

import "sirclo/entities"

type Book interface {
	GetAllBook() ([]entities.Book, error)
	GetBook(id int) (entities.Book, error)
	CreateBook(book entities.Book) error
	DeleteBook(id int) error
	EditBook(book entities.Book, id int) error
	RentBook(bookId int, userId int) error
	ReturnBook(bookId int, userId int) error
}
