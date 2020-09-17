package validators

import (
	"github.com/bpsaunders/user-api/models"
	"regexp"
)

const firstNameField = "first_name"
const lastNameField = "last_name"
const emailField = "email"
const countryField = "country"

var nameRegex = regexp.MustCompile("^[\\w'\\-,.][^0-9_!¡?÷¿/\\\\+=@#$%ˆ&*(){}|~<>;:[\\]]*$")
var emailRegex = regexp.MustCompile("^[\\w-.]+@([\\w-]+\\.)+[\\w-]{2,4}$")
var countryRegex = regexp.MustCompile("^[A-Z]{2}$")

// UserValidate provides an interface by which to validate a user
type UserValidate interface {
	Validate(rest *models.User) []ValidationError
}

// UserValidator implements the UserValidate interface
type UserValidator struct{}

// NewUserValidator returns a new concrete implementation of the UserValidate interface
func NewUserValidator() UserValidate {
	return &UserValidator{}
}

// Validate provides functionality with which to validate a user resource
func (*UserValidator) Validate(rest *models.User) []ValidationError {

	validationErrors := make([]ValidationError, 0)

	validateNameField(rest.FirstName, firstNameField, &validationErrors)
	validateNameField(rest.LastName, lastNameField, &validationErrors)
	validateEmail(rest.Email, &validationErrors)
	validateCountry(rest.Country, &validationErrors)

	return validationErrors
}

func validateNameField(name string, nameField string, validationErrors *[]ValidationError) {

	if name == "" {
		// Reject if name is blank
		*validationErrors = append(*validationErrors, newValidationError(jsonFieldPrefix+nameField, mandatoryElementMissing))
	} else if len(name) < 2 || len(name) > 30 {
		// Reject if name is fewer than chars, or longer than 30 chars
		params := map[string]interface{}{
			minChars: 2,
			maxChars: 30,
		}
		*validationErrors = append(*validationErrors, newValidationErrorWithParams(jsonFieldPrefix+nameField, invalidLength, params))
	} else if !nameRegex.MatchString(name) {
		// Reject if name contains improper characters
		*validationErrors = append(*validationErrors, newValidationError(jsonFieldPrefix+nameField, invalidChars))
	}
}

func validateEmail(email string, validationErrors *[]ValidationError) {

	if email == "" {
		// Reject if email is blank
		*validationErrors = append(*validationErrors, newValidationError(jsonFieldPrefix+emailField, mandatoryElementMissing))
	} else if len(email) > 120 {
		// Reject if email is longer than 120 chars
		params := map[string]interface{}{
			maxChars: 120,
		}
		*validationErrors = append(*validationErrors, newValidationErrorWithParams(jsonFieldPrefix+emailField, invalidLength, params))
	} else if !emailRegex.MatchString(email) {
		// Reject if email is not a valid format
		*validationErrors = append(*validationErrors, newValidationError(jsonFieldPrefix+emailField, invalidFormat))
	}
}

func validateCountry(country string, validationErrors *[]ValidationError) {

	if country == "" {
		// Reject if country is blank
		*validationErrors = append(*validationErrors, newValidationError(jsonFieldPrefix+countryField, mandatoryElementMissing))
	} else if !countryRegex.MatchString(country) {
		// Reject if country doesn't match a country code format
		// TODO: this is a rudimentary check - in future we should look to make sure this is a valid country code
		*validationErrors = append(*validationErrors, newValidationError(jsonFieldPrefix+countryField, invalidCountryCode))
	}
}
