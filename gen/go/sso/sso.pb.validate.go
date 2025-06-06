// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: sso/sso.proto

package ssov1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on RegisterRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegisterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterRequestMultiError, or nil if none found.
func (m *RegisterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = RegisterRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetUsername()); l < 3 || l > 20 {
		err := RegisterRequestValidationError{
			field:  "Username",
			reason: "value length must be between 3 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_RegisterRequest_Username_Pattern.MatchString(m.GetUsername()) {
		err := RegisterRequestValidationError{
			field:  "Username",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9_]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 8 {
		err := RegisterRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 8 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_RegisterRequest_Password_Pattern.MatchString(m.GetPassword()) {
		err := RegisterRequestValidationError{
			field:  "Password",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9!@#$%^&*]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RegisterRequestMultiError(errors)
	}

	return nil
}

func (m *RegisterRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *RegisterRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// RegisterRequestMultiError is an error wrapping multiple validation errors
// returned by RegisterRequest.ValidateAll() if the designated constraints
// aren't met.
type RegisterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterRequestMultiError) AllErrors() []error { return m }

// RegisterRequestValidationError is the validation error returned by
// RegisterRequest.Validate if the designated constraints aren't met.
type RegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterRequestValidationError) ErrorName() string { return "RegisterRequestValidationError" }

// Error satisfies the builtin error interface
func (e RegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterRequestValidationError{}

var _RegisterRequest_Username_Pattern = regexp.MustCompile("^[a-zA-Z0-9_]+$")

var _RegisterRequest_Password_Pattern = regexp.MustCompile("^[a-zA-Z0-9!@#$%^&*]+$")

// Validate checks the field values on RegisterResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegisterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterResponseMultiError, or nil if none found.
func (m *RegisterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return RegisterResponseMultiError(errors)
	}

	return nil
}

// RegisterResponseMultiError is an error wrapping multiple validation errors
// returned by RegisterResponse.ValidateAll() if the designated constraints
// aren't met.
type RegisterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterResponseMultiError) AllErrors() []error { return m }

// RegisterResponseValidationError is the validation error returned by
// RegisterResponse.Validate if the designated constraints aren't met.
type RegisterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterResponseValidationError) ErrorName() string { return "RegisterResponseValidationError" }

