package helper

import "net/http"

func ErrorNotNil(err error, message string) {
	if err != nil {
		Logger.Warn().Err(err).Msg(message)
		return
	}
}

var (
	ErrUserNotFound        = "Error user not found"
	ErrBadRequest          = "Bad request"
	ErrInternalServerError = "Internal Server Error"
	ErrJwtFailed           = "Error Jwt Failed"
)

type AppError struct {
	// Error      error  `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func NewAppError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: HandleStatusCode(message),
	}
}

func HandleStatusCode(message string) int {
	var statusCode int
	switch message {
	case ErrUserNotFound:
		statusCode = http.StatusNotFound
	case ErrBadRequest:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}
	return statusCode
}

func (ae *AppError) Error() string {
	return ae.Message
}
