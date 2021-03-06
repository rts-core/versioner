// Code generated by go-swagger; DO NOT EDIT.

package applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetVersionParams creates a new GetVersionParams object
//
// There are no default values defined in the spec.
func NewGetVersionParams() GetVersionParams {

	return GetVersionParams{}
}

// GetVersionParams contains all the bound params for the get version operation
// typically these are obtained from a http.Request
//
// swagger:parameters getVersion
type GetVersionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ID of the application.
	  Required: true
	  In: path
	*/
	ApplicationID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetVersionParams() beforehand.
func (o *GetVersionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rApplicationID, rhkApplicationID, _ := route.Params.GetOK("applicationId")
	if err := o.bindApplicationID(rApplicationID, rhkApplicationID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindApplicationID binds and validates parameter ApplicationID from path.
func (o *GetVersionParams) bindApplicationID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ApplicationID = raw

	return nil
}
