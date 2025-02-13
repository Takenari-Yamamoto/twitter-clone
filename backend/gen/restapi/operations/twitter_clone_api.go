// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/Takenari-Yamamoto/twitter-clone/gen/restapi/operations/auth"
	"github.com/Takenari-Yamamoto/twitter-clone/gen/restapi/operations/tweets"
)

// NewTwitterCloneAPI creates a new TwitterClone instance
func NewTwitterCloneAPI(spec *loads.Document) *TwitterCloneAPI {
	return &TwitterCloneAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		TweetsDeleteTweetsTweetIDHandler: tweets.DeleteTweetsTweetIDHandlerFunc(func(params tweets.DeleteTweetsTweetIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tweets.DeleteTweetsTweetID has not yet been implemented")
		}),
		TweetsDeleteTweetsTweetIDLikeHandler: tweets.DeleteTweetsTweetIDLikeHandlerFunc(func(params tweets.DeleteTweetsTweetIDLikeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tweets.DeleteTweetsTweetIDLike has not yet been implemented")
		}),
		TweetsGetTweetsHandler: tweets.GetTweetsHandlerFunc(func(params tweets.GetTweetsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tweets.GetTweets has not yet been implemented")
		}),
		TweetsGetTweetsTweetIDHandler: tweets.GetTweetsTweetIDHandlerFunc(func(params tweets.GetTweetsTweetIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tweets.GetTweetsTweetID has not yet been implemented")
		}),
		TweetsGetUsersUserIDTweetsHandler: tweets.GetUsersUserIDTweetsHandlerFunc(func(params tweets.GetUsersUserIDTweetsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tweets.GetUsersUserIDTweets has not yet been implemented")
		}),
		AuthPostAuthLoginHandler: auth.PostAuthLoginHandlerFunc(func(params auth.PostAuthLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthLogin has not yet been implemented")
		}),
		AuthPostAuthSignupHandler: auth.PostAuthSignupHandlerFunc(func(params auth.PostAuthSignupParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthSignup has not yet been implemented")
		}),
		TweetsPostTweetsHandler: tweets.PostTweetsHandlerFunc(func(params tweets.PostTweetsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tweets.PostTweets has not yet been implemented")
		}),
		TweetsPostTweetsTweetIDLikeHandler: tweets.PostTweetsTweetIDLikeHandlerFunc(func(params tweets.PostTweetsTweetIDLikeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tweets.PostTweetsTweetIDLike has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*TwitterCloneAPI Twitter Clone API仕様 */
type TwitterCloneAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// BearerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	BearerAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// TweetsDeleteTweetsTweetIDHandler sets the operation handler for the delete tweets tweet ID operation
	TweetsDeleteTweetsTweetIDHandler tweets.DeleteTweetsTweetIDHandler
	// TweetsDeleteTweetsTweetIDLikeHandler sets the operation handler for the delete tweets tweet ID like operation
	TweetsDeleteTweetsTweetIDLikeHandler tweets.DeleteTweetsTweetIDLikeHandler
	// TweetsGetTweetsHandler sets the operation handler for the get tweets operation
	TweetsGetTweetsHandler tweets.GetTweetsHandler
	// TweetsGetTweetsTweetIDHandler sets the operation handler for the get tweets tweet ID operation
	TweetsGetTweetsTweetIDHandler tweets.GetTweetsTweetIDHandler
	// TweetsGetUsersUserIDTweetsHandler sets the operation handler for the get users user ID tweets operation
	TweetsGetUsersUserIDTweetsHandler tweets.GetUsersUserIDTweetsHandler
	// AuthPostAuthLoginHandler sets the operation handler for the post auth login operation
	AuthPostAuthLoginHandler auth.PostAuthLoginHandler
	// AuthPostAuthSignupHandler sets the operation handler for the post auth signup operation
	AuthPostAuthSignupHandler auth.PostAuthSignupHandler
	// TweetsPostTweetsHandler sets the operation handler for the post tweets operation
	TweetsPostTweetsHandler tweets.PostTweetsHandler
	// TweetsPostTweetsTweetIDLikeHandler sets the operation handler for the post tweets tweet ID like operation
	TweetsPostTweetsTweetIDLikeHandler tweets.PostTweetsTweetIDLikeHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *TwitterCloneAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *TwitterCloneAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *TwitterCloneAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TwitterCloneAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TwitterCloneAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TwitterCloneAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TwitterCloneAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TwitterCloneAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TwitterCloneAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TwitterCloneAPI
func (o *TwitterCloneAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BearerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.TweetsDeleteTweetsTweetIDHandler == nil {
		unregistered = append(unregistered, "tweets.DeleteTweetsTweetIDHandler")
	}
	if o.TweetsDeleteTweetsTweetIDLikeHandler == nil {
		unregistered = append(unregistered, "tweets.DeleteTweetsTweetIDLikeHandler")
	}
	if o.TweetsGetTweetsHandler == nil {
		unregistered = append(unregistered, "tweets.GetTweetsHandler")
	}
	if o.TweetsGetTweetsTweetIDHandler == nil {
		unregistered = append(unregistered, "tweets.GetTweetsTweetIDHandler")
	}
	if o.TweetsGetUsersUserIDTweetsHandler == nil {
		unregistered = append(unregistered, "tweets.GetUsersUserIDTweetsHandler")
	}
	if o.AuthPostAuthLoginHandler == nil {
		unregistered = append(unregistered, "auth.PostAuthLoginHandler")
	}
	if o.AuthPostAuthSignupHandler == nil {
		unregistered = append(unregistered, "auth.PostAuthSignupHandler")
	}
	if o.TweetsPostTweetsHandler == nil {
		unregistered = append(unregistered, "tweets.PostTweetsHandler")
	}
	if o.TweetsPostTweetsTweetIDLikeHandler == nil {
		unregistered = append(unregistered, "tweets.PostTweetsTweetIDLikeHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TwitterCloneAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TwitterCloneAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "Bearer":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.BearerAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *TwitterCloneAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *TwitterCloneAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *TwitterCloneAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TwitterCloneAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the twitter clone API
func (o *TwitterCloneAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TwitterCloneAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/tweets/{tweetId}"] = tweets.NewDeleteTweetsTweetID(o.context, o.TweetsDeleteTweetsTweetIDHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/tweets/{tweetId}/like"] = tweets.NewDeleteTweetsTweetIDLike(o.context, o.TweetsDeleteTweetsTweetIDLikeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tweets"] = tweets.NewGetTweets(o.context, o.TweetsGetTweetsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tweets/{tweetId}"] = tweets.NewGetTweetsTweetID(o.context, o.TweetsGetTweetsTweetIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{userId}/tweets"] = tweets.NewGetUsersUserIDTweets(o.context, o.TweetsGetUsersUserIDTweetsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/login"] = auth.NewPostAuthLogin(o.context, o.AuthPostAuthLoginHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/signup"] = auth.NewPostAuthSignup(o.context, o.AuthPostAuthSignupHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tweets"] = tweets.NewPostTweets(o.context, o.TweetsPostTweetsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tweets/{tweetId}/like"] = tweets.NewPostTweetsTweetIDLike(o.context, o.TweetsPostTweetsTweetIDLikeHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TwitterCloneAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *TwitterCloneAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *TwitterCloneAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *TwitterCloneAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *TwitterCloneAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
