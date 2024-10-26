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
	createUserUseCase     in.CreateUser
	confirmAccountUseCase in.ConfirmAccount
}

func NewHandler(
	cu in.CreateUser,
	ca in.ConfirmAccount,
) Handler {
	return Handler{
		createUserUseCase:     cu,
		confirmAccountUseCase: ca,
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
		CreateUser: *mapper.ToCreateUser(createUser),
	}

	err := h.createUserUseCase.Execute(&createUserIntention)
	if err != nil {
		resp := dto.ResponseDTO{Message: err.Error()}
		utils.ServiceResponse(responseWriter, resp, http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusCreated)
}

func (h Handler) ConfirmAccount(responseWriter http.ResponseWriter, ctx *http.Request) {
	confirmAccount := new(dto.ConfirmAccountDto)
	if err := json.NewDecoder(ctx.Body).Decode(confirmAccount); err != nil {
		resp := dto.ResponseDTO{Message: defines.CannotParseJSONRequest}
		utils.ServiceResponse(responseWriter, resp, http.StatusBadRequest)
		return
	}

	if !confirmAccount.IsValid() {
		resp := dto.ResponseDTO{Message: defines.InvalidRequestBody}
		utils.ServiceResponse(responseWriter, resp, http.StatusBadRequest)
		return
	}

	confirmAccountIntention := entity.ConfirmAccountIntention{
		SlugID: confirmAccount.SlugID,
	}

	err := h.confirmAccountUseCase.Execute(&confirmAccountIntention)
	if err != nil {
		resp := dto.ResponseDTO{Message: err.Error()}
		utils.ServiceResponse(responseWriter, resp, http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}
