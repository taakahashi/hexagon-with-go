package application_test

import (
	"github.com/taakahashi/hexagon-with-go/application"
	"github.com/stretchr/testify/require"
	"testing"
)


func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = DISABLED
	product.Price = 10

	err := product.Enabled()
	require.Nil(t, err)
}