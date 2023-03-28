package tests

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"test_1/internal/server"
	service2 "test_1/internal/service"
	"testing"
)

func Test_DeleteBook(t *testing.T) {

	repo := NewMockRepo()

	service := service2.NewService(repo)

	router, serv := httpServer(t)
	defer serv.Close()

	err := server.RegisterControllers(router, service)
	if err != nil {
		t.Error(err)
	}

	type req struct {
		id int
	}

	type input struct {
		query  string
		result req
		svcErr error
	}

	type output struct {
		code int
	}

	tests := []struct {
		name   string
		input  input
		output output
	}{
		{
			name: "bad request - book not find",
			input: input{
				query:  "/v1/book",
				result: req{id: 1},
				svcErr: errors.New("no books find"),
			},
			output: output{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "ok request",
			input: input{
				query:  "/v1/book",
				result: req{id: 1},
				svcErr: nil,
			},
			output: output{
				code: http.StatusNoContent,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo.DeleteBooksError = test.input.svcErr

			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest(http.MethodDelete, serv.URL+"/book", bytes.NewReader([]byte(`{"id": 1}`)))
			if err != nil {
				t.Error(err)
			}
			resp, err := serv.Client().Do(req)
			require.NoError(t, err, "client request")

			defer resp.Body.Close()

			assert.Equal(t, test.output.code, resp.StatusCode)
		})
	}
}
