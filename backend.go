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

// Backend Backend
//
// HAProxy backend configuration
// swagger:model backend
type Backend struct {

	// abort on close
	// Enum: [enabled disabled]
	AbortOnClose string `json:"abort_on_close,omitempty"`

	// adv check
	// Enum: [http ssl-hello smtp ldap mysql pgsql tcp]
	AdvCheck string `json:"adv_check,omitempty"`

	// adv check http method
	// Enum: [HEAD PUT POST GET TRACE PATCH]
	AdvCheckHTTPMethod *string `json:"adv_check_http_method,omitempty"`

	// adv check http uri
	AdvCheckHTTPURI *string `json:"adv_check_http_uri,omitempty"`

	// adv check http version
	AdvCheckHTTPVersion string `json:"adv_check_http_version,omitempty"`

	// balance
	// Enum: [roundrobin least-connections hash-uri hash-source]
	Balance string `json:"balance,omitempty"`

	// check fail
	CheckFail *int64 `json:"check_fail,omitempty"`

	// check interval
	CheckInterval *int64 `json:"check_interval,omitempty"`

	// check port
	// Maximum: 65535
	// Minimum: 0
	CheckPort *int64 `json:"check_port,omitempty"`

	// check rise
	CheckRise *int64 `json:"check_rise,omitempty"`

	// check timeout
	CheckTimeout *int64 `json:"check_timeout,omitempty"`

	// connect failure redispatch
	// Enum: [enabled disabled]
	ConnectFailureRedispatch string `json:"connect_failure_redispatch,omitempty"`

	// connect retries
	ConnectRetries *int64 `json:"connect_retries,omitempty"`

	// connect source
	ConnectSource string `json:"connect_source,omitempty"`

	// connect timeout
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`

	// connect transparent
	// Enum: [enabled disabled]
	ConnectTransparent string `json:"connect_transparent,omitempty"`

	// continuous statistics
	// Enum: [enabled]
	ContinuousStatistics string `json:"continuous_statistics,omitempty"`

	// hash type
	// Enum: [map-based-sdbm map-based-djb2 map-based-wt6 consistent-sdbm consistent-djb2 consistent-wt6]
	HashType string `json:"hash_type,omitempty"`

	// hash type modifier
	// Enum: [avalanche]
	HashTypeModifier string `json:"hash_type_modifier,omitempty"`

	// http connection mode
	// Enum: [tunel passive-close forced-close server-close keep-alive]
	HTTPConnectionMode string `json:"http_connection_mode,omitempty"`

	// http cookie
	// Enum: [enabled]
	HTTPCookie string `json:"http_cookie,omitempty"`

	// http cookie mode
	// Enum: [passive passive-silent reset set set-silent session-prefix insert-only insert-only-silent passive-session-prefix]
	HTTPCookieMode string `json:"http_cookie_mode,omitempty"`

	// http cookie name
	HTTPCookieName string `json:"http_cookie_name,omitempty"`

	// http cookie nocache
	// Enum: [enabled disabled]
	HTTPCookieNocache string `json:"http_cookie_nocache,omitempty"`

	// http keepalive timeout
	HTTPKeepaliveTimeout *int64 `json:"http_keepalive_timeout,omitempty"`

	// http pretend keepalive
	// Enum: [enabled disabled]
	HTTPPretendKeepalive string `json:"http_pretend_keepalive,omitempty"`

	// http request timeout
	HTTPRequestTimeout *int64 `json:"http_request_timeout,omitempty"`

	// http xff header insert
	// Enum: [enabled]
	HTTPXffHeaderInsert string `json:"http_xff_header_insert,omitempty"`

	// http xff header insert except
	HTTPXffHeaderInsertExcept string `json:"http_xff_header_insert_except,omitempty"`

	// http xff header insert ifnone
	// Enum: [enabled]
	HTTPXffHeaderInsertIfnone string `json:"http_xff_header_insert_ifnone,omitempty"`

	// http xff header insert name
	HTTPXffHeaderInsertName string `json:"http_xff_header_insert_name,omitempty"`

	// log
	// Enum: [enabled]
	Log string `json:"log,omitempty"`

	// log format
	// Enum: [tcp http clf]
	LogFormat *string `json:"log_format,omitempty"`

	// log ignore null
	// Enum: [enabled disabled]
	LogIgnoreNull string `json:"log_ignore_null,omitempty"`

	// log tag
	LogTag string `json:"log_tag,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// protocol
	// Enum: [http tcp]
	Protocol string `json:"protocol,omitempty"`

	// queued timeout
	QueuedTimeout *int64 `json:"queued_timeout,omitempty"`

	// server inactivity timeout
	ServerInactivityTimeout *int64 `json:"server_inactivity_timeout,omitempty"`

	// server tcp keep alive
	// Enum: [enabled disabled]
	ServerTCPKeepAlive string `json:"server_tcp_keep_alive,omitempty"`

	// stick table
	// Enum: [ip ipv6 integer string binary]
	StickTable string `json:"stick_table,omitempty"`

	// stick table expire
	StickTableExpire *int64 `json:"stick_table_expire,omitempty"`

	// stick table keylen
	StickTableKeylen *int64 `json:"stick_table_keylen,omitempty"`

	// stick table nopurge
	// Enum: [enabled]
	StickTableNopurge string `json:"stick_table_nopurge,omitempty"`

	// stick table peers
	StickTablePeers string `json:"stick_table_peers,omitempty"`

	// stick table size
	StickTableSize *int64 `json:"stick_table_size,omitempty"`

	// stick table store
	StickTableStore string `json:"stick_table_store,omitempty"`

	// tcp smart connect
	// Enum: [enabled disabled]
	TCPSmartConnect string `json:"tcp_smart_connect,omitempty"`

	// tcpreq inspect delay
	TcpreqInspectDelay *int64 `json:"tcpreq_inspect_delay,omitempty"`

	// tcprsp inspect delay
	TcprspInspectDelay *int64 `json:"tcprsp_inspect_delay,omitempty"`

	// tunnel timeout
	TunnelTimeout *int64 `json:"tunnel_timeout,omitempty"`
}

