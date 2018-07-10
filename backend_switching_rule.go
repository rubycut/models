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

// BackendSwitchingRule Backend Switching Rule
//
// HAProxy backend switching rule configuration (corresponds to use_backend directive)
// swagger:model backend_switching_rule
type BackendSwitchingRule struct {

	// cond
	// Enum: [if unless]
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// target farm
	// Required: true
	TargetFarm *string `json:"target_farm"`
}

// Validate validates this backend switching rule
func (m *BackendSwitchingRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetFarm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var backendSwitchingRuleTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendSwitchingRuleTypeCondPropEnum = append(backendSwitchingRuleTypeCondPropEnum, v)
	}
}

const (

	// BackendSwitchingRuleCondIf captures enum value "if"
	BackendSwitchingRuleCondIf string = "if"

	// BackendSwitchingRuleCondUnless captures enum value "unless"
	BackendSwitchingRuleCondUnless string = "unless"
)

// prop value enum
func (m *BackendSwitchingRule) validateCondEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendSwitchingRuleTypeCondPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BackendSwitchingRule) validateCond(formats strfmt.Registry) error {

	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

func (m *BackendSwitchingRule) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *BackendSwitchingRule) validateTargetFarm(formats strfmt.Registry) error {

	if err := validate.Required("target_farm", "body", m.TargetFarm); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackendSwitchingRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackendSwitchingRule) UnmarshalBinary(b []byte) error {
	var res BackendSwitchingRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
