package translateapp_test

import (
	"context"
	"errors"
	"testing"
	"translateapp/internal/logging"
	"translateapp/internal/translateapp"
	"translateapp/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	expected := translateapp.Response{
		Code:    200,
		Message: "success",
		Data: translateapp.ListLanguage{
			Languages: []translateapp.Language{
				{
					Name: "polish",
					Code: "pl",
				},
			},
		},
	}

	t.Run("ReturnsNoErrorOnSuccess", func(t *testing.T) {
		lt := mocks.LibreTranslator{}
		lt.On("GetLanguages", mock.Anything).Return(&expected.Data, nil)

		service := translateapp.NewService(logging.DefaultLogger().Desugar(), &lt)

		res, err := service.GetLanguages(context.Background())
		require.NoError(t, err)

		require.Equal(t, expected, *res)

		lt.AssertNumberOfCalls(t, "GetLanguages", 1)
	})

	t.Run("ReturnsErrorOnFailure", func(t *testing.T) {
		lt := mocks.LibreTranslator{}
		lt.On("GetLanguages", mock.Anything).Return(&expected.Data, errors.New("error"))

		service := translateapp.NewService(logging.DefaultLogger().Desugar(), &lt)

		res, err := service.GetLanguages(context.Background())
		require.Error(t, err)

		require.Nil(t, res)

		lt.AssertNumberOfCalls(t, "GetLanguages", 1)
	})
}