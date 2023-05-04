package repo

import (
	"context"
	"errors"
	"go-mongo/configs"
	"go-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type BookRepo interface {
	CreateBook(*models.Book) error
	UpdateBook(primitive.ObjectID, *models.Book) error
	DeleteBook(primitive.ObjectID) (*models.Book, error)
	DeleteMultipleBooks([]primitive.ObjectID) error
	GetBookById(primitive.ObjectID) (*models.Book, error)
	ListBooks() (*[]models.Book, error)
	UpsertBook(*models.Book) error
	CreateManyBooks([]interface{}) error
}

type MongoBookRepo struct {
	collection *mongo.Collection
}

func NewMongoBookRepo() *MongoBookRepo {
	return &MongoBookRepo{collection: configs.GetCollection("books")}
}

func (r *MongoBookRepo) CreateBook(book *models.Book) error {

	_, err := r.collection.InsertOne(context.Background(), book)
	if err != nil {
		return errors.Join(errors.New("MONGO: Failed to insert new book"), err)
	}

	return nil
}

func (r *MongoBookRepo) CreateManyBooks(books []interface{}) error {

	_, err := r.collection.InsertMany(context.Background(), books)

	if err != nil {
		return errors.Join(errors.New("MONGO: Failed to insert books"), err)
	}

	return nil
}

func (r *MongoBookRepo) UpdateBook(id primitive.ObjectID, book *models.Book) error {

	result, err := r.collection.UpdateByID(context.Background(), id, bson.D{{"$set", book}})
	if err != nil {
		return errors.New("MONGO: Failed to update the book")
	}
	if result.MatchedCount == 0 {
		return errors.New("MONGO: ID doesn't exists")
	}

	return nil
}

func (r *MongoBookRepo) DeleteBook(id primitive.ObjectID) (*models.Book, error) {

	book := &models.Book{}
	err := r.collection.FindOneAndDelete(context.Background(), bson.D{{"_id", id}}).Decode(book)

	if err != nil {
		return book, errors.New("MONGO: Failed to delete the book")
	}

	return book, nil
}

func (r *MongoBookRepo) GetBookById(id primitive.ObjectID) (*models.Book, error) {

	result := &models.Book{}

	err := r.collection.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(result)

	if err != nil {
		return result, errors.Join(errors.New("MONGO: Failed to find the book"), err)
	}

	return result, nil
}

func (r *MongoBookRepo) ListBooks() (*[]models.Book, error) {

	books := &[]models.Book{}

	cursor, err := r.collection.Find(context.Background(), bson.D{})

	defer cursor.Close(context.Background())

	if err != nil {
		return nil, errors.New("MONGO: Failed to fetch book records from DB")
	}

	if err := cursor.All(context.Background(), books); err != nil {
		return nil, errors.Join(errors.New("MONGO: Failed to fetch book records"), err)
	}

	return books, nil
}

func (r *MongoBookRepo) UpsertBook(book *models.Book) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := models.Book{
		Author: book.Author,
		Title:  book.Title,
	}

	update := bson.D{{"$set", book}}

	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		return errors.New("MONGO: Failed to upsert")
	}
	return nil
}

func (r *MongoBookRepo) DeleteMultipleBooks(ids []primitive.ObjectID) error {

	filter := bson.D{{"_id", bson.D{{"$in", ids}}}}
	result, err := r.collection.DeleteMany(context.Background(), filter)
	if err != nil || result.DeletedCount == 0 {
		return errors.New("MONGO: Failed to delete the book")
	}

	return nil
}
