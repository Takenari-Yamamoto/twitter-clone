// Code generated by go-swagger; DO NOT EDIT.

package tweets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetTweetsTweetIDParams creates a new GetTweetsTweetIDParams object
//
// There are no default values defined in the spec.
func NewGetTweetsTweetIDParams() GetTweetsTweetIDParams {

	return GetTweetsTweetIDParams{}
}

// GetTweetsTweetIDParams contains all the bound params for the get tweets tweet ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetTweetsTweetID
type GetTweetsTweetIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	TweetID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTweetsTweetIDParams() beforehand.
func (o *GetTweetsTweetIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rTweetID, rhkTweetID, _ := route.Params.GetOK("tweetId")
	if err := o.bindTweetID(rTweetID, rhkTweetID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTweetID binds and validates parameter TweetID from path.
func (o *GetTweetsTweetIDParams) bindTweetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("tweetId", "path", "int64", raw)
	}
	o.TweetID = value

	return nil
}
