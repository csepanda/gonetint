package apiv0

const (
    SERVER_SIDE_ERROR    = 0xBEAF
    CLIENT_REQUEST_ERROR = 0xDEAD
)

type serviceErrorType int
type serviceError     struct {
    errType serviceErrorType
    message string
}

func (e *serviceError) Error() string {
    return e.message
}

func serverError(msg string) error {
    return &serviceError{SERVER_SIDE_ERROR, msg}
}

func clientError(msg string) error {
    return &serviceError{CLIENT_REQUEST_ERROR, msg}
}
