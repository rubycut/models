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

// TCPRule TCP Rule
//
// HAProxy TCP rule configuration (corresponds to tcp-request and tcp-response directive)
// swagger:model tcp_rule
type TCPRule struct {

	// cond
	// Enum: [if unless]
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// type
	// Required: true
	// Enum: [accept reject]
	Type *string `json:"type"`
}

// Validate validates this tcp rule
func (m *TCPRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tcpRuleTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpRuleTypeCondPropEnum = append(tcpRuleTypeCondPropEnum, v)
	}
}

const (

	// TCPRuleCondIf captures enum value "if"
	TCPRuleCondIf string = "if"

	// TCPRuleCondUnless captures enum value "unless"
	TCPRuleCondUnless string = "unless"
)

// prop value enum
func (m *TCPRule) validateCondEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tcpRuleTypeCondPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TCPRule) validateCond(formats strfmt.Registry) error {

	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

func (m *TCPRule) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var tcpRuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["accept","reject"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpRuleTypeTypePropEnum = append(tcpRuleTypeTypePropEnum, v)
	}
}

const (

	// TCPRuleTypeAccept captures enum value "accept"
	TCPRuleTypeAccept string = "accept"

	// TCPRuleTypeReject captures enum value "reject"
	TCPRuleTypeReject string = "reject"
)

// prop value enum
func (m *TCPRule) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tcpRuleTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TCPRule) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TCPRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TCPRule) UnmarshalBinary(b []byte) error {
	var res TCPRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
