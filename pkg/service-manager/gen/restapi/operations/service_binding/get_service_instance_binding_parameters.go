// Code generated by go-swagger; DO NOT EDIT.

package service_binding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetServiceInstanceBindingParams creates a new GetServiceInstanceBindingParams object
// with the default values initialized.
func NewGetServiceInstanceBindingParams() GetServiceInstanceBindingParams {
	var ()
	return GetServiceInstanceBindingParams{}
}

// GetServiceInstanceBindingParams contains all the bound params for the get service instance binding operation
// typically these are obtained from a http.Request
//
// swagger:parameters getServiceInstanceBinding
type GetServiceInstanceBindingParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Name of service instance which is bound
	  Required: true
	  Pattern: ^[\w\d\-]+$
	  In: path
	*/
	ServiceInstanceName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetServiceInstanceBindingParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rServiceInstanceName, rhkServiceInstanceName, _ := route.Params.GetOK("serviceInstanceName")
	if err := o.bindServiceInstanceName(rServiceInstanceName, rhkServiceInstanceName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceInstanceBindingParams) bindServiceInstanceName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ServiceInstanceName = raw

	if err := o.validateServiceInstanceName(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetServiceInstanceBindingParams) validateServiceInstanceName(formats strfmt.Registry) error {

	if err := validate.Pattern("serviceInstanceName", "path", o.ServiceInstanceName, `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}
