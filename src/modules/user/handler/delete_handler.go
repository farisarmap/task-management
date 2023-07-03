package handler

import (
	"encoding/json"
	"net/http"
	"task-management-be/src/helper"

	"github.com/go-chi/chi/v5"
)

func (h *UserHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	err := h.Service.Delete(request.Context(), id)
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	if err != nil {
		writer.WriteHeader(helper.HandleStatusCode(err.Error()))
		encoder.Encode(&err)
		return
	}
	response := helper.GlobalResponse{
		Code:   200,
		Status: "Success",
	}

	encoder.Encode(&response)
}
