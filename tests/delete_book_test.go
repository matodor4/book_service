package tests

import (
	"errors"
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

	err := server.RegisterControllers(router.Group("/v1"), service)
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
		//body string
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
				//body: `{"error":"no books find","code":"BAD_REQUEST"}`,
				code: http.StatusNotFound,
			},
		},
		{
			name: "ok request",
			input: input{
				query:  "/v1/book",
				result: req{id: 1},
			},
			output: output{
				//body: `{"error":"no books find","code":"BAD_REQUEST"}`,
				code: http.StatusOK,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo.GetBookError = test.input.svcErr
			//repo.GetBooksResult = test.input.result
			t.Log("URL", serv.URL+test.input.query)

			if err != nil {
				t.Fatal(err)
			}

			//resp, err := serv.Client().Do(req)
			//require.NoError(t, err, "client request")
			//
			//defer resp.Body.Close()
			//
			//assert.Equal(t, test.output.code, resp.StatusCode)
		})
	}
}
