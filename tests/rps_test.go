package tests

import (
	"github.com/stretchr/testify/require"
	"test_1/internal/domain"
	"test_1/internal/libs/ext_api"
	"test_1/internal/repository"
	"test_1/internal/server"
	service2 "test_1/internal/service"
	"testing"
)

func Test_RSP(t *testing.T) {
	takes := 11
	rps := 10
	countOverRPS := takes - rps

	repo := NewMockRepo()

	service := service2.NewService(repo)

	router, serv := httpServer(t)
	router.Use(ext_api.RPSLimiter())
	defer serv.Close()
	err := server.RegisterControllers(router, service)
	if err != nil {
		t.Error(err)
	}

	var count int
	type input struct {
		query  string
		result []domain.Book
	}

	tests := []struct {
		name  string
		input input
	}{
		{
			name: "rps test",
			input: input{
				query: "/books",
				result: []domain.Book{
					repository.BookOne, repository.BookTwo, repository.BookThree,
				},
			},
		},
	}

	for _, test := range tests {
		for i := 0; i < takes; i++ {

			t.Run(test.name, func(t *testing.T) {

				repo.GetBooksResult = test.input.result

				resp, err := serv.Client().Get(serv.URL + test.input.query)

				if resp.StatusCode != 200 {
					count++
				}
				t.Log(resp.StatusCode)
				require.NoError(t, err, "client request")

				defer resp.Body.Close()
			})
		}

		if count != countOverRPS {
			t.Errorf("rps count error: want %v but got %v", 1, count)
		}
	}
}