// Validate validates this backend
func (m *Backend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbortOnClose(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdvCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdvCheckHTTPMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCheckPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectFailureRedispatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectTransparent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContinuousStatistics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHashType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHashTypeModifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPConnectionMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPCookie(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPCookieMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPCookieNocache(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPPretendKeepalive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPXffHeaderInsert(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPXffHeaderInsertIfnone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogIgnoreNull(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerTCPKeepAlive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStickTable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStickTableNopurge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTCPSmartConnect(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var backendTypeAbortOnClosePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeAbortOnClosePropEnum = append(backendTypeAbortOnClosePropEnum, v)
	}
}

const (

	// BackendAbortOnCloseEnabled captures enum value "enabled"
	BackendAbortOnCloseEnabled string = "enabled"

	// BackendAbortOnCloseDisabled captures enum value "disabled"
	BackendAbortOnCloseDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateAbortOnCloseEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeAbortOnClosePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateAbortOnClose(formats strfmt.Registry) error {

	if swag.IsZero(m.AbortOnClose) { // not required
		return nil
	}

	// value enum
	if err := m.validateAbortOnCloseEnum("abort_on_close", "body", m.AbortOnClose); err != nil {
		return err
	}

	return nil
}

var backendTypeAdvCheckPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","ssl-hello","smtp","ldap","mysql","pgsql","tcp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeAdvCheckPropEnum = append(backendTypeAdvCheckPropEnum, v)
	}
}

const (

	// BackendAdvCheckHTTP captures enum value "http"
	BackendAdvCheckHTTP string = "http"

	// BackendAdvCheckSslHello captures enum value "ssl-hello"
	BackendAdvCheckSslHello string = "ssl-hello"

	// BackendAdvCheckSMTP captures enum value "smtp"
	BackendAdvCheckSMTP string = "smtp"

	// BackendAdvCheckLdap captures enum value "ldap"
	BackendAdvCheckLdap string = "ldap"

	// BackendAdvCheckMysql captures enum value "mysql"
	BackendAdvCheckMysql string = "mysql"

	// BackendAdvCheckPgsql captures enum value "pgsql"
	BackendAdvCheckPgsql string = "pgsql"

	// BackendAdvCheckTCP captures enum value "tcp"
	BackendAdvCheckTCP string = "tcp"
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

var backendTypeAdvCheckHTTPMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HEAD","PUT","POST","GET","TRACE","PATCH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeAdvCheckHTTPMethodPropEnum = append(backendTypeAdvCheckHTTPMethodPropEnum, v)
	}
}

const (

	// BackendAdvCheckHTTPMethodHEAD captures enum value "HEAD"
	BackendAdvCheckHTTPMethodHEAD string = "HEAD"

	// BackendAdvCheckHTTPMethodPUT captures enum value "PUT"
	BackendAdvCheckHTTPMethodPUT string = "PUT"

	// BackendAdvCheckHTTPMethodPOST captures enum value "POST"
	BackendAdvCheckHTTPMethodPOST string = "POST"

	// BackendAdvCheckHTTPMethodGET captures enum value "GET"
	BackendAdvCheckHTTPMethodGET string = "GET"

	// BackendAdvCheckHTTPMethodTRACE captures enum value "TRACE"
	BackendAdvCheckHTTPMethodTRACE string = "TRACE"

	// BackendAdvCheckHTTPMethodPATCH captures enum value "PATCH"
	BackendAdvCheckHTTPMethodPATCH string = "PATCH"
)

// prop value enum
func (m *Backend) validateAdvCheckHTTPMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeAdvCheckHTTPMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateAdvCheckHTTPMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.AdvCheckHTTPMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateAdvCheckHTTPMethodEnum("adv_check_http_method", "body", *m.AdvCheckHTTPMethod); err != nil {
		return err
	}

	return nil
}

var backendTypeBalancePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["roundrobin","least-connections","hash-uri","hash-source"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeBalancePropEnum = append(backendTypeBalancePropEnum, v)
	}
}

const (

	// BackendBalanceRoundrobin captures enum value "roundrobin"
	BackendBalanceRoundrobin string = "roundrobin"

	// BackendBalanceLeastConnections captures enum value "least-connections"
	BackendBalanceLeastConnections string = "least-connections"

	// BackendBalanceHashURI captures enum value "hash-uri"
	BackendBalanceHashURI string = "hash-uri"

	// BackendBalanceHashSource captures enum value "hash-source"
	BackendBalanceHashSource string = "hash-source"
)

// prop value enum
func (m *Backend) validateBalanceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeBalancePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateBalance(formats strfmt.Registry) error {

	if swag.IsZero(m.Balance) { // not required
		return nil
	}

	// value enum
	if err := m.validateBalanceEnum("balance", "body", m.Balance); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateCheckPort(formats strfmt.Registry) error {

	if swag.IsZero(m.CheckPort) { // not required
		return nil
	}

	if err := validate.MinimumInt("check_port", "body", int64(*m.CheckPort), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("check_port", "body", int64(*m.CheckPort), 65535, false); err != nil {
		return err
	}

	return nil
}

var backendTypeConnectFailureRedispatchPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeConnectFailureRedispatchPropEnum = append(backendTypeConnectFailureRedispatchPropEnum, v)
	}
}

const (

	// BackendConnectFailureRedispatchEnabled captures enum value "enabled"
	BackendConnectFailureRedispatchEnabled string = "enabled"

	// BackendConnectFailureRedispatchDisabled captures enum value "disabled"
	BackendConnectFailureRedispatchDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateConnectFailureRedispatchEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeConnectFailureRedispatchPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateConnectFailureRedispatch(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectFailureRedispatch) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectFailureRedispatchEnum("connect_failure_redispatch", "body", m.ConnectFailureRedispatch); err != nil {
		return err
	}

	return nil
}

var backendTypeConnectTransparentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeConnectTransparentPropEnum = append(backendTypeConnectTransparentPropEnum, v)
	}
}

