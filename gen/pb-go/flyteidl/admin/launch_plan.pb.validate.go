// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/admin/launch_plan.proto

package admin

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _launch_plan_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on LaunchPlanCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LaunchPlanCreateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanCreateRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSpec()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanCreateRequestValidationError{
				field:  "Spec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// LaunchPlanCreateRequestValidationError is the validation error returned by
// LaunchPlanCreateRequest.Validate if the designated constraints aren't met.
type LaunchPlanCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LaunchPlanCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LaunchPlanCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LaunchPlanCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LaunchPlanCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LaunchPlanCreateRequestValidationError) ErrorName() string {
	return "LaunchPlanCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e LaunchPlanCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLaunchPlanCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LaunchPlanCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LaunchPlanCreateRequestValidationError{}

// Validate checks the field values on LaunchPlanCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LaunchPlanCreateResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// LaunchPlanCreateResponseValidationError is the validation error returned by
// LaunchPlanCreateResponse.Validate if the designated constraints aren't met.
type LaunchPlanCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LaunchPlanCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LaunchPlanCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LaunchPlanCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LaunchPlanCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LaunchPlanCreateResponseValidationError) ErrorName() string {
	return "LaunchPlanCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LaunchPlanCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLaunchPlanCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LaunchPlanCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LaunchPlanCreateResponseValidationError{}

// Validate checks the field values on LaunchPlan with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LaunchPlan) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSpec()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanValidationError{
				field:  "Spec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetClosure()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanValidationError{
				field:  "Closure",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// LaunchPlanValidationError is the validation error returned by
// LaunchPlan.Validate if the designated constraints aren't met.
type LaunchPlanValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LaunchPlanValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LaunchPlanValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LaunchPlanValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LaunchPlanValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LaunchPlanValidationError) ErrorName() string { return "LaunchPlanValidationError" }

// Error satisfies the builtin error interface
func (e LaunchPlanValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLaunchPlan.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LaunchPlanValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LaunchPlanValidationError{}

// Validate checks the field values on LaunchPlanList with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LaunchPlanList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetLaunchPlans() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LaunchPlanListValidationError{
					field:  fmt.Sprintf("LaunchPlans[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Token

	return nil
}

// LaunchPlanListValidationError is the validation error returned by
// LaunchPlanList.Validate if the designated constraints aren't met.
type LaunchPlanListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LaunchPlanListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LaunchPlanListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LaunchPlanListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LaunchPlanListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LaunchPlanListValidationError) ErrorName() string { return "LaunchPlanListValidationError" }

// Error satisfies the builtin error interface
func (e LaunchPlanListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLaunchPlanList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LaunchPlanListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LaunchPlanListValidationError{}

// Validate checks the field values on Auth with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Auth) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AssumableIamRole

	// no validation rules for KubernetesServiceAccount

	return nil
}

// AuthValidationError is the validation error returned by Auth.Validate if the
// designated constraints aren't met.
type AuthValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthValidationError) ErrorName() string { return "AuthValidationError" }

