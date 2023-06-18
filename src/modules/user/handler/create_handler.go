package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-management-be/src/modules/user/model/data/request"
	"task-management-be/src/modules/user/model/data/response"
)

func (controller *UserController) Create(writer http.ResponseWriter, req *http.Request) {
	var requestUser request.UserRequest

	decode := json.NewDecoder(req.Body)
	err := decode.Decode(&requestUser)
	if err != nil {
		fmt.Println(err, "<< err decode")
		panic(err)
	}

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
