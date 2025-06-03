package errors

type CustomError interface {
	error
	Code() string
}

type ServiceError struct {
	code    string
	message string
	err     error
}

func (e *ServiceError) Error() string {
	if e.err != nil {
		return e.message + ": " + e.err.Error()
	}
	return e.message
}

func (e *ServiceError) Code() string {
	return e.code
}

func (e *ServiceError) Unwrap() error {
	return e.err
}
