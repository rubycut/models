// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Bind Bind
//
// HAProxy frontend bind configuration
// swagger:model bind
type Bind struct {

	// address
	// Pattern: ^[^\s]+$
	Address string `json:"address,omitempty"`

	// name
	// Required: true
	// Pattern: ^[^\s]+$
	Name string `json:"name"`

	// port
	// Maximum: 65535
	// Minimum: 0
	Port *int64 `json:"port,omitempty"`

	// process
	// Pattern: ^[^\s]+$
	Process string `json:"process,omitempty"`

	// ssl
	Ssl bool `json:"ssl,omitempty"`

	// ssl cafile
	// Pattern: ^[^\s]+$
	SslCafile string `json:"ssl_cafile,omitempty"`

	// ssl certificate
	// Pattern: ^[^\s]+$
	SslCertificate string `json:"ssl_certificate,omitempty"`

	// tcp user timeout
	TCPUserTimeout *int64 `json:"tcp_user_timeout,omitempty"`

	// transparent
	Transparent bool `json:"transparent,omitempty"`
}

// Validate validates this bind
func (m *Bind) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslCafile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslCertificate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Bind) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if err := validate.Pattern("address", "body", string(m.Address), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Bind) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Bind) validatePort(formats strfmt.Registry) error {

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

func (m *Bind) validateProcess(formats strfmt.Registry) error {

	if swag.IsZero(m.Process) { // not required
		return nil
	}

	if err := validate.Pattern("process", "body", string(m.Process), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Bind) validateSslCafile(formats strfmt.Registry) error {

	if swag.IsZero(m.SslCafile) { // not required
		return nil
	}

	if err := validate.Pattern("ssl_cafile", "body", string(m.SslCafile), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Bind) validateSslCertificate(formats strfmt.Registry) error {

	if swag.IsZero(m.SslCertificate) { // not required
		return nil
	}

	if err := validate.Pattern("ssl_certificate", "body", string(m.SslCertificate), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Bind) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Bind) UnmarshalBinary(b []byte) error {
	var res Bind
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
