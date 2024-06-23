package services

import (
	"github.com/sandeepputta2001/bookservicemongoandsql/interfaces"
	"github.com/sandeepputta2001/bookservicemongoandsql/models"
)


type bookService struct {  
	repo interfaces.BookRepositary 
}

func NewBookService(repo interfaces.BookRepositary) interfaces.BookService{ 
	return &bookService{repo: repo}
} 


func (s *bookService) GetBooks() ([]models.Book , error) {
	return  s.repo.FindAll()

}

func (s *bookService) GetBook(id string) (models.Book , error){
	return s.repo.FindByID(id)
}


func (s *bookService) CreateBook(book models.Book) (error) {
	return s.repo.Create(book)
}

func (s *bookService) DeleteBook(id string ) (error) {
	return s.repo.Delete(id)
}

func (s *bookService) UpdateBook(book models.Book) (error) {
	return s.repo.Update(book)
}

