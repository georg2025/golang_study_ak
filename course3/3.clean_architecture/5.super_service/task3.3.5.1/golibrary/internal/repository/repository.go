package repository

import (
	"database/sql"
	entities "golibrary/entities"

	_ "github.com/lib/pq"
)

type LibraryService struct {
	Base *sql.DB
}

type Librarer interface {
	AddAuthor(authorName string) error
	AddBook(book entities.Book) error
	AddUser(user entities.User) error
	TakeBook(userID, bookID int) error
	ReturnBook(book entities.Book) error
	GetAuthorByID(id int) (*entities.Author, error)
	GetAllUsers() ([]entities.User, error)
	GetAllBooksTakenByUser(userID int) ([]*entities.Book, error)
	GetAllAuthors() ([]entities.Author, error)
	GetAllAuthorBooks(authorID int) ([]*entities.Book, error)
	HowManyAuthorsExist() (int, error)
	HowManyBooksExist() (int, error)
	HowManyUsersExist() (int, error)
	GetBookInfo(id int) (entities.Book, error)
}

func NewLibraryService() *LibraryService {
	panic("implement me")
}

func (ls *LibraryService) AddAuthor(authorName string) error {
	panic("implement me")
}

func (ls *LibraryService) AddBook(book entities.Book) error {
	panic("implement me")
}

func (ls *LibraryService) AddUser(user entities.User) error {
	panic("implement me")
}

func (ls *LibraryService) TakeBook(userID, bookID int) error {
	panic("implement me")
}

func (ls *LibraryService) ReturnBook(book entities.Book) error {
	panic("implement me")
}

func (ls *LibraryService) GetAllUsers() ([]entities.User, error) {
	panic("implement me")
}

func (ls *LibraryService) GetAllBooksTakenByUser(userID int) ([]*entities.Book, error) {
	panic("implement me")
}

func (ls *LibraryService) GetAllAuthors() ([]entities.Author, error) {
	panic("implement me")
}

func (ls *LibraryService) GetAllAuthorBooks(authorID int) ([]*entities.Book, error) {
	panic("implement me")
}

func (ls *LibraryService) HowManyAuthorsExist() (int, error) {
	panic("implement me")
}

func (ls *LibraryService) HowManyBooksExist() (int, error) {
	panic("implement me")
}

func (ls *LibraryService) HowManyUsersExist() (int, error) {
	panic("implement me")
}

func (ls *LibraryService) GetAuthorByID(id int) (*entities.Author, error) {
	panic("implement me")
}

func (ls *LibraryService) GetBookInfo(id int) (entities.Book, error) {
	panic("implement me")
}
