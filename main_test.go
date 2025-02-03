package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseURL(t *testing.T) {
	t.Run("Successfully parses URL", func(t *testing.T) {
		testURL := "http://localhost:3000/products"
		expPath := "/products"

		actPath := parseURL(testURL)
		require.Equal(t, expPath, actPath)
	})
}

func TestFetchProductsDataHandler(t *testing.T) {
	t.Run("Successfully fetches products from dummyjson", func(t *testing.T) {
		products := fetchProductsDataHandler()
		actLen := len(products.Products)
		require.Greater(t, actLen, 0)
		require.NotEmpty(t, products.Products)
	})
}
