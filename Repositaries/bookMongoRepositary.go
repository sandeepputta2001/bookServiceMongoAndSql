package repositaries

import (
	"context"
	"errors"
	"fmt"

	"github.com/sandeepputta2001/bookservicemongoandsql/interfaces"
	"github.com/sandeepputta2001/bookservicemongoandsql/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepositaryMongo struct {
	collection *mongo.Collection 
}

func NewBookRepositaryMongo(collection *mongo.Collection) interfaces.BookRepositary{ 
  return &BookRepositaryMongo{collection: collection} 
} 


func (b *BookRepositaryMongo) FindAll() ([]models.Book , error){
     
	var books []models.Book

	cursor , err := b.collection.Find(context.Background() , bson.D{})
	if err != nil {
		return nil , err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {

		var book models.Book

		if err := cursor.Decode(&book) ; err != nil {
			return nil ,err
		}

		books = append(books, book)

	}

	if err := cursor.Err(); err != nil  {
		return nil , err
	}

	return books, nil

}

func (b *BookRepositaryMongo) FindByID(id string ) (models.Book , error) {

	var book models.Book

	objID , err := primitive.ObjectIDFromHex(id) 

	if err != nil {
		return models.Book{} , err 
	}

	err = b.collection.FindOne(context.Background() , bson.M{"_id":objID}).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments{
			return models.Book{} , errors.New("book not found")
		}
		return models.Book{} , err
	}

	return book , nil
}

func (b *BookRepositaryMongo) Create(book models.Book) (error) { 

	_ , err := b.collection.InsertOne(context.Background() , book)

	if err != nil {
		return  err
	}

	return nil 
}

func (b *BookRepositaryMongo) Update(book models.Book) (error) { 
	
	updateId , err  := primitive.ObjectIDFromHex(fmt.Sprintf("%v",book.ID))

	if err != nil { 
		return err 
	}

	filter := bson.M{"id":updateId} 

	_ , err = b.collection.UpdateOne(context.Background() , filter , bson.D{
		{Key: "$set" , Value: bson.D{
			{Key: "title" , Value: book.Title},
			{Key: "author",Value: book.Author},
			{Key: "year_published",Value: book.Year_Published},
		}},
	} ,
  )
  
  if err != nil {
	return err 
  }

  return nil 
}

func (b *BookRepositaryMongo) Delete(id string ) (error) {
	objID , err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_ , err = b.collection.DeleteOne(context.Background() , bson.M{"_id":objID})

	if err != nil {
		return err 
	}

	return nil 
}