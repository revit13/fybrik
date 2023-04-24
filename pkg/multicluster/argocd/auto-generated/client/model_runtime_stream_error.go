/*
Consolidate Services

Description of all APIs

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
)

// RuntimeStreamError struct for RuntimeStreamError
type RuntimeStreamError struct {
	Details *[]ProtobufAny `json:"details,omitempty"`
	GrpcCode *int32 `json:"grpc_code,omitempty"`
	HttpCode *int32 `json:"http_code,omitempty"`
	HttpStatus *string `json:"http_status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewRuntimeStreamError instantiates a new RuntimeStreamError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuntimeStreamError() *RuntimeStreamError {
	this := RuntimeStreamError{}
	return &this
}

// NewRuntimeStreamErrorWithDefaults instantiates a new RuntimeStreamError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuntimeStreamErrorWithDefaults() *RuntimeStreamError {
	this := RuntimeStreamError{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *RuntimeStreamError) GetDetails() []ProtobufAny {
	if o == nil || o.Details == nil {
		var ret []ProtobufAny
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeStreamError) GetDetailsOk() (*[]ProtobufAny, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *RuntimeStreamError) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []ProtobufAny and assigns it to the Details field.
func (o *RuntimeStreamError) SetDetails(v []ProtobufAny) {
	o.Details = &v
}

// GetGrpcCode returns the GrpcCode field value if set, zero value otherwise.
func (o *RuntimeStreamError) GetGrpcCode() int32 {
	if o == nil || o.GrpcCode == nil {
		var ret int32
		return ret
	}
	return *o.GrpcCode
}

// GetGrpcCodeOk returns a tuple with the GrpcCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeStreamError) GetGrpcCodeOk() (*int32, bool) {
	if o == nil || o.GrpcCode == nil {
		return nil, false
	}
	return o.GrpcCode, true
}

// HasGrpcCode returns a boolean if a field has been set.
func (o *RuntimeStreamError) HasGrpcCode() bool {
	if o != nil && o.GrpcCode != nil {
		return true
	}

	return false
}

// SetGrpcCode gets a reference to the given int32 and assigns it to the GrpcCode field.
func (o *RuntimeStreamError) SetGrpcCode(v int32) {
	o.GrpcCode = &v
}

// GetHttpCode returns the HttpCode field value if set, zero value otherwise.
func (o *RuntimeStreamError) GetHttpCode() int32 {
	if o == nil || o.HttpCode == nil {
		var ret int32
		return ret
	}
	return *o.HttpCode
}

// GetHttpCodeOk returns a tuple with the HttpCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeStreamError) GetHttpCodeOk() (*int32, bool) {
	if o == nil || o.HttpCode == nil {
		return nil, false
	}
	return o.HttpCode, true
}

// HasHttpCode returns a boolean if a field has been set.
func (o *RuntimeStreamError) HasHttpCode() bool {
	if o != nil && o.HttpCode != nil {
		return true
	}

	return false
}

// SetHttpCode gets a reference to the given int32 and assigns it to the HttpCode field.
func (o *RuntimeStreamError) SetHttpCode(v int32) {
	o.HttpCode = &v
}

// GetHttpStatus returns the HttpStatus field value if set, zero value otherwise.
func (o *RuntimeStreamError) GetHttpStatus() string {
	if o == nil || o.HttpStatus == nil {
		var ret string
		return ret
	}
	return *o.HttpStatus
}

// GetHttpStatusOk returns a tuple with the HttpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeStreamError) GetHttpStatusOk() (*string, bool) {
	if o == nil || o.HttpStatus == nil {
		return nil, false
	}
	return o.HttpStatus, true
}

// HasHttpStatus returns a boolean if a field has been set.
func (o *RuntimeStreamError) HasHttpStatus() bool {
	if o != nil && o.HttpStatus != nil {
		return true
	}

	return false
}

// SetHttpStatus gets a reference to the given string and assigns it to the HttpStatus field.
func (o *RuntimeStreamError) SetHttpStatus(v string) {
	o.HttpStatus = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RuntimeStreamError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeStreamError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RuntimeStreamError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RuntimeStreamError) SetMessage(v string) {
	o.Message = &v
}

func (o RuntimeStreamError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.GrpcCode != nil {
		toSerialize["grpc_code"] = o.GrpcCode
	}
	if o.HttpCode != nil {
		toSerialize["http_code"] = o.HttpCode
	}
	if o.HttpStatus != nil {
		toSerialize["http_status"] = o.HttpStatus
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableRuntimeStreamError struct {
	value *RuntimeStreamError
	isSet bool
}

func (v NullableRuntimeStreamError) Get() *RuntimeStreamError {
	return v.value
}

func (v *NullableRuntimeStreamError) Set(val *RuntimeStreamError) {
	v.value = val
	v.isSet = true
}

func (v NullableRuntimeStreamError) IsSet() bool {
	return v.isSet
}

func (v *NullableRuntimeStreamError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuntimeStreamError(val *RuntimeStreamError) *NullableRuntimeStreamError {
	return &NullableRuntimeStreamError{value: val, isSet: true}
}

func (v NullableRuntimeStreamError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuntimeStreamError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


