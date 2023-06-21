package interfaces

import "net/http"

type IUserHandler interface {
	Create(writer http.ResponseWriter, request *http.Request)
}
