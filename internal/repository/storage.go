package repository

import (
	"sync"
	"test_1/internal/domain"
)

type SafeKeyValueStore struct {
	data map[string]domain.Book
	mux  sync.Mutex
}

func (store *SafeKeyValueStore) Put(key string, value domain.Book) {
	store.mux.Lock()
	store.data[key] = value
	store.mux.Unlock()
}

func (store *SafeKeyValueStore) Get(key string) domain.Book {
	return store.data[key]
}

func (store *SafeKeyValueStore) Remove(key string) {
	store.mux.Lock()
	delete(store.data, key)
	store.mux.Unlock()
}
