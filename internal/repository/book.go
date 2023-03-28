package repository

import (
	"context"
	"errors"
	"sync"
	"test_1/internal/domain"
)

// BookList for test too
var bookList = map[string]domain.Book{
	"1": BookOne,
	"2": BookTwo,
	"3": BookThree,
}

var BookOne = domain.Book{
	ID:            "1",
	Title:         "The Hitchhiker's Guide to the Galaxy",
	Author:        "Douglas Adams",
	PublisherYear: "1979",
}
var BookTwo = domain.Book{
	ID:            "2",
	Title:         "Pride and Prejudice",
	Author:        "Jane Austen",
	PublisherYear: "1979",
}
var BookThree = domain.Book{
	ID:            "3",
	Title:         "To Kill a Mockingbird",
	Author:        "Harper Lee",
	PublisherYear: "1960",
}

var storage = SafeKeyValueStore{
	data: bookList,
	mux:  &sync.Mutex{},
}

type BookRepo struct {
	storage SafeKeyValueStore
}

func New() *BookRepo {
	r := &BookRepo{}
	r.storage = storage
	return r
}

func (rep *BookRepo) GetBooks(_ context.Context) ([]domain.Book, error) {
	var books = make([]domain.Book, 0, len(rep.storage.data))

	for _, book := range rep.storage.data {
		books = append(books, book)
	}
	return books, nil
}

func (rep *BookRepo) DeleteBookByID(_ context.Context, id string) error {
	if _, find := rep.storage.data[id]; find {
		rep.storage.Remove(id)
		return nil
	}

	return errors.New("book not find")
}
