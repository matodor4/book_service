package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"test_1/internal/domain"
	"test_1/internal/server"
	service2 "test_1/internal/service"
	"testing"
	"time"
)

func Test_GetBooks(t *testing.T) {

	repo := NewMockRepo()

	service := service2.NewService(repo)

	router, serv := httpServer(t)
	defer serv.Close()
	err := server.RegisterControllers(router.Group("/v1"), service)
	if err != nil {
		t.Error(err)
	}

	type input struct {
		query  string
		result []domain.Book
		svcErr error
	}

	type output struct {
		//body string
		code int
	}

	tests := []struct {
		name   string
		input  input
		output output
	}{
		{
			name: "bad request - empty storage",
			input: input{
				query:  "/v1/books",
				result: nil,
				svcErr: errors.New("no books find"),
			},
			output: output{
				//body: `{"error":"no books find","code":"BAD_REQUEST"}`,
				code: http.StatusNotFound,
			},
		},
		{
			name: "ok request",
			input: input{
				query: "/v1/books",
				result: []domain.Book{
					{
						ID:            "1",
						Title:         "title_1",
						PublisherYear: time.Now(),
					},
					{
						ID:            "2",
						Title:         "title_2",
						PublisherYear: time.Now(),
					},
					{
						ID:            "3",
						Title:         "title_3",
						PublisherYear: time.Now(),
					},
				},
				svcErr: nil,
			},
			output: output{
				//body: ``,
				code: http.StatusOK,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo.GetBookError = test.input.svcErr
			repo.GetBooksResult = test.input.result

			resp, err := serv.Client().Get(serv.URL + test.input.query)

			require.NoError(t, err, "client request")

			defer resp.Body.Close()

			assert.Equal(t, test.output.code, resp.StatusCode)
		})
	}

}
