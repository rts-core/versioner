// Code generated by go-swagger; DO NOT EDIT.

package applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"versioner/api/generated/models"
)

// GetVersionOKCode is the HTTP code returned for type GetVersionOK
const GetVersionOKCode int = 200

/*GetVersionOK Current Version Number

swagger:response getVersionOK
*/
type GetVersionOK struct {

	/*
	  In: Body
	*/
	Payload *models.ApplicationVersion `json:"body,omitempty"`
}

// NewGetVersionOK creates GetVersionOK with default headers values
func NewGetVersionOK() *GetVersionOK {

	return &GetVersionOK{}
}

// WithPayload adds the payload to the get version o k response
func (o *GetVersionOK) WithPayload(payload *models.ApplicationVersion) *GetVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get version o k response
func (o *GetVersionOK) SetPayload(payload *models.ApplicationVersion) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVersionNotFoundCode is the HTTP code returned for type GetVersionNotFound
const GetVersionNotFoundCode int = 404

/*GetVersionNotFound No version on record

swagger:response getVersionNotFound
*/
type GetVersionNotFound struct {
}

// NewGetVersionNotFound creates GetVersionNotFound with default headers values
func NewGetVersionNotFound() *GetVersionNotFound {

	return &GetVersionNotFound{}
}

// WriteResponse to the client
func (o *GetVersionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*GetVersionDefault Unexpected Error

swagger:response getVersionDefault
*/
type GetVersionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.DefaultError `json:"body,omitempty"`
}

// NewGetVersionDefault creates GetVersionDefault with default headers values
func NewGetVersionDefault(code int) *GetVersionDefault {
	if code <= 0 {
		code = 500
	}

	return &GetVersionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get version default response
func (o *GetVersionDefault) WithStatusCode(code int) *GetVersionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get version default response
func (o *GetVersionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get version default response
func (o *GetVersionDefault) WithPayload(payload *models.DefaultError) *GetVersionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get version default response
func (o *GetVersionDefault) SetPayload(payload *models.DefaultError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVersionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
