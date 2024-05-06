package validator

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestCustomValidator_Validate(t *testing.T) {
	// Create a new instance of CustomValidator
	cv := NewValidator()

	// Test case 1: Valid input
	validInput := struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}{
		Name:  "John",
		Email: "john@example.com",
	}
	err := cv.Validate(validInput)
	assert.NoError(t, err)

	// Test case 2: Invalid input
	invalidInput := struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}{
		Name:  "John",
		Email: "invalidemail", // Invalid email format
	}
	err = cv.Validate(invalidInput)
	assert.Error(t, err)
	assert.IsType(t, &ErrorValidation{}, err)
}

func TestCustomValidator_ValidateStruct(t *testing.T) {
	// Create a new instance of CustomValidator
	cv := NewValidator()

	// Test case 1: Valid input
	validInput := struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}{
		Name:  "John",
		Email: "john@example.com",
	}
	err := cv.ValidateStruct(validInput)
	assert.NoError(t, err)

	// Test case 2: Invalid input
	invalidInput := struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}{
		Name:  "John",
		Email: "invalidemail", // Invalid email format
	}
	err = cv.ValidateStruct(invalidInput)
	assert.Error(t, err)
	assert.IsType(t, validator.ValidationErrors{}, err)
}
