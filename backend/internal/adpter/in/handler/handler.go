package handler

import (
	"backend/internal/adpter/dto"
	"backend/internal/application/port/in"
	"backend/internal/defines"
	"backend/internal/utils"
	"encoding/json"
	"net/http"
)

type Handler struct {
	createUserUseCase in.CreateUser
}

func NewHandler(
	cu in.CreateUser,
) Handler {
	return Handler{
		createUserUseCase: cu,
	}
}

func (h Handler) CreateUser(responseWriter http.ResponseWriter, ctx *http.Request) {
	createUser := new(dto.CreateUserDTO)
	if err := json.NewDecoder(ctx.Body).Decode(createUser); err != nil {
		resp := dto.ResponseDTO{Message: defines.InvalidRequestBody}
		utils.ServiceResponse(responseWriter, &resp, http.StatusBadRequest)
		return
	}
	responseWriter.WriteHeader(http.StatusCreated)
}