// Error satisfies the builtin error interface
func (e AuthValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuth.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthValidationError{}

// Validate checks the field values on LaunchPlanSpec with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LaunchPlanSpec) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetWorkflowId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanSpecValidationError{
				field:  "WorkflowId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEntityMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanSpecValidationError{
				field:  "EntityMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDefaultInputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanSpecValidationError{
				field:  "DefaultInputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFixedInputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanSpecValidationError{
				field:  "FixedInputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Role

	if v, ok := interface{}(m.GetLabels()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanSpecValidationError{
				field:  "Labels",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAnnotations()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanSpecValidationError{
				field:  "Annotations",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAuth()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanSpecValidationError{
				field:  "Auth",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAuthRole()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanSpecValidationError{
				field:  "AuthRole",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSecurityContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanSpecValidationError{
				field:  "SecurityContext",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetQualityOfService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanSpecValidationError{
				field:  "QualityOfService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRawOutputDataConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanSpecValidationError{
				field:  "RawOutputDataConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MaxParallelism

	if v, ok := interface{}(m.GetInterruptible()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanSpecValidationError{
				field:  "Interruptible",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OverwriteCache

	if v, ok := interface{}(m.GetEnvs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanSpecValidationError{
				field:  "Envs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for key, val := range m.GetTaskNodeRuntimeOverrides() {
		_ = val

		// no validation rules for TaskNodeRuntimeOverrides[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LaunchPlanSpecValidationError{
					field:  fmt.Sprintf("TaskNodeRuntimeOverrides[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LaunchPlanSpecValidationError is the validation error returned by
// LaunchPlanSpec.Validate if the designated constraints aren't met.
type LaunchPlanSpecValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LaunchPlanSpecValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LaunchPlanSpecValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LaunchPlanSpecValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LaunchPlanSpecValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LaunchPlanSpecValidationError) ErrorName() string { return "LaunchPlanSpecValidationError" }

// Error satisfies the builtin error interface
func (e LaunchPlanSpecValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLaunchPlanSpec.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LaunchPlanSpecValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LaunchPlanSpecValidationError{}

// Validate checks the field values on LaunchPlanClosure with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *LaunchPlanClosure) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for State

	if v, ok := interface{}(m.GetExpectedInputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanClosureValidationError{
				field:  "ExpectedInputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetExpectedOutputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanClosureValidationError{
				field:  "ExpectedOutputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanClosureValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanClosureValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// LaunchPlanClosureValidationError is the validation error returned by
// LaunchPlanClosure.Validate if the designated constraints aren't met.
type LaunchPlanClosureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LaunchPlanClosureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LaunchPlanClosureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LaunchPlanClosureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LaunchPlanClosureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LaunchPlanClosureValidationError) ErrorName() string {
	return "LaunchPlanClosureValidationError"
}

// Error satisfies the builtin error interface
func (e LaunchPlanClosureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLaunchPlanClosure.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LaunchPlanClosureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LaunchPlanClosureValidationError{}

// Validate checks the field values on LaunchPlanMetadata with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LaunchPlanMetadata) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSchedule()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanMetadataValidationError{
				field:  "Schedule",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetNotifications() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LaunchPlanMetadataValidationError{
					field:  fmt.Sprintf("Notifications[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LaunchPlanMetadataValidationError is the validation error returned by
// LaunchPlanMetadata.Validate if the designated constraints aren't met.
type LaunchPlanMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LaunchPlanMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LaunchPlanMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LaunchPlanMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LaunchPlanMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LaunchPlanMetadataValidationError) ErrorName() string {
	return "LaunchPlanMetadataValidationError"
}

// Error satisfies the builtin error interface
func (e LaunchPlanMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLaunchPlanMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LaunchPlanMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LaunchPlanMetadataValidationError{}

// Validate checks the field values on LaunchPlanUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LaunchPlanUpdateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LaunchPlanUpdateRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for State

	return nil
}

// LaunchPlanUpdateRequestValidationError is the validation error returned by
// LaunchPlanUpdateRequest.Validate if the designated constraints aren't met.
type LaunchPlanUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LaunchPlanUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LaunchPlanUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LaunchPlanUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LaunchPlanUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LaunchPlanUpdateRequestValidationError) ErrorName() string {
	return "LaunchPlanUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e LaunchPlanUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLaunchPlanUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LaunchPlanUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LaunchPlanUpdateRequestValidationError{}

// Validate checks the field values on LaunchPlanUpdateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LaunchPlanUpdateResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// LaunchPlanUpdateResponseValidationError is the validation error returned by
// LaunchPlanUpdateResponse.Validate if the designated constraints aren't met.
type LaunchPlanUpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LaunchPlanUpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LaunchPlanUpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LaunchPlanUpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LaunchPlanUpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LaunchPlanUpdateResponseValidationError) ErrorName() string {
	return "LaunchPlanUpdateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LaunchPlanUpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLaunchPlanUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LaunchPlanUpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LaunchPlanUpdateResponseValidationError{}

// Validate checks the field values on ActiveLaunchPlanRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ActiveLaunchPlanRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ActiveLaunchPlanRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ActiveLaunchPlanRequestValidationError is the validation error returned by
// ActiveLaunchPlanRequest.Validate if the designated constraints aren't met.
type ActiveLaunchPlanRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActiveLaunchPlanRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActiveLaunchPlanRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActiveLaunchPlanRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActiveLaunchPlanRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActiveLaunchPlanRequestValidationError) ErrorName() string {
	return "ActiveLaunchPlanRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ActiveLaunchPlanRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActiveLaunchPlanRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActiveLaunchPlanRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActiveLaunchPlanRequestValidationError{}

// Validate checks the field values on ActiveLaunchPlanListRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ActiveLaunchPlanListRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Project

	// no validation rules for Domain

	// no validation rules for Limit

	// no validation rules for Token

	if v, ok := interface{}(m.GetSortBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ActiveLaunchPlanListRequestValidationError{
				field:  "SortBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ActiveLaunchPlanListRequestValidationError is the validation error returned
// by ActiveLaunchPlanListRequest.Validate if the designated constraints
// aren't met.
type ActiveLaunchPlanListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActiveLaunchPlanListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActiveLaunchPlanListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActiveLaunchPlanListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActiveLaunchPlanListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActiveLaunchPlanListRequestValidationError) ErrorName() string {
	return "ActiveLaunchPlanListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ActiveLaunchPlanListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActiveLaunchPlanListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActiveLaunchPlanListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActiveLaunchPlanListRequestValidationError{}
