// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutTimeseriesDeviceIDOKCode is the HTTP code returned for type PutTimeseriesDeviceIDOK
const PutTimeseriesDeviceIDOKCode int = 200

/*PutTimeseriesDeviceIDOK Timeseries was pushed.

swagger:response putTimeseriesDeviceIdOK
*/
type PutTimeseriesDeviceIDOK struct {
}

// NewPutTimeseriesDeviceIDOK creates PutTimeseriesDeviceIDOK with default headers values
func NewPutTimeseriesDeviceIDOK() *PutTimeseriesDeviceIDOK {

	return &PutTimeseriesDeviceIDOK{}
}

// WriteResponse to the client
func (o *PutTimeseriesDeviceIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutTimeseriesDeviceIDBadRequestCode is the HTTP code returned for type PutTimeseriesDeviceIDBadRequest
const PutTimeseriesDeviceIDBadRequestCode int = 400

/*PutTimeseriesDeviceIDBadRequest The specified device ID is invalid (e.g. not a number).

swagger:response putTimeseriesDeviceIdBadRequest
*/
type PutTimeseriesDeviceIDBadRequest struct {
}

// NewPutTimeseriesDeviceIDBadRequest creates PutTimeseriesDeviceIDBadRequest with default headers values
func NewPutTimeseriesDeviceIDBadRequest() *PutTimeseriesDeviceIDBadRequest {

	return &PutTimeseriesDeviceIDBadRequest{}
}

// WriteResponse to the client
func (o *PutTimeseriesDeviceIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PutTimeseriesDeviceIDNotFoundCode is the HTTP code returned for type PutTimeseriesDeviceIDNotFound
const PutTimeseriesDeviceIDNotFoundCode int = 404

/*PutTimeseriesDeviceIDNotFound A device with the specified ID was not found.

swagger:response putTimeseriesDeviceIdNotFound
*/
type PutTimeseriesDeviceIDNotFound struct {
}

// NewPutTimeseriesDeviceIDNotFound creates PutTimeseriesDeviceIDNotFound with default headers values
func NewPutTimeseriesDeviceIDNotFound() *PutTimeseriesDeviceIDNotFound {

	return &PutTimeseriesDeviceIDNotFound{}
}

// WriteResponse to the client
func (o *PutTimeseriesDeviceIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*PutTimeseriesDeviceIDDefault Unexpected error

swagger:response putTimeseriesDeviceIdDefault
*/
type PutTimeseriesDeviceIDDefault struct {
	_statusCode int
}

// NewPutTimeseriesDeviceIDDefault creates PutTimeseriesDeviceIDDefault with default headers values
func NewPutTimeseriesDeviceIDDefault(code int) *PutTimeseriesDeviceIDDefault {
	if code <= 0 {
		code = 500
	}

	return &PutTimeseriesDeviceIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put timeseries device ID default response
func (o *PutTimeseriesDeviceIDDefault) WithStatusCode(code int) *PutTimeseriesDeviceIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put timeseries device ID default response
func (o *PutTimeseriesDeviceIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *PutTimeseriesDeviceIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
