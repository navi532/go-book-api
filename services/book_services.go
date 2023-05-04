package services

import (
	"errors"
	"go-mongo/controllers/requests/book"
	"go-mongo/models"
	"go-mongo/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUpdateBookService(request *requests.CreateUpdateBookRequest) error {
	var r repo.BookRepo = repo.NewMongoBookRepo()
	return r.UpsertBook(
		&models.Book{
			Title:     request.Title,
			Author:    request.Author,
			PageCount: request.PageCount,
		})
}

func CreateBookService(request *requests.CreateRequest) error {
	var r repo.BookRepo = repo.NewMongoBookRepo()
	return r.CreateBook(&models.Book{
		Title:     request.Title,
		Author:    request.Author,
		PageCount: request.PageCount,
	})
}

func CreateManyBookService(request *[]requests.CreateRequest) error {
	var r repo.BookRepo = repo.NewMongoBookRepo()

	var books []interface{}

	for _, eachRequest := range *request {
		books = append(books, models.Book{
			Title:     eachRequest.Title,
			Author:    eachRequest.Author,
			PageCount: eachRequest.PageCount,
		})
	}

	return r.CreateManyBooks(books)
}

func UpdateBookService(request *requests.UpdateRequest, id string) error {
	var r repo.BookRepo = repo.NewMongoBookRepo()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ERROR: Invalid Object ID")
	}
	return r.UpdateBook(objID, &models.Book{Title: request.Title, Author: request.Author, PageCount: request.PageCount})
}

func GetBookByIdService(id string) (*models.Book, error) {
	var r repo.BookRepo = repo.NewMongoBookRepo()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ERROR: Invalid Object ID")
	}

	return r.GetBookById(objID)
}

func ListBookService() (*[]models.Book, error) {
	var r repo.BookRepo = repo.NewMongoBookRepo()
	return r.ListBooks()
}

func DeleteBookService(id string) (*models.Book, error) {
	var r repo.BookRepo = repo.NewMongoBookRepo()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ERROR: Invalid Object ID")
	}
	return r.DeleteBook(objID)
}

func DeleteMultipleBookService(request *requests.DeleteBookRequest) error {
	var r repo.BookRepo = repo.NewMongoBookRepo()

	var ids []primitive.ObjectID

	for _, val := range request.IDS {
		objID, err := primitive.ObjectIDFromHex(val)
		if err != nil {
			return errors.New("ERROR: Invalid Object ID")
		}
		ids = append(ids, objID)
	}

	return r.DeleteMultipleBooks(ids)

}
