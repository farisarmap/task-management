package handler

import (
	"encoding/json"
	"net/http"
	"task-management-be/src/helper"
	"task-management-be/src/modules/user/model/data/request"
	"task-management-be/src/modules/user/model/data/response"
)

func (controller *UserHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var requestUser request.UserRequest

	decode := json.NewDecoder(req.Body)
	err := decode.Decode(&requestUser)

	helper.ErrorNotNil(err, "Failed to decode user data")

	userResponse := controller.Service.Create(req.Context(), requestUser)

	response := response.GlobalResponse{
		Code:   200,
		Status: "Success",
		Data:   userResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}
