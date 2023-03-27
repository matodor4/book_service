package tests

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"test_1/internal/domain"
	"testing"
)

func httpServer(t *testing.T) (*gin.Engine, *httptest.Server) {
	t.Helper()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	server := httptest.NewServer(router)

	return router, server
}

type mockBookRepo struct {
	GetBooksResult []domain.Book
	GetBookError   error

	DeleteBooksError error
}

func NewMockRepo() *mockBookRepo {
	return &mockBookRepo{}
}

func (mock mockBookRepo) GetBooks(_ context.Context) ([]domain.Book, error) {
	return mock.GetBooksResult, mock.GetBookError
}

func (mock mockBookRepo) DeleteBookByID(_ context.Context, _ string) error {
	fmt.Println("DELETE MOCK")
	return mock.DeleteBooksError
}
