// Code generated by go-swagger; DO NOT EDIT.

package tweets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostTweetsHandlerFunc turns a function with the right signature into a post tweets handler
type PostTweetsHandlerFunc func(PostTweetsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostTweetsHandlerFunc) Handle(params PostTweetsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostTweetsHandler interface for that can handle valid post tweets params
type PostTweetsHandler interface {
	Handle(PostTweetsParams, interface{}) middleware.Responder
}

// NewPostTweets creates a new http.Handler for the post tweets operation
func NewPostTweets(ctx *middleware.Context, handler PostTweetsHandler) *PostTweets {
	return &PostTweets{Context: ctx, Handler: handler}
}

/*
	PostTweets swagger:route POST /tweets tweets postTweets

ツイート投稿
*/
type PostTweets struct {
	Context *middleware.Context
	Handler PostTweetsHandler
}

func (o *PostTweets) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostTweetsParams()
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

// PostTweetsBody post tweets body
//
// swagger:model PostTweetsBody
type PostTweetsBody struct {

	// content
	// Required: true
	Content *string `json:"content"`
}

// Validate validates this post tweets body
func (o *PostTweetsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTweetsBody) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"content", "body", o.Content); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post tweets body based on context it is used
func (o *PostTweetsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostTweetsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTweetsBody) UnmarshalBinary(b []byte) error {
	var res PostTweetsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
