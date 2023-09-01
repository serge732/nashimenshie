package apperror

type AppError struct {
	Err    error
	Custom string
}

func NewAppError(err error, custom string) *AppError {
	return &AppError{
		Err:    err,
		Custom: custom,
	}
}

func (e *AppError) Error() string {
	return e.Custom
}

func (e *AppError) Unwrap() error {
	return e.Err
}
