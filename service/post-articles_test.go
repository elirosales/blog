package service_test

import (
	"fmt"
	"testing"

	"github.com/elizabethrosales/blog/service"
	"github.com/elizabethrosales/blog/testutils"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestPostArticle(t *testing.T) {
	s := testutils.NewSuite(t, "../testutils/testdata/test_articles.sql", "../.env")
	log := logrus.New()

	tests := []struct {
		name   string
		input  service.PostArticlesRequest
		output error
	}{
		{
			name: "success",
			input: service.PostArticlesRequest{
				Title:   "Hello World",
				Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
				Author:  "John Doe",
			},
		},
		{
			name: "success",
			input: service.PostArticlesRequest{
				Title: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			},
			output: fmt.Errorf("value too long"),
		},
	}

	for _, tt := range tests {
		s.Run(t, tt.name, func(t *testing.T) {
			svc := service.New(log, s.DB)
			resp, err := svc.PostArticle(tt.input)
			if err != nil {
				require.ErrorContains(t, err, tt.output.Error())
				return
			}

			require.NoError(t, err)
			require.NotEmpty(t, resp)
		})
	}
}
