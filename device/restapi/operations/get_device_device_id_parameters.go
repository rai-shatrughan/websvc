// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetDeviceDeviceIDParams creates a new GetDeviceDeviceIDParams object
//
// There are no default values defined in the spec.
func NewGetDeviceDeviceIDParams() GetDeviceDeviceIDParams {

	return GetDeviceDeviceIDParams{}
}

// GetDeviceDeviceIDParams contains all the bound params for the get device device ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetDeviceDeviceID
type GetDeviceDeviceIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	DeviceID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetDeviceDeviceIDParams() beforehand.
func (o *GetDeviceDeviceIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rDeviceID, rhkDeviceID, _ := route.Params.GetOK("deviceId")
	if err := o.bindDeviceID(rDeviceID, rhkDeviceID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDeviceID binds and validates parameter DeviceID from path.
func (o *GetDeviceDeviceIDParams) bindDeviceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("deviceId", "path", "strfmt.UUID", raw)
	}
	o.DeviceID = *(value.(*strfmt.UUID))

	if err := o.validateDeviceID(formats); err != nil {
		return err
	}

	return nil
}

// validateDeviceID carries on validations for parameter DeviceID
func (o *GetDeviceDeviceIDParams) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.FormatOf("deviceId", "path", "uuid", o.DeviceID.String(), formats); err != nil {
		return err
	}
	return nil
}
