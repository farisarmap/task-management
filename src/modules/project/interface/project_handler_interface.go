package interfaces

import "net/http"

type IProjectHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}
