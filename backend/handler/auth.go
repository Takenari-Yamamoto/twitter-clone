package handler

import (
	"github.com/Takenari-Yamamoto/twitter-clone/gen/models"
	"github.com/Takenari-Yamamoto/twitter-clone/gen/restapi/operations/auth"
	"github.com/Takenari-Yamamoto/twitter-clone/service"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) HandleSignup(params auth.PostAuthSignupParams) middleware.Responder {
	user, err := h.authService.Signup(params.HTTPRequest.Context(), service.SignupInput{
		Username: *params.Body.Username,
		Email:    *params.Body.Email,
		Password: *params.Body.Password,
	})
	if err != nil {
		return auth.NewPostAuthSignupBadRequest().WithPayload(&models.Error{
			Code:    400,
			Message: err.Error(),
		})
	}

	return auth.NewPostAuthSignupCreated().WithPayload(&models.User{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		DisplayName: *user.DisplayName.Ptr(),
		Bio:         *user.Bio.Ptr(),
		AvatarURL:   *user.AvatarURL.Ptr(),
		CreatedAt:   strfmt.DateTime(user.CreatedAt),
		UpdatedAt:   strfmt.DateTime(user.UpdatedAt),
	})
}

func (h *AuthHandler) HandleLogin(params auth.PostAuthLoginParams) middleware.Responder {
	response, err := h.authService.Login(params.HTTPRequest.Context(), service.LoginInput{
		Email:    *params.Body.Email,
		Password: *params.Body.Password,
	})
	if err != nil {
		return auth.NewPostAuthLoginUnauthorized().WithPayload(&models.Error{
			Code:    401,
			Message: "Invalid credentials",
		})
	}

	return auth.NewPostAuthLoginOK().WithPayload(&auth.PostAuthLoginOKBody{
		Token: response.Token,
		User: &models.User{
			ID:          response.User.ID,
			Username:    response.User.Username,
			Email:       response.User.Email,
			DisplayName: *response.User.DisplayName.Ptr(),
			Bio:         *response.User.Bio.Ptr(),
			AvatarURL:   *response.User.AvatarURL.Ptr(),
			CreatedAt:   strfmt.DateTime(response.User.CreatedAt),
			UpdatedAt:   strfmt.DateTime(response.User.UpdatedAt),
		},
	})
} 