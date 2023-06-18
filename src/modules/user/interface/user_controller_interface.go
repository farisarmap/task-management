package interfaces

import "net/http"

type IUserController interface {
	Create(writer http.ResponseWriter, request *http.Request)
}
