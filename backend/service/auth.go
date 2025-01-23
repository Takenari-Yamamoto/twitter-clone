package service

import (
	"context"
	"errors"
	"time"

	"github.com/Takenari-Yamamoto/twitter-clone/config"
	"github.com/Takenari-Yamamoto/twitter-clone/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db boil.ContextExecutor
}

func NewAuthService(db boil.ContextExecutor) *AuthService {
	return &AuthService{db: db}
}

type SignupInput struct {
	Username string
	Email    string
	Password string
}

func (s *AuthService) Signup(ctx context.Context, input SignupInput) (*models.User, error) {
	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
	}

	err = user.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return user, nil
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginResponse struct {
	User  *models.User
	Token string
}

func (s *AuthService) Login(ctx context.Context, input LoginInput) (*LoginResponse, error) {
	user, err := models.Users(models.UserWhere.Email.EQ(input.Email)).One(ctx, s.db)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// パスワードの検証
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// JWTトークンの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		User:  user,
		Token: tokenString,
	}, nil
} 