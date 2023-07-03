package handler

import (
	"encoding/json"
	"net/http"
	"task-management-be/src/helper"
)

func (h *UserHandler) List(writer http.ResponseWriter, request *http.Request) {
	users, err := h.Service.List(request.Context())

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
		Data:   users,
	}

	encoder.Encode(response)
}