// Error satisfies the builtin error interface
func (e RegisterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterResponseValidationError{}

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRequestMultiError, or
// nil if none found.
func (m *LoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetPassword()) < 8 {
		err := LoginRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 8 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_LoginRequest_Password_Pattern.MatchString(m.GetPassword()) {
		err := LoginRequestValidationError{
			field:  "Password",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9!@#$%^&*]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAppId() <= 0 {
		err := LoginRequestValidationError{
			field:  "AppId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	switch v := m.Identifier.(type) {
	case *LoginRequest_Email:
		if v == nil {
			err := LoginRequestValidationError{
				field:  "Identifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if err := m._validateEmail(m.GetEmail()); err != nil {
			err = LoginRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *LoginRequest_Username:
		if v == nil {
			err := LoginRequestValidationError{
				field:  "Identifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if l := utf8.RuneCountInString(m.GetUsername()); l < 3 || l > 20 {
			err := LoginRequestValidationError{
				field:  "Username",
				reason: "value length must be between 3 and 20 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if !_LoginRequest_Username_Pattern.MatchString(m.GetUsername()) {
			err := LoginRequestValidationError{
				field:  "Username",
				reason: "value does not match regex pattern \"^[a-zA-Z0-9_]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return LoginRequestMultiError(errors)
	}

	return nil
}

func (m *LoginRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *LoginRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// LoginRequestMultiError is an error wrapping multiple validation errors
// returned by LoginRequest.ValidateAll() if the designated constraints aren't met.
type LoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRequestMultiError) AllErrors() []error { return m }

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

var _LoginRequest_Username_Pattern = regexp.MustCompile("^[a-zA-Z0-9_]+$")

var _LoginRequest_Password_Pattern = regexp.MustCompile("^[a-zA-Z0-9!@#$%^&*]+$")

// Validate checks the field values on LoginResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginResponseMultiError, or
// nil if none found.
func (m *LoginResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return LoginResponseMultiError(errors)
	}

	return nil
}

// LoginResponseMultiError is an error wrapping multiple validation errors
// returned by LoginResponse.ValidateAll() if the designated constraints
// aren't met.
type LoginResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginResponseMultiError) AllErrors() []error { return m }

// LoginResponseValidationError is the validation error returned by
// LoginResponse.Validate if the designated constraints aren't met.
type LoginResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginResponseValidationError) ErrorName() string { return "LoginResponseValidationError" }

// Error satisfies the builtin error interface
func (e LoginResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginResponseValidationError{}

// Validate checks the field values on IsAdminRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IsAdminRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IsAdminRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IsAdminRequestMultiError,
// or nil if none found.
func (m *IsAdminRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IsAdminRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetRequstingUserId() <= 0 {
		err := IsAdminRequestValidationError{
			field:  "RequstingUserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTargetUserId() <= 0 {
		err := IsAdminRequestValidationError{
			field:  "TargetUserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IsAdminRequestMultiError(errors)
	}

	return nil
}

// IsAdminRequestMultiError is an error wrapping multiple validation errors
// returned by IsAdminRequest.ValidateAll() if the designated constraints
// aren't met.
type IsAdminRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IsAdminRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IsAdminRequestMultiError) AllErrors() []error { return m }

// IsAdminRequestValidationError is the validation error returned by
// IsAdminRequest.Validate if the designated constraints aren't met.
type IsAdminRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IsAdminRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IsAdminRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IsAdminRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IsAdminRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IsAdminRequestValidationError) ErrorName() string { return "IsAdminRequestValidationError" }

// Error satisfies the builtin error interface
func (e IsAdminRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIsAdminRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IsAdminRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IsAdminRequestValidationError{}

// Validate checks the field values on IsAdminResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IsAdminResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IsAdminResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IsAdminResponseMultiError, or nil if none found.
func (m *IsAdminResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *IsAdminResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IsAdmin

	if len(errors) > 0 {
		return IsAdminResponseMultiError(errors)
	}

	return nil
}

// IsAdminResponseMultiError is an error wrapping multiple validation errors
// returned by IsAdminResponse.ValidateAll() if the designated constraints
// aren't met.
type IsAdminResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IsAdminResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IsAdminResponseMultiError) AllErrors() []error { return m }

// IsAdminResponseValidationError is the validation error returned by
// IsAdminResponse.Validate if the designated constraints aren't met.
type IsAdminResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IsAdminResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IsAdminResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IsAdminResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IsAdminResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IsAdminResponseValidationError) ErrorName() string { return "IsAdminResponseValidationError" }

// Error satisfies the builtin error interface
func (e IsAdminResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIsAdminResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IsAdminResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IsAdminResponseValidationError{}

// Validate checks the field values on GrantAdminRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GrantAdminRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GrantAdminRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GrantAdminRequestMultiError, or nil if none found.
func (m *GrantAdminRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GrantAdminRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetRequstingUserId() <= 0 {
		err := GrantAdminRequestValidationError{
			field:  "RequstingUserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTargetUserId() <= 0 {
		err := GrantAdminRequestValidationError{
			field:  "TargetUserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GrantAdminRequestMultiError(errors)
	}

	return nil
}

// GrantAdminRequestMultiError is an error wrapping multiple validation errors
// returned by GrantAdminRequest.ValidateAll() if the designated constraints
// aren't met.
type GrantAdminRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrantAdminRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrantAdminRequestMultiError) AllErrors() []error { return m }

// GrantAdminRequestValidationError is the validation error returned by
// GrantAdminRequest.Validate if the designated constraints aren't met.
type GrantAdminRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrantAdminRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrantAdminRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrantAdminRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrantAdminRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrantAdminRequestValidationError) ErrorName() string {
	return "GrantAdminRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GrantAdminRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrantAdminRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrantAdminRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrantAdminRequestValidationError{}

// Validate checks the field values on GrantAdminResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GrantAdminResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GrantAdminResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GrantAdminResponseMultiError, or nil if none found.
func (m *GrantAdminResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GrantAdminResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return GrantAdminResponseMultiError(errors)
	}

	return nil
}

// GrantAdminResponseMultiError is an error wrapping multiple validation errors
// returned by GrantAdminResponse.ValidateAll() if the designated constraints
// aren't met.
type GrantAdminResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrantAdminResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrantAdminResponseMultiError) AllErrors() []error { return m }

// GrantAdminResponseValidationError is the validation error returned by
// GrantAdminResponse.Validate if the designated constraints aren't met.
type GrantAdminResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrantAdminResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrantAdminResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrantAdminResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrantAdminResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrantAdminResponseValidationError) ErrorName() string {
	return "GrantAdminResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GrantAdminResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrantAdminResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrantAdminResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrantAdminResponseValidationError{}

// Validate checks the field values on RevokeAdminRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RevokeAdminRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RevokeAdminRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RevokeAdminRequestMultiError, or nil if none found.
func (m *RevokeAdminRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RevokeAdminRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetRequstingUserId() <= 0 {
		err := RevokeAdminRequestValidationError{
			field:  "RequstingUserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTargetUserId() <= 0 {
		err := RevokeAdminRequestValidationError{
			field:  "TargetUserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RevokeAdminRequestMultiError(errors)
	}

	return nil
}

// RevokeAdminRequestMultiError is an error wrapping multiple validation errors
// returned by RevokeAdminRequest.ValidateAll() if the designated constraints
// aren't met.
type RevokeAdminRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RevokeAdminRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RevokeAdminRequestMultiError) AllErrors() []error { return m }

// RevokeAdminRequestValidationError is the validation error returned by
// RevokeAdminRequest.Validate if the designated constraints aren't met.
type RevokeAdminRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RevokeAdminRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RevokeAdminRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RevokeAdminRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RevokeAdminRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RevokeAdminRequestValidationError) ErrorName() string {
	return "RevokeAdminRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RevokeAdminRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRevokeAdminRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RevokeAdminRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RevokeAdminRequestValidationError{}

// Validate checks the field values on RevokeAdminResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RevokeAdminResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RevokeAdminResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RevokeAdminResponseMultiError, or nil if none found.
func (m *RevokeAdminResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RevokeAdminResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return RevokeAdminResponseMultiError(errors)
	}

	return nil
}

// RevokeAdminResponseMultiError is an error wrapping multiple validation
// errors returned by RevokeAdminResponse.ValidateAll() if the designated
// constraints aren't met.
type RevokeAdminResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RevokeAdminResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RevokeAdminResponseMultiError) AllErrors() []error { return m }

// RevokeAdminResponseValidationError is the validation error returned by
// RevokeAdminResponse.Validate if the designated constraints aren't met.
type RevokeAdminResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RevokeAdminResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RevokeAdminResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RevokeAdminResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RevokeAdminResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RevokeAdminResponseValidationError) ErrorName() string {
	return "RevokeAdminResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RevokeAdminResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRevokeAdminResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RevokeAdminResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RevokeAdminResponseValidationError{}
