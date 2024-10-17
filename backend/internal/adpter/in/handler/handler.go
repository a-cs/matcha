package handler

import (
	"backend/internal/adpter/dto"
	"backend/internal/application/domain/entity"
	"backend/internal/application/mapper"
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
	createUser := new(dto.CreateUserDto)
	if err := json.NewDecoder(ctx.Body).Decode(createUser); err != nil {
		resp := dto.ResponseDTO{Message: defines.CannotParseJSONRequest}
		utils.ServiceResponse(responseWriter, resp, http.StatusBadRequest)
		return
	}

	if !createUser.IsValid() {
		resp := dto.ResponseDTO{Message: defines.InvalidRequestBody}
		utils.ServiceResponse(responseWriter, resp, http.StatusBadRequest)
		return
	}

	createUserIntention := entity.CreateUserIntention{
		User: *mapper.ToCreateUser(createUser),
	}

	err := h.createUserUseCase.Execute(&createUserIntention)
	if err != nil {
		resp := dto.ResponseDTO{Message: err.Error()}
		utils.ServiceResponse(responseWriter, resp, http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusCreated)
}
