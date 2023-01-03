package exceptions

type ErrorBadRequest struct {
	ErrorMessage string
}

func (err ErrorBadRequest) Error() string {
	return err.ErrorMessage
}

func NewErrorBadRequest(error string) ErrorBadRequest {
	return ErrorBadRequest{ErrorMessage: error}
}
