package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {

	tax, err := CalculateTax(1000.0)

	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0.0)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, tax)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repo := &TaxRepositoryMock{}
	repo.On("SaveTax", 10.0).Return(nil)
	repo.On("SaveTax", 0.0).Return(errors.New("amount must be greater than 0"))

	err := CalculateTaxAndSave(1000.0, repo)

	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repo)
	assert.Error(t, err, "amount must be greater than 0")
	repo.AssertExpectations(t)
	repo.AssertNumberOfCalls(t, "SaveTax", 2)
}
