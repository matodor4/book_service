package tests

import (
	"github.com/stretchr/testify/require"
	"io"
	"test_1/internal/domain"
	"test_1/internal/libs/ext_api"
	"test_1/internal/server"
	service2 "test_1/internal/service"
	"testing"
	"time"
)

func Test_RSP(t *testing.T) {
	repo := NewMockRepo()

	service := service2.NewService(repo)

	router, serv := httpServer(t)
	router.Use(ext_api.RPSLimiter())
	defer serv.Close()
	err := server.RegisterControllers(router.Group("/v1"), service)
	if err != nil {
		t.Error(err)
	}

	var count int
	type input struct {
		query  string
		result []domain.Book
	}

	type output struct {
		//body string
		code int
	}

	tests := []struct {
		name  string
		input input
	}{
		{
			name: "rps test",
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
			},
		},
	}

	for _, test := range tests {
		for i := 0; i < 21; i++ {

			t.Run(test.name, func(t *testing.T) {

				repo.GetBooksResult = test.input.result

				resp, err := serv.Client().Get(serv.URL + test.input.query)
				_, err = io.ReadAll(resp.Body)
				if err != nil {
					t.Error(err)
				}
				if resp.StatusCode != 200 {
					count++
				}
				t.Log(resp.StatusCode)
				require.NoError(t, err, "client request")

				defer resp.Body.Close()
				time.Sleep(time.Second / 10)
			})
			t.Log("count ", count)
		}
	}
}
