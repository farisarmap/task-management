package interfaces

import "net/http"

type IUserHandler interface {
	Create(writer http.ResponseWriter, request *http.Request)
	FindById(writer http.ResponseWriter, req *http.Request)
	List(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
}
