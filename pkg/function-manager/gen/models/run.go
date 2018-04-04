///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Run run
// swagger:model Run
type Run struct {

	// blocking
	Blocking bool `json:"blocking,omitempty"`

	// event
	Event *CloudEvent `json:"event,omitempty"`

	// executed time
	// Read Only: true
	ExecutedTime int64 `json:"executedTime,omitempty"`

	// finished time
	// Read Only: true
	FinishedTime int64 `json:"finishedTime,omitempty"`

	// function Id
	// Read Only: true
	FunctionID string `json:"functionId,omitempty"`

	// function name
	// Read Only: true
	FunctionName string `json:"functionName,omitempty"`

	// input
	Input interface{} `json:"input,omitempty"`

	// logs
	Logs []string `json:"logs"`

	// name
	// Read Only: true
	Name strfmt.UUID `json:"name,omitempty"`

	// output
	// Read Only: true
	Output interface{} `json:"output,omitempty"`

	// reason
	Reason []string `json:"reason"`

	// secrets
	Secrets []string `json:"secrets"`

	// services
	Services []string `json:"services"`

	// status
	Status Status `json:"status,omitempty"`

	// tags
	Tags []*Tag `json:"tags"`
}

// Validate validates this run
func (m *Run) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLogs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSecrets(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Run) validateEvent(formats strfmt.Registry) error {

	if swag.IsZero(m.Event) { // not required
		return nil
	}

	if m.Event != nil {

		if err := m.Event.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event")
			}
			return err
		}

	}

	return nil
}

func (m *Run) validateLogs(formats strfmt.Registry) error {

	if swag.IsZero(m.Logs) { // not required
		return nil
	}

	return nil
}

func (m *Run) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.FormatOf("name", "body", "uuid", m.Name.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Run) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	return nil
}

func (m *Run) validateSecrets(formats strfmt.Registry) error {

	if swag.IsZero(m.Secrets) { // not required
		return nil
	}

	return nil
}

func (m *Run) validateServices(formats strfmt.Registry) error {

	if swag.IsZero(m.Services) { // not required
		return nil
	}

	return nil
}

func (m *Run) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *Run) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {

			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Run) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Run) UnmarshalBinary(b []byte) error {
	var res Run
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
