// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"websvc/timeseries/models"
)

// GetTimeseriesDeviceIDHandlerFunc turns a function with the right signature into a get timeseries device ID handler
type GetTimeseriesDeviceIDHandlerFunc func(GetTimeseriesDeviceIDParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTimeseriesDeviceIDHandlerFunc) Handle(params GetTimeseriesDeviceIDParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetTimeseriesDeviceIDHandler interface for that can handle valid get timeseries device ID params
type GetTimeseriesDeviceIDHandler interface {
	Handle(GetTimeseriesDeviceIDParams, *models.Principal) middleware.Responder
}

// NewGetTimeseriesDeviceID creates a new http.Handler for the get timeseries device ID operation
func NewGetTimeseriesDeviceID(ctx *middleware.Context, handler GetTimeseriesDeviceIDHandler) *GetTimeseriesDeviceID {
	return &GetTimeseriesDeviceID{Context: ctx, Handler: handler}
}

/* GetTimeseriesDeviceID swagger:route GET /timeseries/{deviceId} getTimeseriesDeviceId

get timeseries data for a device by ID.

*/
type GetTimeseriesDeviceID struct {
	Context *middleware.Context
	Handler GetTimeseriesDeviceIDHandler
}

func (o *GetTimeseriesDeviceID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTimeseriesDeviceIDParams()
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