const (

	// BackendConnectTransparentEnabled captures enum value "enabled"
	BackendConnectTransparentEnabled string = "enabled"

	// BackendConnectTransparentDisabled captures enum value "disabled"
	BackendConnectTransparentDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateConnectTransparentEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeConnectTransparentPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateConnectTransparent(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectTransparent) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectTransparentEnum("connect_transparent", "body", m.ConnectTransparent); err != nil {
		return err
	}

	return nil
}

var backendTypeContinuousStatisticsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeContinuousStatisticsPropEnum = append(backendTypeContinuousStatisticsPropEnum, v)
	}
}

const (

	// BackendContinuousStatisticsEnabled captures enum value "enabled"
	BackendContinuousStatisticsEnabled string = "enabled"
)

// prop value enum
func (m *Backend) validateContinuousStatisticsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeContinuousStatisticsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateContinuousStatistics(formats strfmt.Registry) error {

	if swag.IsZero(m.ContinuousStatistics) { // not required
		return nil
	}

	// value enum
	if err := m.validateContinuousStatisticsEnum("continuous_statistics", "body", m.ContinuousStatistics); err != nil {
		return err
	}

	return nil
}

var backendTypeHashTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["map-based-sdbm","map-based-djb2","map-based-wt6","consistent-sdbm","consistent-djb2","consistent-wt6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHashTypePropEnum = append(backendTypeHashTypePropEnum, v)
	}
}

