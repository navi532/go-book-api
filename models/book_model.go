package models

type Book struct {
	ID        string `bson:"_id,omitempty"`
	Title     string `bson:"title,omitempty"`
	Author    string `bson:"author,omitempty"`
	PageCount int    `bson:"page_count,omitempty"`
}
