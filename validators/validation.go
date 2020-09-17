package validators

const jsonFieldPrefix = "$."

const mandatoryElementMissing = "mandatory_element_missing"
const invalidLength = "invalid_length"
const invalidChars = "invalid_characters"
const invalidFormat = "invalid_format"
const invalidCountryCode = "invalid_country_code"

const minChars = "min_chars"
const maxChars = "max_chars"

// ValidationError holds details of any validation errors
type ValidationError struct {
	Field  string                 `json:"field"`
	Error  string                 `json:"error"`
	Params map[string]interface{} `json:"params,omitempty"`
}

func newValidationError(field string, error string) ValidationError {
	return ValidationError{
		Field: field,
		Error: error,
	}
}

func newValidationErrorWithParams(field string, error string, params map[string]interface{}) ValidationError {
	return ValidationError{
		Field:  field,
		Error:  error,
		Params: params,
	}
}