const (

	// BackendHashTypeMapBasedSdbm captures enum value "map-based-sdbm"
	BackendHashTypeMapBasedSdbm string = "map-based-sdbm"

	// BackendHashTypeMapBasedDjb2 captures enum value "map-based-djb2"
	BackendHashTypeMapBasedDjb2 string = "map-based-djb2"

	// BackendHashTypeMapBasedWt6 captures enum value "map-based-wt6"
	BackendHashTypeMapBasedWt6 string = "map-based-wt6"

	// BackendHashTypeConsistentSdbm captures enum value "consistent-sdbm"
	BackendHashTypeConsistentSdbm string = "consistent-sdbm"

	// BackendHashTypeConsistentDjb2 captures enum value "consistent-djb2"
	BackendHashTypeConsistentDjb2 string = "consistent-djb2"

	// BackendHashTypeConsistentWt6 captures enum value "consistent-wt6"
	BackendHashTypeConsistentWt6 string = "consistent-wt6"
)

// prop value enum
func (m *Backend) validateHashTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHashTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHashType(formats strfmt.Registry) error {

	if swag.IsZero(m.HashType) { // not required
		return nil
	}

	// value enum
	if err := m.validateHashTypeEnum("hash_type", "body", m.HashType); err != nil {
		return err
	}

	return nil
}

var backendTypeHashTypeModifierPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["avalanche"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHashTypeModifierPropEnum = append(backendTypeHashTypeModifierPropEnum, v)
	}
}

const (

	// BackendHashTypeModifierAvalanche captures enum value "avalanche"
	BackendHashTypeModifierAvalanche string = "avalanche"
)

// prop value enum
func (m *Backend) validateHashTypeModifierEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHashTypeModifierPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHashTypeModifier(formats strfmt.Registry) error {

	if swag.IsZero(m.HashTypeModifier) { // not required
		return nil
	}

	// value enum
	if err := m.validateHashTypeModifierEnum("hash_type_modifier", "body", m.HashTypeModifier); err != nil {
		return err
	}

	return nil
}

var backendTypeHTTPConnectionModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tunel","passive-close","forced-close","server-close","keep-alive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHTTPConnectionModePropEnum = append(backendTypeHTTPConnectionModePropEnum, v)
	}
}

