package handler

import (
	"encoding/json"
	"net/http"
	"task-management-be/src/domain"
	"task-management-be/src/helper"
)

func (h *UserHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var requestUser domain.UserRequest

	decode := json.NewDecoder(req.Body)
	err := decode.Decode(&requestUser)

	helper.ErrorNotNil(err, "Failed to decode user data")

	userResponse, err := h.Service.Create(req.Context(), requestUser)
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
		Data:   userResponse,
	}

	encoder.Encode(response)
}
