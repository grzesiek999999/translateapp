package client

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"translateapp/internal/logging"
	"translateapp/internal/translateapp"
)

func TestClient(t *testing.T) {
	expected := translateapp.ListLanguage{
		Languages: []translateapp.Language{
			{
				Name: "English",
				Code: "en",
			},
			{
				Name: "Polish",
				Code: "pl",
			},
		},
	}

	t.Run("ReturnsNoErrorOnSuccess", func(t *testing.T) {
		client := NewLibreTranslateClient(logging.DefaultLogger().Desugar())
		res, err := client.GetLanguages(context.Background())
		require.NoError(t, err)
		require.Equal(t, expected, *res)

	})
}
