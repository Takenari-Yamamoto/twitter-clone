// Code generated by go-swagger; DO NOT EDIT.

package tweets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostTweetsTweetIDLikeHandlerFunc turns a function with the right signature into a post tweets tweet ID like handler
type PostTweetsTweetIDLikeHandlerFunc func(PostTweetsTweetIDLikeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostTweetsTweetIDLikeHandlerFunc) Handle(params PostTweetsTweetIDLikeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostTweetsTweetIDLikeHandler interface for that can handle valid post tweets tweet ID like params
type PostTweetsTweetIDLikeHandler interface {
	Handle(PostTweetsTweetIDLikeParams, interface{}) middleware.Responder
}

// NewPostTweetsTweetIDLike creates a new http.Handler for the post tweets tweet ID like operation
func NewPostTweetsTweetIDLike(ctx *middleware.Context, handler PostTweetsTweetIDLikeHandler) *PostTweetsTweetIDLike {
	return &PostTweetsTweetIDLike{Context: ctx, Handler: handler}
}

/*
	PostTweetsTweetIDLike swagger:route POST /tweets/{tweetId}/like tweets postTweetsTweetIdLike

ツイートにいいねをする
*/
type PostTweetsTweetIDLike struct {
	Context *middleware.Context
	Handler PostTweetsTweetIDLikeHandler
}

func (o *PostTweetsTweetIDLike) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostTweetsTweetIDLikeParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
