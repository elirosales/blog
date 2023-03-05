package service_test

import (
	"fmt"
	"testing"

	"github.com/elizabethrosales/blog/service"
	"github.com/elizabethrosales/blog/testutils"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestGetArticle(t *testing.T) {
	s := testutils.NewSuite(t, "../testutils/testdata/test_articles.sql", "../.env")
	log := logrus.New()

	tests := []struct {
		name   string
		input  string
		output service.Article
		err    error
	}{
		{
			name:  "success",
			input: "d9af0398-a57f-48fc-88ee-9a8015206cf7",
			output: service.Article{
				ID:      "d9af0398-a57f-48fc-88ee-9a8015206cf7",
				Title:   "Hello World 1",
				Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
				Author:  "Meredith Grey",
			},
		},
		{
			name:  "not found",
			input: "invalid",
			err:   fmt.Errorf("record not found"),
		},
	}

	for _, tt := range tests {
		s.Run(t, tt.name, func(t *testing.T) {
			svc := service.New(log, s.DB)
			resp, err := svc.GetArticle(tt.input)
			if err != nil {
				require.Equal(t, tt.err, err)
				return
			}

			require.NoError(t, err)
			assertArticle(t, tt.output, resp[0])
		})
	}
}

func TestGetArticles(t *testing.T) {
	s := testutils.NewSuite(t, "../testutils/testdata/test_articles.sql", "../.env")
	log := logrus.New()

	tests := []struct {
		name   string
		output map[string]service.Article
		err    error
	}{
		{
			name: "success",
			output: map[string]service.Article{
				"d9af0398-a57f-48fc-88ee-9a8015206cf7": {
					ID:      "d9af0398-a57f-48fc-88ee-9a8015206cf7",
					Title:   "Hello World 1",
					Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
					Author:  "Meredith Grey",
				},
				"91e51c20-9e3b-4067-87df-3414717bef1e": {
					ID:      "91e51c20-9e3b-4067-87df-3414717bef1e",
					Title:   "Hello World 2",
					Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
					Author:  "Cristina Yang",
				},
				"ed0ade7c-f826-4448-aa60-9e1f41986e6a": {
					ID:      "ed0ade7c-f826-4448-aa60-9e1f41986e6a",
					Title:   "Hello World 3",
					Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
					Author:  "Alex Karev",
				},
				"0b503e50-9c74-4a96-9947-d0ea92d8a8f2": {
					ID:      "0b503e50-9c74-4a96-9947-d0ea92d8a8f2",
					Title:   "Hello World 4",
					Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
					Author:  "Lexie Grey",
				},
				"5d131ff7-527b-432c-a1ad-1ce0808e000d": {
					ID:      "5d131ff7-527b-432c-a1ad-1ce0808e000d",
					Title:   "Hello World 5",
					Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
					Author:  "Derek Shepherd",
				},
			},
		},
	}

	for _, tt := range tests {
		s.Run(t, tt.name, func(t *testing.T) {
			svc := service.New(log, s.DB)
			resp, err := svc.GetArticles()
			if err != nil {
				require.Equal(t, tt.err, err)
				return
			}

			require.NoError(t, err)
			for _, data := range resp {
				assertArticle(t, tt.output[data.ID], data)
			}
		})
	}
}

func assertArticle(t *testing.T, expect, actual service.Article) {
	require.NotEmpty(t, actual.ID)
	require.NotEmpty(t, actual.Title)
	require.NotEmpty(t, actual.Content)
	require.NotEmpty(t, actual.Author)

	require.Equal(t, expect.ID, actual.ID)
	require.Equal(t, expect.Title, actual.Title)
	require.Equal(t, expect.Content, actual.Content)
	require.Equal(t, expect.Author, actual.Author)
}
