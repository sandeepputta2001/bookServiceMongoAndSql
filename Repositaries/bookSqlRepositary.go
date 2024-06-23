package repositaries

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/sandeepputta2001/bookservicemongoandsql/interfaces"
	"github.com/sandeepputta2001/bookservicemongoandsql/models"
)


type BookRepositarySQL struct {
	db *sql.DB
}

func NewBookRepositarySQL(db *sql.DB) interfaces.BookRepositary {
	return &BookRepositarySQL{db: db}
}


func (r *BookRepositarySQL) FindAll() ([]models.Book , error) {

	var books []models.Book

	rows , err := r.db.Query("SELECT id , title , author , year_published FROM books")

	fmt.Print(rows)

	if err != nil {
		return nil , err
	}

	defer rows.Close()

	for rows.Next(){
		var book models.Book 

		err := rows.Scan(&book.ID , &book.Title , &book.Author,&book.Year_Published)

		if err != nil {
			return nil , err
		}

		books = append(books, book)

	}

	return books , nil
}

func (r *BookRepositarySQL) FindByID(id string ) (models.Book , error) {
	var book models.Book

	row  := r.db.QueryRow("SELECT id , title , author , year_published FROM books WHERE id = ? " , id)

	err := row.Scan(&book.ID, &book.Title , &book.Author,&book.Year_Published)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{} , errors.New("book not found")
		}
		return models.Book{} , err 
	}

	return book , nil 
}

func (r *BookRepositarySQL) Create(book models.Book) error { 
	_ , err := r.db.Exec("INSERT INTO books(id , title , author, year_published)  VALUES(? , ? , ?,?)" , book.ID , book.Title , book.Author,book.Year_Published)

	if err != nil {
		return err 
	}

	return nil 
}

func ( r *BookRepositarySQL) Update(book models.Book) error {
	_ , err := r.db.Exec("UPDATE books SET title=? , author=? , year_published=?  WHERE id=?",book.Title, book.Author ,book.Year_Published,  book.ID)
	if err !=nil {
		return err 
	}

	return nil 
}

func ( r *BookRepositarySQL) Delete(id string ) error {
	_ , err := r.db.Exec("DELETE FROM books WHERE id=?",id)

	if err != nil {
		return err 
	}

	return nil 
}  