const (

	// BackendHTTPConnectionModeTunel captures enum value "tunel"
	BackendHTTPConnectionModeTunel string = "tunel"

	// BackendHTTPConnectionModePassiveClose captures enum value "passive-close"
	BackendHTTPConnectionModePassiveClose string = "passive-close"

	// BackendHTTPConnectionModeForcedClose captures enum value "forced-close"
	BackendHTTPConnectionModeForcedClose string = "forced-close"

	// BackendHTTPConnectionModeServerClose captures enum value "server-close"
	BackendHTTPConnectionModeServerClose string = "server-close"

	// BackendHTTPConnectionModeKeepAlive captures enum value "keep-alive"
	BackendHTTPConnectionModeKeepAlive string = "keep-alive"
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

var backendTypeHTTPCookiePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHTTPCookiePropEnum = append(backendTypeHTTPCookiePropEnum, v)
	}
}

const (

	// BackendHTTPCookieEnabled captures enum value "enabled"
	BackendHTTPCookieEnabled string = "enabled"
)

// prop value enum
func (m *Backend) validateHTTPCookieEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHTTPCookiePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHTTPCookie(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPCookie) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPCookieEnum("http_cookie", "body", m.HTTPCookie); err != nil {
		return err
	}

	return nil
}

var backendTypeHTTPCookieModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["passive","passive-silent","reset","set","set-silent","session-prefix","insert-only","insert-only-silent","passive-session-prefix"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHTTPCookieModePropEnum = append(backendTypeHTTPCookieModePropEnum, v)
	}
}

const (

	// BackendHTTPCookieModePassive captures enum value "passive"
	BackendHTTPCookieModePassive string = "passive"

	// BackendHTTPCookieModePassiveSilent captures enum value "passive-silent"
	BackendHTTPCookieModePassiveSilent string = "passive-silent"

	// BackendHTTPCookieModeReset captures enum value "reset"
	BackendHTTPCookieModeReset string = "reset"

	// BackendHTTPCookieModeSet captures enum value "set"
	BackendHTTPCookieModeSet string = "set"

	// BackendHTTPCookieModeSetSilent captures enum value "set-silent"
	BackendHTTPCookieModeSetSilent string = "set-silent"

	// BackendHTTPCookieModeSessionPrefix captures enum value "session-prefix"
	BackendHTTPCookieModeSessionPrefix string = "session-prefix"

	// BackendHTTPCookieModeInsertOnly captures enum value "insert-only"
	BackendHTTPCookieModeInsertOnly string = "insert-only"

	// BackendHTTPCookieModeInsertOnlySilent captures enum value "insert-only-silent"
	BackendHTTPCookieModeInsertOnlySilent string = "insert-only-silent"

	// BackendHTTPCookieModePassiveSessionPrefix captures enum value "passive-session-prefix"
	BackendHTTPCookieModePassiveSessionPrefix string = "passive-session-prefix"
)

// prop value enum
func (m *Backend) validateHTTPCookieModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHTTPCookieModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHTTPCookieMode(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPCookieMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPCookieModeEnum("http_cookie_mode", "body", m.HTTPCookieMode); err != nil {
		return err
	}

	return nil
}

var backendTypeHTTPCookieNocachePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHTTPCookieNocachePropEnum = append(backendTypeHTTPCookieNocachePropEnum, v)
	}
}

const (

	// BackendHTTPCookieNocacheEnabled captures enum value "enabled"
	BackendHTTPCookieNocacheEnabled string = "enabled"

	// BackendHTTPCookieNocacheDisabled captures enum value "disabled"
	BackendHTTPCookieNocacheDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateHTTPCookieNocacheEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHTTPCookieNocachePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHTTPCookieNocache(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPCookieNocache) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPCookieNocacheEnum("http_cookie_nocache", "body", m.HTTPCookieNocache); err != nil {
		return err
	}

	return nil
}

var backendTypeHTTPPretendKeepalivePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHTTPPretendKeepalivePropEnum = append(backendTypeHTTPPretendKeepalivePropEnum, v)
	}
}

const (

	// BackendHTTPPretendKeepaliveEnabled captures enum value "enabled"
	BackendHTTPPretendKeepaliveEnabled string = "enabled"

	// BackendHTTPPretendKeepaliveDisabled captures enum value "disabled"
	BackendHTTPPretendKeepaliveDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateHTTPPretendKeepaliveEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHTTPPretendKeepalivePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHTTPPretendKeepalive(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPPretendKeepalive) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPPretendKeepaliveEnum("http_pretend_keepalive", "body", m.HTTPPretendKeepalive); err != nil {
		return err
	}

	return nil
}

var backendTypeHTTPXffHeaderInsertPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHTTPXffHeaderInsertPropEnum = append(backendTypeHTTPXffHeaderInsertPropEnum, v)
	}
}

const (

	// BackendHTTPXffHeaderInsertEnabled captures enum value "enabled"
	BackendHTTPXffHeaderInsertEnabled string = "enabled"
)

// prop value enum
func (m *Backend) validateHTTPXffHeaderInsertEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHTTPXffHeaderInsertPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHTTPXffHeaderInsert(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPXffHeaderInsert) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPXffHeaderInsertEnum("http_xff_header_insert", "body", m.HTTPXffHeaderInsert); err != nil {
		return err
	}

	return nil
}

var backendTypeHTTPXffHeaderInsertIfnonePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHTTPXffHeaderInsertIfnonePropEnum = append(backendTypeHTTPXffHeaderInsertIfnonePropEnum, v)
	}
}

const (

	// BackendHTTPXffHeaderInsertIfnoneEnabled captures enum value "enabled"
	BackendHTTPXffHeaderInsertIfnoneEnabled string = "enabled"
)

// prop value enum
func (m *Backend) validateHTTPXffHeaderInsertIfnoneEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHTTPXffHeaderInsertIfnonePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHTTPXffHeaderInsertIfnone(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPXffHeaderInsertIfnone) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPXffHeaderInsertIfnoneEnum("http_xff_header_insert_ifnone", "body", m.HTTPXffHeaderInsertIfnone); err != nil {
		return err
	}

	return nil
}

var backendTypeLogPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeLogPropEnum = append(backendTypeLogPropEnum, v)
	}
}

const (

	// BackendLogEnabled captures enum value "enabled"
	BackendLogEnabled string = "enabled"
)

// prop value enum
func (m *Backend) validateLogEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeLogPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateLog(formats strfmt.Registry) error {

	if swag.IsZero(m.Log) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogEnum("log", "body", m.Log); err != nil {
		return err
	}

	return nil
}

var backendTypeLogFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","http","clf"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeLogFormatPropEnum = append(backendTypeLogFormatPropEnum, v)
	}
}

const (

	// BackendLogFormatTCP captures enum value "tcp"
	BackendLogFormatTCP string = "tcp"

	// BackendLogFormatHTTP captures enum value "http"
	BackendLogFormatHTTP string = "http"

	// BackendLogFormatClf captures enum value "clf"
	BackendLogFormatClf string = "clf"
)

// prop value enum
func (m *Backend) validateLogFormatEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeLogFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateLogFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.LogFormat) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogFormatEnum("log_format", "body", *m.LogFormat); err != nil {
		return err
	}

	return nil
}

var backendTypeLogIgnoreNullPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeLogIgnoreNullPropEnum = append(backendTypeLogIgnoreNullPropEnum, v)
	}
}

const (

	// BackendLogIgnoreNullEnabled captures enum value "enabled"
	BackendLogIgnoreNullEnabled string = "enabled"

	// BackendLogIgnoreNullDisabled captures enum value "disabled"
	BackendLogIgnoreNullDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateLogIgnoreNullEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeLogIgnoreNullPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateLogIgnoreNull(formats strfmt.Registry) error {

	if swag.IsZero(m.LogIgnoreNull) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogIgnoreNullEnum("log_ignore_null", "body", m.LogIgnoreNull); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var backendTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","tcp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeProtocolPropEnum = append(backendTypeProtocolPropEnum, v)
	}
}

