// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutUserUsernameHandlerFunc turns a function with the right signature into a put user username handler
type PutUserUsernameHandlerFunc func(PutUserUsernameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutUserUsernameHandlerFunc) Handle(params PutUserUsernameParams) middleware.Responder {
	return fn(params)
}

// PutUserUsernameHandler interface for that can handle valid put user username params
type PutUserUsernameHandler interface {
	Handle(PutUserUsernameParams) middleware.Responder
}

// NewPutUserUsername creates a new http.Handler for the put user username operation
func NewPutUserUsername(ctx *middleware.Context, handler PutUserUsernameHandler) *PutUserUsername {
	return &PutUserUsername{Context: ctx, Handler: handler}
}

/*PutUserUsername swagger:route PUT /user/{username} user putUserUsername

Updated user

Change information about user

*/
type PutUserUsername struct {
	Context *middleware.Context
	Handler PutUserUsernameHandler
}

func (o *PutUserUsername) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutUserUsernameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
