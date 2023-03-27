package repository

import (
	"context"
	"errors"
	"sync"
	"test_1/internal/domain"
	"time"
)

var bookList = map[string]domain.Book{
	"1": {
		ID:            "1",
		Title:         "title_1",
		PublisherYear: time.Now(),
	},
	"2": {
		ID:            "2",
		Title:         "title_2",
		PublisherYear: time.Now(),
	},
	"3": {
		ID:            "3",
		Title:         "title_3",
		PublisherYear: time.Now(),
	},
}

var storage = SafeKeyValueStore{
	data: bookList,
	mux:  sync.Mutex{},
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
