package exceptions

type ErrorNotFound struct {
	ErrorMessage string
}

func (err ErrorNotFound) Error() string {
	return err.ErrorMessage
}

func NewErrorNotFound(error string) ErrorNotFound {
	return ErrorNotFound{ErrorMessage: error}
}
