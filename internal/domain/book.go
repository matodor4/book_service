package domain

import (
	"time"
)

type Book struct {
	ID            string
	Title         string
	author        string
	PublisherYear time.Time
}
