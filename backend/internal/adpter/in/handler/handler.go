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
	"strings"
)

type Handler struct {
	createUserUseCase     in.CreateUser
	confirmAccountUseCase in.ConfirmAccount
	loginUseCase          in.Login
	getProfile            in.GetProfile
	updateProfile         in.UpdateProfile
}

func NewHandler(
	cu in.CreateUser,
	ca in.ConfirmAccount,
	l in.Login,
	gp in.GetProfile,
	up in.UpdateProfile,
) Handler {
	return Handler{
		createUserUseCase:     cu,
		confirmAccountUseCase: ca,
		loginUseCase:          l,
		getProfile:            gp,
		updateProfile:         up,
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

func (h Handler) Login(responseWriter http.ResponseWriter, ctx *http.Request) {
	login := new(dto.LoginDto)
	if err := json.NewDecoder(ctx.Body).Decode(login); err != nil {
		resp := dto.ResponseDTO{Message: defines.CannotParseJSONRequest}
		utils.ServiceResponse(responseWriter, resp, http.StatusBadRequest)
		return
	}

	if !login.IsValid() {
		resp := dto.ResponseDTO{Message: defines.InvalidRequestBody}
		utils.ServiceResponse(responseWriter, resp, http.StatusBadRequest)
		return
	}

	loginIntention := entity.LoginIntention{
		Email:    login.Email,
		Password: login.Password,
	}

	err := h.loginUseCase.Execute(&loginIntention)
	if err != nil {
		resp := dto.ResponseDTO{Message: err.Error()}
		utils.ServiceResponse(responseWriter, resp, http.StatusInternalServerError)
		return
	}

	resp := mapper.ToJwtResponseDto(&loginIntention.JwtResponse)
	utils.ServiceResponse(responseWriter, resp, http.StatusOK)
}

func (h Handler) GetProfile(responseWriter http.ResponseWriter, ctx *http.Request) {
	username := ctx.URL.Path[len("/profile/"):]
	if len(username) == 0 {
		resp := dto.ResponseDTO{Message: defines.UsernameMustToBeProvided}
		utils.ServiceResponse(responseWriter, resp, http.StatusBadRequest)
		return
	}

	if utils.IsSQLInjection(username) {
		resp := dto.ResponseDTO{Message: defines.InvalidRequestBody}
		utils.ServiceResponse(responseWriter, resp, http.StatusBadRequest)
		return
	}

	profileIntention := entity.ProfileIntention{
		Username: username,
	}

	err := h.getProfile.Execute(&profileIntention)
	if err != nil {
		resp := dto.ResponseDTO{Message: err.Error()}
		utils.ServiceResponse(responseWriter, resp, http.StatusInternalServerError)
		return
	}

	resp := mapper.FromProfileToProfileFrontDto(&profileIntention.Profile)
	utils.ServiceResponse(responseWriter, resp, http.StatusOK)
}

func (h Handler) UpdateProfile(responseWriter http.ResponseWriter, ctx *http.Request) {
	jwtToken := ctx.Header.Get("Authorization")
	if len(jwtToken) == 0 || !strings.HasPrefix(jwtToken, "Bearer ") ||
		len(jwtToken[len("Bearer "):]) == 0 || utils.IsSQLInjection(jwtToken) {
		resp := dto.ResponseDTO{Message: defines.AuthorizationMustToBeProvided}
		utils.ServiceResponse(responseWriter, resp, http.StatusUnauthorized)
		return
	}

	profile := new(dto.ProfileFrontDto)
	if err := json.NewDecoder(ctx.Body).Decode(profile); err != nil {
		resp := dto.ResponseDTO{Message: defines.CannotParseJSONRequest}
		utils.ServiceResponse(responseWriter, resp, http.StatusBadRequest)
		return
	}

	if !profile.IsValid() {
		resp := dto.ResponseDTO{Message: defines.InvalidRequestBody}
		utils.ServiceResponse(responseWriter, resp, http.StatusBadRequest)
		return
	}

	profileIntention := entity.ProfileIntention{
		ProfileRequest: *profile,
		JwtToken:       jwtToken[len("Bearer "):],
	}

	err := h.updateProfile.Execute(&profileIntention)
	if err != nil {
		resp := dto.ResponseDTO{Message: err.Error()}
		utils.ServiceResponse(responseWriter, resp, http.StatusInternalServerError)
		return
	}

	utils.ServiceResponse(responseWriter, profileIntention.ProfileRequest, http.StatusOK)
}
