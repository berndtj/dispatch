///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package v1

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NO TESTS

// EventDriver driver
// swagger:model EventDriver
type EventDriver struct {

	// config
	Config []*Config `json:"config"`

	// created time
	// Read Only: true
	CreatedTime int64 `json:"created-time,omitempty"`

	// id
	// Read Only: true
	ID strfmt.UUID `json:"id,omitempty"`

	// kind
	// Read Only: true
	// Pattern: ^[\w\d\-]+$
	Kind string `json:"kind,omitempty"`

	// modified time
	// Read Only: true
	ModifiedTime int64 `json:"modified-time,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// secrets
	Secrets []string `json:"secrets"`

	// status
	// Read Only: true
	Status Status `json:"status,omitempty"`

	// tags
	Tags []*Tag `json:"tags"`

	// type
	// Required: true
	// Max Length: 32
	Type *string `json:"type"`
}

// Validate validates this driver
func (m *EventDriver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventDriver) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	for i := 0; i < len(m.Config); i++ {

		if swag.IsZero(m.Config[i]) { // not required
			continue
		}

		if m.Config[i] != nil {

			if err := m.Config[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("config" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *EventDriver) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}
	return nil
}

func (m *EventDriver) validateKind(formats strfmt.Registry) error {
	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if err := validate.Pattern("kind", "body", string(m.Kind), `^[\w\d\-]+$`); err != nil {
		return err
	}
	return nil
}

func (m *EventDriver) validateName(formats strfmt.Registry) error {
	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}
	return nil
}

func (m *EventDriver) validateStatus(formats strfmt.Registry) error {
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

func (m *EventDriver) validateTags(formats strfmt.Registry) error {
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

func (m *EventDriver) validateType(formats strfmt.Registry) error {
	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.MaxLength("type", "body", string(*m.Type), 32); err != nil {
		return err
	}
	return nil
}

// MarshalBinary interface implementation
func (m *EventDriver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventDriver) UnmarshalBinary(b []byte) error {
	var res EventDriver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}