package handler

import (
	"encoding/json"
	"net/http"
	"task-management-be/src/domain"
	"task-management-be/src/helper"

	"github.com/go-chi/chi/v5"
)

func (h *UserHandler) Update(writer http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)

	var requestUser domain.UserRequest

	decode := json.NewDecoder(req.Body)
	err := decode.Decode(&requestUser)
	if err != nil {
		writer.WriteHeader(helper.HandleStatusCode(err.Error()))
		encoder.Encode(&err)
		return
	}

	err = h.Service.Update(req.Context(), id, requestUser)

	if err != nil {
		writer.WriteHeader(helper.HandleStatusCode(err.Error()))
		encoder.Encode(&err)
		return
	}
	response := helper.GlobalResponse{
		Code:   200,
		Status: "Success update user",
	}

	encoder.Encode(&response)
}
