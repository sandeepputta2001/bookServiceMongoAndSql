package models


type Book struct {
	ID              interface{}         `json:"id,omitempty" bson:"_id,omitempty"`
	Title           string              `json:"title" bson:"title" validate:"required"`
	Author          string              `json:"author" bson:"author" validate:"required"`
	Year_Published  int                 `json:"year_published" bson:"year_published" validate:"required"`
} 

