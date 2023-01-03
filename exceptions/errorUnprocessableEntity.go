package exceptions

type ErrorUnprocessableEntity struct {
	ErrorMessage string
	Param        string
}

func (err ErrorUnprocessableEntity) Error() string {
	return err.ErrorMessage
}

func NewErrorUnprocessableEntity(error string, param string) ErrorUnprocessableEntity {
	return ErrorUnprocessableEntity{ErrorMessage: error, Param: param}
}
