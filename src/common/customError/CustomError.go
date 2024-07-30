package customerError

type CustomError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error implements error.
func (c *CustomError) Error() string {
	panic("unimplemented")
}

type CustomerErrEnum string

const (
	UNAUTHORIZE      CustomerErrEnum = "UNAUTHORIZE"
	INVALID_TOKEN    CustomerErrEnum = "INVALID_TOKEN"
	INVALID_PASSWORD CustomerErrEnum = "INVALID_PASSWORD"
	INVALID_FIELD    CustomerErrEnum = "INVALID_FIELD"
	CANNT_CREATE_DOC CustomerErrEnum = "CANNT_CREATE_DOC"
)

func ResponseError(code CustomerErrEnum, message string) interface{} {
	return map[string]interface{}{
		"code":    code,
		"message": message,
	}
}
