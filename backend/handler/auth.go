package handler

import (
	"fmt"

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
	if authService == nil {
		panic("authService must not be nil")
	}
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) HandleSignup(params auth.PostAuthSignupParams) middleware.Responder {
	fmt.Println("Signup called")
	if h.authService == nil {
		return auth.NewPostAuthSignupBadRequest().WithPayload(&models.Error{
			Code:    500,
			Message: "Internal server error: auth service not initialized",
		})
	}

	if params.Body.Username == nil || params.Body.Email == nil || params.Body.Password == nil {
		return auth.NewPostAuthSignupBadRequest().WithPayload(&models.Error{
			Code:    400,
			Message: "Invalid request: missing required fields",
		})
	}

	user, err := h.authService.Signup(params.HTTPRequest.Context(), service.SignupInput{
		Username: *params.Body.Username,
		Email:    *params.Body.Email,
		Password: *params.Body.Password,
	})
	fmt.Println("User created:", user)

	if err != nil {
		return auth.NewPostAuthSignupBadRequest().WithPayload(&models.Error{
			Code:    400,
			Message: err.Error(),
		})
	}

	// nilチェックを追加
	displayName := ""
	if user.DisplayName.Valid {
		displayName = user.DisplayName.String
	}

	bio := ""
	if user.Bio.Valid {
		bio = user.Bio.String
	}

	avatarURL := ""
	if user.AvatarURL.Valid {
		avatarURL = user.AvatarURL.String
	}

	return auth.NewPostAuthSignupCreated().WithPayload(&models.User{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		DisplayName: displayName,
		Bio:         bio,
		AvatarURL:   avatarURL,
		CreatedAt:   strfmt.DateTime(user.CreatedAt),
		UpdatedAt:   strfmt.DateTime(user.UpdatedAt),
	})
}

func (h *AuthHandler) HandleLogin(params auth.PostAuthLoginParams) middleware.Responder {
	fmt.Println("Login called")
	if h.authService == nil {
		return auth.NewPostAuthLoginUnauthorized().WithPayload(&models.Error{
			Code:    500,
			Message: "Internal server error: auth service not initialized",
		})
	}

	if params.Body.Email == nil || params.Body.Password == nil {
		return auth.NewPostAuthLoginUnauthorized().WithPayload(&models.Error{
			Code:    400,
			Message: "Invalid request: missing required fields",
		})
	}

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

	if response == nil || response.User == nil {
		return auth.NewPostAuthLoginUnauthorized().WithPayload(&models.Error{
			Code:    500,
			Message: "Internal server error: invalid response",
		})
	}

	// nilチェックを追加
	displayName := ""
	if response.User.DisplayName.Valid {
		displayName = response.User.DisplayName.String
	}

	bio := ""
	if response.User.Bio.Valid {
		bio = response.User.Bio.String
	}

	avatarURL := ""
	if response.User.AvatarURL.Valid {
		avatarURL = response.User.AvatarURL.String
	}

	return auth.NewPostAuthLoginOK().WithPayload(&auth.PostAuthLoginOKBody{
		Token: response.Token,
		User: &models.User{
			ID:          response.User.ID,
			Username:    response.User.Username,
			Email:       response.User.Email,
			DisplayName: displayName,
			Bio:         bio,
			AvatarURL:   avatarURL,
			CreatedAt:   strfmt.DateTime(response.User.CreatedAt),
			UpdatedAt:   strfmt.DateTime(response.User.UpdatedAt),
		},
	})
}
