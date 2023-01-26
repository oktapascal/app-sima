package exceptions

type ErrorUnprocessableEntity struct {
	ErrorMessage string `json:"error_message"`
	Param        string `json:"param"`
}

func (err ErrorUnprocessableEntity) Error() string {
	return err.ErrorMessage
}

func NewErrorUnprocessableEntity(error string, param string) ErrorUnprocessableEntity {
	return ErrorUnprocessableEntity{ErrorMessage: error, Param: param}
}
