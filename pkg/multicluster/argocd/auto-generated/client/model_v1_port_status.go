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

// V1PortStatus struct for V1PortStatus
type V1PortStatus struct {
	Error *string `json:"error,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
}

// NewV1PortStatus instantiates a new V1PortStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PortStatus() *V1PortStatus {
	this := V1PortStatus{}
	return &this
}

// NewV1PortStatusWithDefaults instantiates a new V1PortStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PortStatusWithDefaults() *V1PortStatus {
	this := V1PortStatus{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V1PortStatus) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PortStatus) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V1PortStatus) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *V1PortStatus) SetError(v string) {
	o.Error = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *V1PortStatus) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PortStatus) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *V1PortStatus) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *V1PortStatus) SetPort(v int32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *V1PortStatus) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PortStatus) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *V1PortStatus) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *V1PortStatus) SetProtocol(v string) {
	o.Protocol = &v
}

func (o V1PortStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	return json.Marshal(toSerialize)
}

type NullableV1PortStatus struct {
	value *V1PortStatus
	isSet bool
}

func (v NullableV1PortStatus) Get() *V1PortStatus {
	return v.value
}

func (v *NullableV1PortStatus) Set(val *V1PortStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PortStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PortStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PortStatus(val *V1PortStatus) *NullableV1PortStatus {
	return &NullableV1PortStatus{value: val, isSet: true}
}

func (v NullableV1PortStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PortStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


