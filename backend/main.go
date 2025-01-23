package main

import (
	"log"

	"github.com/Takenari-Yamamoto/twitter-clone/gen/restapi"
	"github.com/Takenari-Yamamoto/twitter-clone/gen/restapi/operations"
	"github.com/go-openapi/loads"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTwitterCloneAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 8080

	// ハンドラーの設定
	// TODO: データベース接続とサービスの初期化
	// authHandler := handler.NewAuthHandler(authService)
	// api.AuthPostAuthSignupHandler = auth.PostAuthSignupHandlerFunc(authHandler.HandleSignup)
	// api.AuthPostAuthLoginHandler = auth.PostAuthLoginHandlerFunc(authHandler.HandleLogin)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
} 