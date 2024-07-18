package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)

	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0.0)

	assert.Error(t, err, "amount must be greater than zeros")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "greater than zeros")
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil).Once()
	repository.On("SaveTax", 0.0).Return(errors.New("amount must be greater than zeros")).Once()
	//repository.On("SaveTax", mock.Anything).Return(errors.New("amount must be greater than zeros"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "amount must be greater than zeros")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
