package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	repositaries "github.com/sandeepputta2001/bookservicemongoandsql/Repositaries"
	services "github.com/sandeepputta2001/bookservicemongoandsql/Services"
	"github.com/sandeepputta2001/bookservicemongoandsql/config"
	"github.com/sandeepputta2001/bookservicemongoandsql/handlers"
	"github.com/sandeepputta2001/bookservicemongoandsql/helpers"
	"github.com/sandeepputta2001/bookservicemongoandsql/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {  


	config := config.GetConfig()

	var booksRepo interfaces.BookRepositary

	if config.RepoType == "mongo" { 
    
	clientOptions := options.Client().ApplyURI(config.MongoURI) 
	client  , err  := mongo.Connect(context.Background() , clientOptions) 
	log.Print("connected to mongo database ")

	if err != nil { 
		 log.Fatal(err) 
	}

	booksCollection := client.Database(config.DatabaseName).Collection("books")

	booksRepo = repositaries.NewBookRepositaryMongo(booksCollection) 

    } 	else if config.RepoType == "sql" { 
	  sqlDB , err := sql.Open("mysql",fmt.Sprintf("sandeep:sandeepputta@tcp(localhost:3306)/%s",config.DatabaseName)) 
	 

	  if err != nil {
		log.Fatal(err )
	  }

	err = sqlDB.Ping()

	if err != nil {
		sqlDB.Close()
		log.Fatal(err)
	  }

	log.Print("connected  to sql db") 

	tableExists := helpers.TableExists(sqlDB,"books")

	if !tableExists { 
		booksTable := `
	  CREATE TABLE books ( 
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL,
		year_published INT NOT NULL
	  )
	`
	_ , err = sqlDB.Exec(booksTable)

	if err != nil {
		log.Fatal("error creating books table")
	}

	log.Print("Books table created successfully")

	} else {
		log.Print("Table books already existed")
	}

	

	booksRepo = repositaries.NewBookRepositarySQL(sqlDB)

    } else {
		log.Fatal("unsupported repositary type") 
	}  
	
	bookservice := services.NewBookService(booksRepo) 

	bookHandler := handlers.NewBookHandler(bookservice) 

	

	r := mux.NewRouter()

	r.HandleFunc("/books",bookHandler.GetBooks).Methods("GET") 
	r.HandleFunc("/books/{id}",bookHandler.GetBook).Methods("GET")
	r.HandleFunc("/books" , bookHandler.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}",bookHandler.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/{id}",bookHandler.UpdateBook).Methods("PUT")

	fmt.Println("server is listening at port 5003") 


	log.Fatal(http.ListenAndServe("localhost:5003", r))


}  