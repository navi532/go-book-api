package models

type Book struct {
	ID        string   `bson:"_id,omitempty"`
	Title     string   `bson:"title,omitempty"`
	Authors   []Author `bson:"authors,omitempty"`
	PageCount int      `bson:"page_count,omitempty"`
}
