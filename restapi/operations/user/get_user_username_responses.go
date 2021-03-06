// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/crioto/swagger-playground/models"
)

// GetUserUsernameOKCode is the HTTP code returned for type GetUserUsernameOK
const GetUserUsernameOKCode int = 200

/*GetUserUsernameOK Successful

swagger:response getUserUsernameOK
*/
type GetUserUsernameOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserUsernameOK creates GetUserUsernameOK with default headers values
func NewGetUserUsernameOK() *GetUserUsernameOK {
	return &GetUserUsernameOK{}
}

// WithPayload adds the payload to the get user username o k response
func (o *GetUserUsernameOK) WithPayload(payload *models.User) *GetUserUsernameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user username o k response
func (o *GetUserUsernameOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserUsernameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserUsernameBadRequestCode is the HTTP code returned for type GetUserUsernameBadRequest
const GetUserUsernameBadRequestCode int = 400

/*GetUserUsernameBadRequest Invalid username

swagger:response getUserUsernameBadRequest
*/
type GetUserUsernameBadRequest struct {
}

// NewGetUserUsernameBadRequest creates GetUserUsernameBadRequest with default headers values
func NewGetUserUsernameBadRequest() *GetUserUsernameBadRequest {
	return &GetUserUsernameBadRequest{}
}

// WriteResponse to the client
func (o *GetUserUsernameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetUserUsernameNotFoundCode is the HTTP code returned for type GetUserUsernameNotFound
const GetUserUsernameNotFoundCode int = 404

/*GetUserUsernameNotFound User not found

swagger:response getUserUsernameNotFound
*/
type GetUserUsernameNotFound struct {
}

// NewGetUserUsernameNotFound creates GetUserUsernameNotFound with default headers values
func NewGetUserUsernameNotFound() *GetUserUsernameNotFound {
	return &GetUserUsernameNotFound{}
}

// WriteResponse to the client
func (o *GetUserUsernameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
