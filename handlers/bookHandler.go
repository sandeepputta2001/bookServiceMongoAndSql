package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/sandeepputta2001/bookservicemongoandsql/config"
	"github.com/sandeepputta2001/bookservicemongoandsql/interfaces"
	"github.com/sandeepputta2001/bookservicemongoandsql/models"
	"github.com/sandeepputta2001/bookservicemongoandsql/validators"
)


type bookHandler struct {
	service interfaces.BookService
}

func NewBookHandler(service interfaces.BookService) *bookHandler {
	return &bookHandler{service: service}
}


func (h *bookHandler) GetBooks(w http.ResponseWriter , r *http.Request) {
	books , err := h.service.GetBooks() 

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	json.NewEncoder(w).Encode(books)
} 

func (h *bookHandler) GetBook(w http.ResponseWriter , r *http.Request) {
	params := mux.Vars(r) 

	id := params["id"]
	book , err := h.service.GetBook(id)

	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func (b *bookHandler) CreateBook(w http.ResponseWriter , r *http.Request) { 

	var book models.Book 

	
	err := json.NewDecoder(r.Body).Decode(&book)
	config := config.GetConfig()

	if config.RepoType == "mongo" { 
      	book.ID = primitive.NewObjectID() 
	}

	if err != nil {
		http.Error(w , err.Error(), http.StatusBadRequest)
		return
	}

	if err := validators.ModelValidator(book) ; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    
	 err = b.service.CreateBook(book) 

	if err != nil {
		 http.Error(w, err.Error(),http.StatusInternalServerError)
		 return
	}
    
	json.NewEncoder(w).Encode(book) 
}

func (h *bookHandler) UpdateBook(w http.ResponseWriter , r *http.Request) { 

	var book models.Book 

	params := mux.Vars(r)

	id := params["id"]

	book.ID = id

	if err := json.NewDecoder(r.Body).Decode(&book) ; err != nil {
		http.Error(w , err.Error(), http.StatusBadRequest)
		return
	}

	if err := validators.ModelValidator(book) ; err != nil { 
		http.Error(w, err.Error() , http.StatusBadRequest)
		return 
	}

	if err := h.service.UpdateBook(book) ; err != nil {
		http.Error(w, err.Error() , http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book) 
}

func (h *bookHandler) DeleteBook(w http.ResponseWriter , r *http.Request) {

	params := mux.Vars(r)

	err := h.service.DeleteBook(params["id"])

	if err != nil {
		http.Error(w, err.Error() , http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)


} 