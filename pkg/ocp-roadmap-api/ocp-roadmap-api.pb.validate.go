// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-roadmap-api/ocp-roadmap-api.proto

package ocp_roadmap_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on Roadmap with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Roadmap) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return RoadmapValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	if m.GetUserId() <= 0 {
		return RoadmapValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for Link

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoadmapValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RoadmapValidationError is the validation error returned by Roadmap.Validate
// if the designated constraints aren't met.
type RoadmapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoadmapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoadmapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoadmapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoadmapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoadmapValidationError) ErrorName() string { return "RoadmapValidationError" }

// Error satisfies the builtin error interface
func (e RoadmapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoadmap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoadmapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoadmapValidationError{}

// Validate checks the field values on CreateRoadmapRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateRoadmapRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRoadmap()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRoadmapRequestValidationError{
				field:  "Roadmap",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateRoadmapRequestValidationError is the validation error returned by
// CreateRoadmapRequest.Validate if the designated constraints aren't met.
type CreateRoadmapRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoadmapRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoadmapRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoadmapRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoadmapRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoadmapRequestValidationError) ErrorName() string {
	return "CreateRoadmapRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRoadmapRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoadmapRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoadmapRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoadmapRequestValidationError{}

// Validate checks the field values on CreateRoadmapResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateRoadmapResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RoadmapId

	return nil
}

// CreateRoadmapResponseValidationError is the validation error returned by
// CreateRoadmapResponse.Validate if the designated constraints aren't met.
type CreateRoadmapResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoadmapResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoadmapResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoadmapResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoadmapResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoadmapResponseValidationError) ErrorName() string {
	return "CreateRoadmapResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRoadmapResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoadmapResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoadmapResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoadmapResponseValidationError{}

// Validate checks the field values on UpdateRoadmapRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateRoadmapRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for Link

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateRoadmapRequestValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateRoadmapRequestValidationError is the validation error returned by
// UpdateRoadmapRequest.Validate if the designated constraints aren't met.
type UpdateRoadmapRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRoadmapRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRoadmapRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRoadmapRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRoadmapRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRoadmapRequestValidationError) ErrorName() string {
	return "UpdateRoadmapRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRoadmapRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRoadmapRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRoadmapRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRoadmapRequestValidationError{}

// Validate checks the field values on UpdateRoadmapResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateRoadmapResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Updated

	return nil
}

// UpdateRoadmapResponseValidationError is the validation error returned by
// UpdateRoadmapResponse.Validate if the designated constraints aren't met.
type UpdateRoadmapResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRoadmapResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRoadmapResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRoadmapResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRoadmapResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRoadmapResponseValidationError) ErrorName() string {
	return "UpdateRoadmapResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRoadmapResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRoadmapResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRoadmapResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRoadmapResponseValidationError{}

// Validate checks the field values on MultiCreateRoadmapRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateRoadmapRequest) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRoadmaps() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiCreateRoadmapRequestValidationError{
					field:  fmt.Sprintf("Roadmaps[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MultiCreateRoadmapRequestValidationError is the validation error returned by
// MultiCreateRoadmapRequest.Validate if the designated constraints aren't met.
type MultiCreateRoadmapRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateRoadmapRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateRoadmapRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateRoadmapRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateRoadmapRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateRoadmapRequestValidationError) ErrorName() string {
	return "MultiCreateRoadmapRequestValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateRoadmapRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateRoadmapRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateRoadmapRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateRoadmapRequestValidationError{}

// Validate checks the field values on MultiCreateRoadmapResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateRoadmapResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// MultiCreateRoadmapResponseValidationError is the validation error returned
// by MultiCreateRoadmapResponse.Validate if the designated constraints aren't met.
type MultiCreateRoadmapResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateRoadmapResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateRoadmapResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateRoadmapResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateRoadmapResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateRoadmapResponseValidationError) ErrorName() string {
	return "MultiCreateRoadmapResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateRoadmapResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateRoadmapResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateRoadmapResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateRoadmapResponseValidationError{}

// Validate checks the field values on DescribeRoadmapRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeRoadmapRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DescribeRoadmapRequestValidationError is the validation error returned by
// DescribeRoadmapRequest.Validate if the designated constraints aren't met.
type DescribeRoadmapRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeRoadmapRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeRoadmapRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeRoadmapRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeRoadmapRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeRoadmapRequestValidationError) ErrorName() string {
	return "DescribeRoadmapRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeRoadmapRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeRoadmapRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeRoadmapRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeRoadmapRequestValidationError{}

// Validate checks the field values on DescribeRoadmapResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeRoadmapResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRoadmap()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeRoadmapResponseValidationError{
				field:  "Roadmap",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeRoadmapResponseValidationError is the validation error returned by
// DescribeRoadmapResponse.Validate if the designated constraints aren't met.
type DescribeRoadmapResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeRoadmapResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeRoadmapResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeRoadmapResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeRoadmapResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeRoadmapResponseValidationError) ErrorName() string {
	return "DescribeRoadmapResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeRoadmapResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeRoadmapResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeRoadmapResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeRoadmapResponseValidationError{}

// Validate checks the field values on ListRoadmapRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListRoadmapRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Limit

	// no validation rules for Offset

	return nil
}

// ListRoadmapRequestValidationError is the validation error returned by
// ListRoadmapRequest.Validate if the designated constraints aren't met.
type ListRoadmapRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRoadmapRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRoadmapRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRoadmapRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRoadmapRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRoadmapRequestValidationError) ErrorName() string {
	return "ListRoadmapRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListRoadmapRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRoadmapRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRoadmapRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRoadmapRequestValidationError{}

// Validate checks the field values on ListRoadmapResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListRoadmapResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRoadmaps() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListRoadmapResponseValidationError{
					field:  fmt.Sprintf("Roadmaps[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListRoadmapResponseValidationError is the validation error returned by
// ListRoadmapResponse.Validate if the designated constraints aren't met.
type ListRoadmapResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRoadmapResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRoadmapResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRoadmapResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRoadmapResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRoadmapResponseValidationError) ErrorName() string {
	return "ListRoadmapResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListRoadmapResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRoadmapResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRoadmapResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRoadmapResponseValidationError{}

// Validate checks the field values on RemoveRoadmapRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveRoadmapRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// RemoveRoadmapRequestValidationError is the validation error returned by
// RemoveRoadmapRequest.Validate if the designated constraints aren't met.
type RemoveRoadmapRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveRoadmapRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveRoadmapRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveRoadmapRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveRoadmapRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveRoadmapRequestValidationError) ErrorName() string {
	return "RemoveRoadmapRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveRoadmapRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveRoadmapRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveRoadmapRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveRoadmapRequestValidationError{}

// Validate checks the field values on RemoveRoadmapResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveRoadmapResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Removed

	return nil
}

// RemoveRoadmapResponseValidationError is the validation error returned by
// RemoveRoadmapResponse.Validate if the designated constraints aren't met.
type RemoveRoadmapResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveRoadmapResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveRoadmapResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveRoadmapResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveRoadmapResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveRoadmapResponseValidationError) ErrorName() string {
	return "RemoveRoadmapResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveRoadmapResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveRoadmapResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveRoadmapResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveRoadmapResponseValidationError{}