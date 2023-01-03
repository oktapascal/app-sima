package exceptions

type ErrorUnauthorized struct {
	ErrorMessage string
}

func (err ErrorUnauthorized) Error() string {
	return err.ErrorMessage
}

func NewErrorUnauthorized(error string) ErrorUnauthorized {
	return ErrorUnauthorized{ErrorMessage: error}
}
