// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"websvc/device/models"
)

// GetDeviceVersionHandlerFunc turns a function with the right signature into a get device version handler
type GetDeviceVersionHandlerFunc func(GetDeviceVersionParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDeviceVersionHandlerFunc) Handle(params GetDeviceVersionParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetDeviceVersionHandler interface for that can handle valid get device version params
type GetDeviceVersionHandler interface {
	Handle(GetDeviceVersionParams, *models.Principal) middleware.Responder
}

// NewGetDeviceVersion creates a new http.Handler for the get device version operation
func NewGetDeviceVersion(ctx *middleware.Context, handler GetDeviceVersionHandler) *GetDeviceVersion {
	return &GetDeviceVersion{Context: ctx, Handler: handler}
}

/* GetDeviceVersion swagger:route GET /device/version getDeviceVersion

Returns current api version.

*/
type GetDeviceVersion struct {
	Context *middleware.Context
	Handler GetDeviceVersionHandler
}

func (o *GetDeviceVersion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetDeviceVersionParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
