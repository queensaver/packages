/*
 * Queensaver API
 *
 * Queensaver API to send in sensor data and retrieve it. It's also used for user management.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
	errorHandler ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"BboxesGet",
			strings.ToUpper("Get"),
			"/v1/bboxes",
			c.BboxesGet,
		},
		{
			"LoginPost",
			strings.ToUpper("Post"),
			"/v1/login",
			c.LoginPost,
		},
		{
			"ScaleGet",
			strings.ToUpper("Get"),
			"/v1/scale",
			c.ScaleGet,
		},
		{
			"UserPost",
			strings.ToUpper("Post"),
			"/v1/user",
			c.UserPost,
		},
	}
}

// BboxesGet - Get QBox metadata
func (c *DefaultApiController) BboxesGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.BboxesGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// LoginPost - Authenticate a user against the system.
func (c *DefaultApiController) LoginPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
				usernameParam := r.FormValue("username")
				passwordParam := r.FormValue("password")
	result, err := c.service.LoginPost(r.Context(), usernameParam, passwordParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ScaleGet - Get Scale values
func (c *DefaultApiController) ScaleGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	qTokenParam := r.Header.Get("Q-Token")
	bhiveIdParam := query.Get("bhive_id")
	epochParam, err := parseInt64Parameter(query.Get("epoch"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	secondsInThePastParam, err := parseInt64Parameter(query.Get("seconds_in_the_past"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.ScaleGet(r.Context(), qTokenParam, bhiveIdParam, epochParam, secondsInThePastParam, tokenParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UserPost - Create a user
func (c *DefaultApiController) UserPost(w http.ResponseWriter, r *http.Request) {
	userParam := User{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserRequired(userParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UserPost(r.Context(), userParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