const (

	// BackendProtocolHTTP captures enum value "http"
	BackendProtocolHTTP string = "http"

	// BackendProtocolTCP captures enum value "tcp"
	BackendProtocolTCP string = "tcp"
)

// prop value enum
func (m *Backend) validateProtocolEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

var backendTypeServerTCPKeepAlivePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeServerTCPKeepAlivePropEnum = append(backendTypeServerTCPKeepAlivePropEnum, v)
	}
}

const (

	// BackendServerTCPKeepAliveEnabled captures enum value "enabled"
	BackendServerTCPKeepAliveEnabled string = "enabled"

	// BackendServerTCPKeepAliveDisabled captures enum value "disabled"
	BackendServerTCPKeepAliveDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateServerTCPKeepAliveEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeServerTCPKeepAlivePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateServerTCPKeepAlive(formats strfmt.Registry) error {

	if swag.IsZero(m.ServerTCPKeepAlive) { // not required
		return nil
	}

	// value enum
	if err := m.validateServerTCPKeepAliveEnum("server_tcp_keep_alive", "body", m.ServerTCPKeepAlive); err != nil {
		return err
	}

	return nil
}

var backendTypeStickTablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ip","ipv6","integer","string","binary"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeStickTablePropEnum = append(backendTypeStickTablePropEnum, v)
	}
}

const (

	// BackendStickTableIP captures enum value "ip"
	BackendStickTableIP string = "ip"

	// BackendStickTableIPV6 captures enum value "ipv6"
	BackendStickTableIPV6 string = "ipv6"

	// BackendStickTableInteger captures enum value "integer"
	BackendStickTableInteger string = "integer"

	// BackendStickTableString captures enum value "string"
	BackendStickTableString string = "string"

	// BackendStickTableBinary captures enum value "binary"
	BackendStickTableBinary string = "binary"
)

// prop value enum
func (m *Backend) validateStickTableEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeStickTablePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateStickTable(formats strfmt.Registry) error {

	if swag.IsZero(m.StickTable) { // not required
		return nil
	}

	// value enum
	if err := m.validateStickTableEnum("stick_table", "body", m.StickTable); err != nil {
		return err
	}

	return nil
}

var backendTypeStickTableNopurgePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeStickTableNopurgePropEnum = append(backendTypeStickTableNopurgePropEnum, v)
	}
}

const (

	// BackendStickTableNopurgeEnabled captures enum value "enabled"
	BackendStickTableNopurgeEnabled string = "enabled"
)

// prop value enum
func (m *Backend) validateStickTableNopurgeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeStickTableNopurgePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateStickTableNopurge(formats strfmt.Registry) error {

	if swag.IsZero(m.StickTableNopurge) { // not required
		return nil
	}

	// value enum
	if err := m.validateStickTableNopurgeEnum("stick_table_nopurge", "body", m.StickTableNopurge); err != nil {
		return err
	}

	return nil
}

var backendTypeTCPSmartConnectPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeTCPSmartConnectPropEnum = append(backendTypeTCPSmartConnectPropEnum, v)
	}
}

const (

	// BackendTCPSmartConnectEnabled captures enum value "enabled"
	BackendTCPSmartConnectEnabled string = "enabled"

	// BackendTCPSmartConnectDisabled captures enum value "disabled"
	BackendTCPSmartConnectDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateTCPSmartConnectEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeTCPSmartConnectPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateTCPSmartConnect(formats strfmt.Registry) error {

	if swag.IsZero(m.TCPSmartConnect) { // not required
		return nil
	}

	// value enum
	if err := m.validateTCPSmartConnectEnum("tcp_smart_connect", "body", m.TCPSmartConnect); err != nil {
		return err
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
