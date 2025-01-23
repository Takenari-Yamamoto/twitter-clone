package main

import (
	"database/sql"
	"log"

	"github.com/Takenari-Yamamoto/twitter-clone/db"
	"github.com/Takenari-Yamamoto/twitter-clone/gen/restapi"
	"github.com/Takenari-Yamamoto/twitter-clone/gen/restapi/operations"
	"github.com/Takenari-Yamamoto/twitter-clone/gen/restapi/operations/auth"
	"github.com/Takenari-Yamamoto/twitter-clone/handler"
	"github.com/Takenari-Yamamoto/twitter-clone/service"
	"github.com/go-openapi/loads"
	"github.com/rs/cors"
)

func configureAPI(api *operations.TwitterCloneAPI, db *sql.DB) {
	// サービスとハンドラーの初期化
	authService := service.NewAuthService(db)
	if authService == nil {
		log.Fatal("Failed to create auth service")
	}

	authHandler := handler.NewAuthHandler(authService)
	if authHandler == nil {
		log.Fatal("Failed to create auth handler")
	}

	// ハンドラーの設定
	api.AuthPostAuthSignupHandler = auth.PostAuthSignupHandlerFunc(authHandler.HandleSignup)
	api.AuthPostAuthLoginHandler = auth.PostAuthLoginHandlerFunc(authHandler.HandleLogin)
}

func main() {
	log.Println("Starting server...")

	// データベース接続
	db, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// データベース接続の確認
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Connected to database")

	// Swagger仕様の読み込み
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// APIサーバーの初期化
	api := operations.NewTwitterCloneAPI(swaggerSpec)

	// APIの設定
	configureAPI(api, db)

	// サーバーの設定
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 8080
	server.Host = "0.0.0.0"

	// CORSの設定
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:3001"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          true,
	})

	// CORSミドルウェアを適用
	handler := corsMiddleware.Handler(api.Serve(nil))
	server.SetHandler(handler)

	log.Printf("Server is ready to handle requests at %s:%d", server.Host, server.Port)

	// サーバー起動
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
