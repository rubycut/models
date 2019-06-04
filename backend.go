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
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Backend Backend
//
// HAProxy backend configuration
// swagger:model backend
type Backend struct {

	// adv check
	// Enum: [ssl-hello-chk smtpchk ldap-check mysql-check pgsql-check tcp-check redis-check]
	AdvCheck string `json:"adv_check,omitempty"`

	// balance
	Balance *Balance `json:"balance,omitempty"`

	// check timeout
	CheckTimeout *int64 `json:"check_timeout,omitempty"`

	// connect timeout
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`

	// contstats
	// Enum: [enabled disabled]
	Contstats string `json:"contstats,omitempty"`

	// cookie
	// Pattern: ^[^\s]+$
	Cookie string `json:"cookie,omitempty"`

	// default server
	DefaultServer *DefaultServer `json:"default_server,omitempty"`

	// forwardfor
	Forwardfor *Forwardfor `json:"forwardfor,omitempty"`

	// http use htx
	// Enum: [enabled disabled]
	HTTPUseHtx string `json:"http-use-htx,omitempty"`

	// http connection mode
	// Enum: [http-tunnel httpclose http-server-close http-keep-alive]
	HTTPConnectionMode string `json:"http_connection_mode,omitempty"`

	// httpchk
	Httpchk *BackendHttpchk `json:"httpchk,omitempty"`

	// log tag
	// Pattern: ^[^\s]+$
	LogTag string `json:"log_tag,omitempty"`

	// mode
	// Enum: [http tcp health]
	Mode string `json:"mode,omitempty"`

	// name
	// Required: true
	// Pattern: ^[A-Za-z0-9-_.:]+$
	Name string `json:"name"`

	// queue timeout
	QueueTimeout *int64 `json:"queue_timeout,omitempty"`

	// redispatch
	Redispatch *BackendRedispatch `json:"redispatch,omitempty"`

	// retries
	Retries *int64 `json:"retries,omitempty"`

	// server timeout
	ServerTimeout *int64 `json:"server_timeout,omitempty"`

	// stick table
	StickTable *BackendStickTable `json:"stick_table,omitempty"`
}

// Validate validates this backend
func (m *Backend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContstats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCookie(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForwardfor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPUseHtx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPConnectionMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHttpchk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedispatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStickTable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var backendTypeAdvCheckPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ssl-hello-chk","smtpchk","ldap-check","mysql-check","pgsql-check","tcp-check","redis-check"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeAdvCheckPropEnum = append(backendTypeAdvCheckPropEnum, v)
	}
}

const (

	// BackendAdvCheckSslHelloChk captures enum value "ssl-hello-chk"
	BackendAdvCheckSslHelloChk string = "ssl-hello-chk"

	// BackendAdvCheckSmtpchk captures enum value "smtpchk"
	BackendAdvCheckSmtpchk string = "smtpchk"

	// BackendAdvCheckLdapCheck captures enum value "ldap-check"
	BackendAdvCheckLdapCheck string = "ldap-check"

	// BackendAdvCheckMysqlCheck captures enum value "mysql-check"
	BackendAdvCheckMysqlCheck string = "mysql-check"

	// BackendAdvCheckPgsqlCheck captures enum value "pgsql-check"
	BackendAdvCheckPgsqlCheck string = "pgsql-check"

	// BackendAdvCheckTCPCheck captures enum value "tcp-check"
	BackendAdvCheckTCPCheck string = "tcp-check"

	// BackendAdvCheckRedisCheck captures enum value "redis-check"
	BackendAdvCheckRedisCheck string = "redis-check"
)

// prop value enum
func (m *Backend) validateAdvCheckEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeAdvCheckPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateAdvCheck(formats strfmt.Registry) error {

	if swag.IsZero(m.AdvCheck) { // not required
		return nil
	}

	// value enum
	if err := m.validateAdvCheckEnum("adv_check", "body", m.AdvCheck); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateBalance(formats strfmt.Registry) error {

	if swag.IsZero(m.Balance) { // not required
		return nil
	}

	if m.Balance != nil {
		if err := m.Balance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance")
			}
			return err
		}
	}

	return nil
}

var backendTypeContstatsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeContstatsPropEnum = append(backendTypeContstatsPropEnum, v)
	}
}

const (

	// BackendContstatsEnabled captures enum value "enabled"
	BackendContstatsEnabled string = "enabled"

	// BackendContstatsDisabled captures enum value "disabled"
	BackendContstatsDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateContstatsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeContstatsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateContstats(formats strfmt.Registry) error {

	if swag.IsZero(m.Contstats) { // not required
		return nil
	}

	// value enum
	if err := m.validateContstatsEnum("contstats", "body", m.Contstats); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateCookie(formats strfmt.Registry) error {

	if swag.IsZero(m.Cookie) { // not required
		return nil
	}

	if err := validate.Pattern("cookie", "body", string(m.Cookie), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateDefaultServer(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultServer) { // not required
		return nil
	}

	if m.DefaultServer != nil {
		if err := m.DefaultServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_server")
			}
			return err
		}
	}

	return nil
}

func (m *Backend) validateForwardfor(formats strfmt.Registry) error {

	if swag.IsZero(m.Forwardfor) { // not required
		return nil
	}

	if m.Forwardfor != nil {
		if err := m.Forwardfor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forwardfor")
			}
			return err
		}
	}

	return nil
}

var backendTypeHTTPUseHtxPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHTTPUseHtxPropEnum = append(backendTypeHTTPUseHtxPropEnum, v)
	}
}

const (

	// BackendHTTPUseHtxEnabled captures enum value "enabled"
	BackendHTTPUseHtxEnabled string = "enabled"

	// BackendHTTPUseHtxDisabled captures enum value "disabled"
	BackendHTTPUseHtxDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateHTTPUseHtxEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHTTPUseHtxPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHTTPUseHtx(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPUseHtx) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPUseHtxEnum("http-use-htx", "body", m.HTTPUseHtx); err != nil {
		return err
	}

	return nil
}

var backendTypeHTTPConnectionModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http-tunnel","httpclose","http-server-close","http-keep-alive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHTTPConnectionModePropEnum = append(backendTypeHTTPConnectionModePropEnum, v)
	}
}

const (

	// BackendHTTPConnectionModeHTTPTunnel captures enum value "http-tunnel"
	BackendHTTPConnectionModeHTTPTunnel string = "http-tunnel"

	// BackendHTTPConnectionModeHttpclose captures enum value "httpclose"
	BackendHTTPConnectionModeHttpclose string = "httpclose"

	// BackendHTTPConnectionModeHTTPServerClose captures enum value "http-server-close"
	BackendHTTPConnectionModeHTTPServerClose string = "http-server-close"

	// BackendHTTPConnectionModeHTTPKeepAlive captures enum value "http-keep-alive"
	BackendHTTPConnectionModeHTTPKeepAlive string = "http-keep-alive"
)

// prop value enum
func (m *Backend) validateHTTPConnectionModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHTTPConnectionModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHTTPConnectionMode(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPConnectionMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPConnectionModeEnum("http_connection_mode", "body", m.HTTPConnectionMode); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateHttpchk(formats strfmt.Registry) error {

	if swag.IsZero(m.Httpchk) { // not required
		return nil
	}

	if m.Httpchk != nil {
		if err := m.Httpchk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpchk")
			}
			return err
		}
	}

	return nil
}

func (m *Backend) validateLogTag(formats strfmt.Registry) error {

	if swag.IsZero(m.LogTag) { // not required
		return nil
	}

	if err := validate.Pattern("log_tag", "body", string(m.LogTag), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var backendTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","tcp","health"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeModePropEnum = append(backendTypeModePropEnum, v)
	}
}

const (

	// BackendModeHTTP captures enum value "http"
	BackendModeHTTP string = "http"

	// BackendModeTCP captures enum value "tcp"
	BackendModeTCP string = "tcp"

	// BackendModeHealth captures enum value "health"
	BackendModeHealth string = "health"
)

// prop value enum
func (m *Backend) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateRedispatch(formats strfmt.Registry) error {

	if swag.IsZero(m.Redispatch) { // not required
		return nil
	}

	if m.Redispatch != nil {
		if err := m.Redispatch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redispatch")
			}
			return err
		}
	}

	return nil
}

func (m *Backend) validateStickTable(formats strfmt.Registry) error {

	if swag.IsZero(m.StickTable) { // not required
		return nil
	}

	if m.StickTable != nil {
		if err := m.StickTable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stick_table")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Backend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Backend) UnmarshalBinary(b []byte) error {
	var res Backend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BackendHttpchk backend httpchk
// swagger:model BackendHttpchk
type BackendHttpchk struct {

	// method
	// Enum: [HEAD PUT POST GET TRACE PATCH]
	Method string `json:"method,omitempty"`

	// uri
	// Pattern: ^[^\s]+$
	URI string `json:"uri,omitempty"`

	// version
	// Pattern: ^[^\s]+$
	Version string `json:"version,omitempty"`
}

// Validate validates this backend httpchk
func (m *BackendHttpchk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var backendHttpchkTypeMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HEAD","PUT","POST","GET","TRACE","PATCH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendHttpchkTypeMethodPropEnum = append(backendHttpchkTypeMethodPropEnum, v)
	}
}

const (

	// BackendHttpchkMethodHEAD captures enum value "HEAD"
	BackendHttpchkMethodHEAD string = "HEAD"

	// BackendHttpchkMethodPUT captures enum value "PUT"
	BackendHttpchkMethodPUT string = "PUT"

	// BackendHttpchkMethodPOST captures enum value "POST"
	BackendHttpchkMethodPOST string = "POST"

	// BackendHttpchkMethodGET captures enum value "GET"
	BackendHttpchkMethodGET string = "GET"

	// BackendHttpchkMethodTRACE captures enum value "TRACE"
	BackendHttpchkMethodTRACE string = "TRACE"

	// BackendHttpchkMethodPATCH captures enum value "PATCH"
	BackendHttpchkMethodPATCH string = "PATCH"
)

// prop value enum
func (m *BackendHttpchk) validateMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendHttpchkTypeMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BackendHttpchk) validateMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.Method) { // not required
		return nil
	}

	// value enum
	if err := m.validateMethodEnum("httpchk"+"."+"method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

func (m *BackendHttpchk) validateURI(formats strfmt.Registry) error {

	if swag.IsZero(m.URI) { // not required
		return nil
	}

	if err := validate.Pattern("httpchk"+"."+"uri", "body", string(m.URI), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *BackendHttpchk) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.Pattern("httpchk"+"."+"version", "body", string(m.Version), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackendHttpchk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackendHttpchk) UnmarshalBinary(b []byte) error {
	var res BackendHttpchk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BackendRedispatch backend redispatch
// swagger:model BackendRedispatch
type BackendRedispatch struct {

	// enabled
	// Required: true
	// Enum: [enabled disabled]
	Enabled *string `json:"enabled"`

	// interval
	Interval int64 `json:"interval,omitempty"`
}

// Validate validates this backend redispatch
func (m *BackendRedispatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var backendRedispatchTypeEnabledPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendRedispatchTypeEnabledPropEnum = append(backendRedispatchTypeEnabledPropEnum, v)
	}
}

const (

	// BackendRedispatchEnabledEnabled captures enum value "enabled"
	BackendRedispatchEnabledEnabled string = "enabled"

	// BackendRedispatchEnabledDisabled captures enum value "disabled"
	BackendRedispatchEnabledDisabled string = "disabled"
)

// prop value enum
func (m *BackendRedispatch) validateEnabledEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendRedispatchTypeEnabledPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BackendRedispatch) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("redispatch"+"."+"enabled", "body", m.Enabled); err != nil {
		return err
	}

	// value enum
	if err := m.validateEnabledEnum("redispatch"+"."+"enabled", "body", *m.Enabled); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackendRedispatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackendRedispatch) UnmarshalBinary(b []byte) error {
	var res BackendRedispatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BackendStickTable backend stick table
// swagger:model BackendStickTable
type BackendStickTable struct {

	// expire
	Expire *int64 `json:"expire,omitempty"`

	// keylen
	Keylen *int64 `json:"keylen,omitempty"`

	// nopurge
	Nopurge bool `json:"nopurge,omitempty"`

	// peers
	// Pattern: ^[^\s]+$
	Peers string `json:"peers,omitempty"`

	// size
	Size *int64 `json:"size,omitempty"`

	// store
	// Pattern: ^[^\s]+$
	Store string `json:"store,omitempty"`

	// type
	// Enum: [ip ipv6 integer string binary]
	Type string `json:"type,omitempty"`
}

// Validate validates this backend stick table
func (m *BackendStickTable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStore(formats); err != nil {
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

func (m *BackendStickTable) validatePeers(formats strfmt.Registry) error {

	if swag.IsZero(m.Peers) { // not required
		return nil
	}

	if err := validate.Pattern("stick_table"+"."+"peers", "body", string(m.Peers), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *BackendStickTable) validateStore(formats strfmt.Registry) error {

	if swag.IsZero(m.Store) { // not required
		return nil
	}

	if err := validate.Pattern("stick_table"+"."+"store", "body", string(m.Store), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var backendStickTableTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ip","ipv6","integer","string","binary"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendStickTableTypeTypePropEnum = append(backendStickTableTypeTypePropEnum, v)
	}
}

const (

	// BackendStickTableTypeIP captures enum value "ip"
	BackendStickTableTypeIP string = "ip"

	// BackendStickTableTypeIPV6 captures enum value "ipv6"
	BackendStickTableTypeIPV6 string = "ipv6"

	// BackendStickTableTypeInteger captures enum value "integer"
	BackendStickTableTypeInteger string = "integer"

	// BackendStickTableTypeString captures enum value "string"
	BackendStickTableTypeString string = "string"

	// BackendStickTableTypeBinary captures enum value "binary"
	BackendStickTableTypeBinary string = "binary"
)

// prop value enum
func (m *BackendStickTable) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendStickTableTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BackendStickTable) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("stick_table"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackendStickTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackendStickTable) UnmarshalBinary(b []byte) error {
	var res BackendStickTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
