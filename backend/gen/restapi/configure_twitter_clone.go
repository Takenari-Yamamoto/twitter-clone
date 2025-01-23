// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Takenari-Yamamoto/twitter-clone/gen/restapi/operations"
	"github.com/Takenari-Yamamoto/twitter-clone/gen/restapi/operations/auth"
	"github.com/Takenari-Yamamoto/twitter-clone/gen/restapi/operations/tweets"
)

//go:generate swagger generate server --target ../../gen --name TwitterClone --spec ../../api/swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.TwitterCloneAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TwitterCloneAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.BearerAuth == nil {
		api.BearerAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.TweetsDeleteTweetsTweetIDHandler == nil {
		api.TweetsDeleteTweetsTweetIDHandler = tweets.DeleteTweetsTweetIDHandlerFunc(func(params tweets.DeleteTweetsTweetIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tweets.DeleteTweetsTweetID has not yet been implemented")
		})
	}
	if api.TweetsDeleteTweetsTweetIDLikeHandler == nil {
		api.TweetsDeleteTweetsTweetIDLikeHandler = tweets.DeleteTweetsTweetIDLikeHandlerFunc(func(params tweets.DeleteTweetsTweetIDLikeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tweets.DeleteTweetsTweetIDLike has not yet been implemented")
		})
	}
	if api.TweetsGetTweetsHandler == nil {
		api.TweetsGetTweetsHandler = tweets.GetTweetsHandlerFunc(func(params tweets.GetTweetsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tweets.GetTweets has not yet been implemented")
		})
	}
	if api.TweetsGetTweetsTweetIDHandler == nil {
		api.TweetsGetTweetsTweetIDHandler = tweets.GetTweetsTweetIDHandlerFunc(func(params tweets.GetTweetsTweetIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tweets.GetTweetsTweetID has not yet been implemented")
		})
	}
	if api.TweetsGetUsersUserIDTweetsHandler == nil {
		api.TweetsGetUsersUserIDTweetsHandler = tweets.GetUsersUserIDTweetsHandlerFunc(func(params tweets.GetUsersUserIDTweetsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tweets.GetUsersUserIDTweets has not yet been implemented")
		})
	}
	if api.AuthPostAuthLoginHandler == nil {
		api.AuthPostAuthLoginHandler = auth.PostAuthLoginHandlerFunc(func(params auth.PostAuthLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthLogin has not yet been implemented")
		})
	}
	if api.AuthPostAuthSignupHandler == nil {
		api.AuthPostAuthSignupHandler = auth.PostAuthSignupHandlerFunc(func(params auth.PostAuthSignupParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthSignup has not yet been implemented")
		})
	}
	if api.TweetsPostTweetsHandler == nil {
		api.TweetsPostTweetsHandler = tweets.PostTweetsHandlerFunc(func(params tweets.PostTweetsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tweets.PostTweets has not yet been implemented")
		})
	}
	if api.TweetsPostTweetsTweetIDLikeHandler == nil {
		api.TweetsPostTweetsTweetIDLikeHandler = tweets.PostTweetsTweetIDLikeHandlerFunc(func(params tweets.PostTweetsTweetIDLikeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tweets.PostTweetsTweetIDLike has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
