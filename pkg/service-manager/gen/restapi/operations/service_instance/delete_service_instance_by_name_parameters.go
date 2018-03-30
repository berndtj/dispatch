///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package service_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteServiceInstanceByNameParams creates a new DeleteServiceInstanceByNameParams object
// no default values defined in spec.
func NewDeleteServiceInstanceByNameParams() DeleteServiceInstanceByNameParams {

	return DeleteServiceInstanceByNameParams{}
}

// DeleteServiceInstanceByNameParams contains all the bound params for the delete service instance by name operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteServiceInstanceByName
type DeleteServiceInstanceByNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Name of service instance to return
	  Required: true
	  Pattern: ^[\w\d\-]+$
	  In: path
	*/
	ServiceInstanceName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteServiceInstanceByNameParams() beforehand.
func (o *DeleteServiceInstanceByNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

func (o *DeleteServiceInstanceByNameParams) bindServiceInstanceName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ServiceInstanceName = raw

	if err := o.validateServiceInstanceName(formats); err != nil {
		return err
	}

	return nil
}

func (o *DeleteServiceInstanceByNameParams) validateServiceInstanceName(formats strfmt.Registry) error {

	if err := validate.Pattern("serviceInstanceName", "path", o.ServiceInstanceName, `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}
