// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SiteBackendsItemsServersItems site backends items servers items
// swagger:model siteBackendsItemsServersItems
type SiteBackendsItemsServersItems struct {

	// address
	Address string `json:"address,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// port
	// Maximum: 65535
	// Minimum: 0
	Port *int64 `json:"port,omitempty"`

	// ssl
	// Enum: [enabled]
	Ssl string `json:"ssl,omitempty"`

	// weight
	Weight *int64 `json:"weight,omitempty"`
}

// Validate validates this site backends items servers items
func (m *SiteBackendsItemsServersItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteBackendsItemsServersItems) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

var siteBackendsItemsServersItemsTypeSslPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteBackendsItemsServersItemsTypeSslPropEnum = append(siteBackendsItemsServersItemsTypeSslPropEnum, v)
	}
}

const (

	// SiteBackendsItemsServersItemsSslEnabled captures enum value "enabled"
	SiteBackendsItemsServersItemsSslEnabled string = "enabled"
)

// prop value enum
func (m *SiteBackendsItemsServersItems) validateSslEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, siteBackendsItemsServersItemsTypeSslPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SiteBackendsItemsServersItems) validateSsl(formats strfmt.Registry) error {

	if swag.IsZero(m.Ssl) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslEnum("ssl", "body", m.Ssl); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteBackendsItemsServersItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteBackendsItemsServersItems) UnmarshalBinary(b []byte) error {
	var res SiteBackendsItemsServersItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}