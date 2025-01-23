package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Takenari-Yamamoto/twitter-clone/config"
	"github.com/Takenari-Yamamoto/twitter-clone/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	fmt.Println("Signup service called")
	// ユーザー名の重複チェック
	exists, err := models.Users(qm.Where("username = ?", input.Username)).Exists(ctx, s.db)
	if err != nil {
		fmt.Printf("Error checking username existence: %v\n", err)
		return nil, err
	}
	if exists {
		fmt.Println("Username already exists:", input.Username)
		return nil, errors.New("username already exists")
	}
	fmt.Println("Username check passed")

	// メールアドレスの重複チェック
	exists, err = models.Users(qm.Where("email = ?", input.Email)).Exists(ctx, s.db)
	if err != nil {
		fmt.Printf("Error checking email existence: %v\n", err)
		return nil, err
	}
	if exists {
		fmt.Println("Email already exists:", input.Email)
		return nil, errors.New("email already exists")
	}
	fmt.Println("Email check passed")

	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error hashing password: %v\n", err)
		return nil, err
	}

	user := &models.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
	}
	fmt.Println("User object created")

	err = user.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		fmt.Printf("Error inserting user: %v\n", err)
		return nil, err
	}
	fmt.Printf("User inserted successfully with ID: %d\n", user.ID)

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
	user, err := models.Users(qm.Where("email = ?", input.Email)).One(ctx, s.db)
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

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
