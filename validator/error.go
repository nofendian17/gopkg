package validator

type ErrorValidation []FieldError

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func (e *ErrorValidation) Error() string {
	return "error validation"
}
