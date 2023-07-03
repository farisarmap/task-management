package handler

import (
	"encoding/json"
	"net/http"
	"task-management-be/src/helper"

	"github.com/go-chi/chi/v5"
)

func (h *UserHandler) FindById(writer http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	user, err := h.Service.FindById(req.Context(), id)
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)

	if err != nil {
		writer.WriteHeader(helper.HandleStatusCode(err.Error()))
		encoder.Encode(err)
		return
	}
	response := helper.GlobalResponse{
		Code:   200,
		Status: "Success",
		Data:   user,
	}

	encoder.Encode(response)
}
