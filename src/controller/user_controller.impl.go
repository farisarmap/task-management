package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-management-be/src/model/data/request"
	"task-management-be/src/model/data/response"
	"task-management-be/src/services"
)

type UserController struct {
	Service services.IUserService
}

func NewUserController(service services.IUserService) IUserController {
	return &UserController{
		Service: service,
	}
}

func (controller *UserController) Create(writer http.ResponseWriter, req *http.Request) {
	var requestUser request.UserRequest

	decode := json.NewDecoder(req.Body)
	err := decode.Decode(&requestUser)
	if err != nil {
		fmt.Println(err, "<< err decode")
		panic(err)
	}

	userResponse := controller.Service.CreateUser(req.Context(), requestUser)

	response := response.GlobalResponse{
		Code:   200,
		Status: "Success",
		Data:   userResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}
