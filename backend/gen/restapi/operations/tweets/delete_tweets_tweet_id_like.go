// Code generated by go-swagger; DO NOT EDIT.

package tweets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteTweetsTweetIDLikeHandlerFunc turns a function with the right signature into a delete tweets tweet ID like handler
type DeleteTweetsTweetIDLikeHandlerFunc func(DeleteTweetsTweetIDLikeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteTweetsTweetIDLikeHandlerFunc) Handle(params DeleteTweetsTweetIDLikeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteTweetsTweetIDLikeHandler interface for that can handle valid delete tweets tweet ID like params
type DeleteTweetsTweetIDLikeHandler interface {
	Handle(DeleteTweetsTweetIDLikeParams, interface{}) middleware.Responder
}

// NewDeleteTweetsTweetIDLike creates a new http.Handler for the delete tweets tweet ID like operation
func NewDeleteTweetsTweetIDLike(ctx *middleware.Context, handler DeleteTweetsTweetIDLikeHandler) *DeleteTweetsTweetIDLike {
	return &DeleteTweetsTweetIDLike{Context: ctx, Handler: handler}
}

/*
	DeleteTweetsTweetIDLike swagger:route DELETE /tweets/{tweetId}/like tweets deleteTweetsTweetIdLike

いいねを取り消す
*/
type DeleteTweetsTweetIDLike struct {
	Context *middleware.Context
	Handler DeleteTweetsTweetIDLikeHandler
}

func (o *DeleteTweetsTweetIDLike) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteTweetsTweetIDLikeParams()
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
