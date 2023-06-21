package handler

import (
	"encoding/json"
	"net/http"
	"task-management-be/src/modules/user/model/data/response"

	"github.com/go-chi/chi/v5"
)

func (handler *UserHandler) FindById(writer http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	user := handler.Service.FindById(req.Context(), id)
	response := response.GlobalResponse{
		Code:   200,
		Status: "Success",
		Data:   user,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}
