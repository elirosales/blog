package models_test

import (
	"fmt"
	"testing"

	"github.com/elizabethrosales/blog/database/models"
	"github.com/elizabethrosales/blog/testutils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestInsertArticle(t *testing.T) {
	s := testutils.NewSuite(t, "../../testutils/testdata/test_articles.sql", "../../.env")

	tests := []struct {
		name   string
		input  models.Article
		output error
	}{
		{
			name: "success",
			input: models.Article{
				UUID:    uuid.NewString(),
				Title:   "Hello World",
				Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
				Author:  "John Doe",
			},
		},
		{
			name: "invalid input",
			input: models.Article{
				UUID: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			},
			output: fmt.Errorf("value too long"),
		},
	}

	for _, tt := range tests {
		s.Run(t, tt.name, func(t *testing.T) {
			a := models.NewArticle(s.DB)
			err := a.InsertArticle(tt.input)
			if err != nil {
				require.ErrorContains(t, err, tt.output.Error())
				return
			}
			require.NoError(t, err)

			resp, err := a.GetArticle(tt.input.UUID)
			require.NoError(t, err)
			assertArticle(t, tt.input, resp)
		})
	}
}

func TestGetArticle(t *testing.T) {
	s := testutils.NewSuite(t, "../../testutils/testdata/test_articles.sql", "../../.env")

	tests := []struct {
		name   string
		input  string
		output models.Article
		err    error
	}{
		{
			name:  "success",
			input: "d9af0398-a57f-48fc-88ee-9a8015206cf7",
			output: models.Article{
				UUID:    "d9af0398-a57f-48fc-88ee-9a8015206cf7",
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
			a := models.NewArticle(s.DB)
			resp, err := a.GetArticle(tt.input)
			if err != nil {
				require.Equal(t, tt.err, err)
				return
			}

			require.NoError(t, err)
			assertArticle(t, tt.output, resp)
		})
	}
}

func TestGetArticles(t *testing.T) {
	s := testutils.NewSuite(t, "../../testutils/testdata/test_articles.sql", "../../.env")

	tests := []struct {
		name   string
		output map[string]models.Article
		err    error
	}{
		{
			name: "success",
			output: map[string]models.Article{
				"d9af0398-a57f-48fc-88ee-9a8015206cf7": {
					UUID:    "d9af0398-a57f-48fc-88ee-9a8015206cf7",
					Title:   "Hello World 1",
					Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
					Author:  "Meredith Grey",
				},
				"91e51c20-9e3b-4067-87df-3414717bef1e": {
					UUID:    "91e51c20-9e3b-4067-87df-3414717bef1e",
					Title:   "Hello World 2",
					Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
					Author:  "Cristina Yang",
				},
				"ed0ade7c-f826-4448-aa60-9e1f41986e6a": {
					UUID:    "ed0ade7c-f826-4448-aa60-9e1f41986e6a",
					Title:   "Hello World 3",
					Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
					Author:  "Alex Karev",
				},
				"0b503e50-9c74-4a96-9947-d0ea92d8a8f2": {
					UUID:    "0b503e50-9c74-4a96-9947-d0ea92d8a8f2",
					Title:   "Hello World 4",
					Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
					Author:  "Lexie Grey",
				},
				"5d131ff7-527b-432c-a1ad-1ce0808e000d": {
					UUID:    "5d131ff7-527b-432c-a1ad-1ce0808e000d",
					Title:   "Hello World 5",
					Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
					Author:  "Derek Shepherd",
				},
			},
		},
	}

	for _, tt := range tests {
		s.Run(t, tt.name, func(t *testing.T) {
			a := models.NewArticle(s.DB)
			resp, err := a.GetArticles()
			if err != nil {
				require.Equal(t, tt.output, err)
				return
			}

			require.NoError(t, err)
			for _, data := range resp {
				assertArticle(t, tt.output[data.UUID], data)
			}
		})
	}
}

func assertArticle(t *testing.T, expect, actual models.Article) {
	require.NotEmpty(t, actual.UUID)
	require.NotEmpty(t, actual.Title)
	require.NotEmpty(t, actual.Content)
	require.NotEmpty(t, actual.Author)

	require.Equal(t, expect.UUID, actual.UUID)
	require.Equal(t, expect.Title, actual.Title)
	require.Equal(t, expect.Content, actual.Content)
	require.Equal(t, expect.Author, actual.Author)
}
