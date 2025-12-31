package application_test

import (
	"go-hexagonal/src/application"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}

	product.Name = "Gas"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "ice cream"
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}

	product.ID = uuid.NewV4().String()
	product.Name = "Milk Shake"
	product.Status = application.DISABLED
	product.Price = 10.0

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())

}
