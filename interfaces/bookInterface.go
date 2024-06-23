package interfaces

import (
	"github.com/sandeepputta2001/bookservicemongoandsql/models"
)

type BookRepositary interface {
	FindAll() ([]models.Book , error)
	FindByID(id string ) (models.Book , error)
	Create(book models.Book) (error)
	Delete(id string ) (error)
	Update(book models.Book) (error)

}


type BookService interface { 
	GetBooks() ([]models.Book , error)
	GetBook(id string ) (models.Book , error)
	CreateBook(book models.Book) (error)
	DeleteBook(id string ) (error)
	UpdateBook(book models.Book) (error)
}