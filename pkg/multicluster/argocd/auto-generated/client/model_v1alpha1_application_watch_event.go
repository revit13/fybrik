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

// V1alpha1ApplicationWatchEvent ApplicationWatchEvent contains information about application change.
type V1alpha1ApplicationWatchEvent struct {
	Application *V1alpha1Application `json:"application,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewV1alpha1ApplicationWatchEvent instantiates a new V1alpha1ApplicationWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ApplicationWatchEvent() *V1alpha1ApplicationWatchEvent {
	this := V1alpha1ApplicationWatchEvent{}
	return &this
}

// NewV1alpha1ApplicationWatchEventWithDefaults instantiates a new V1alpha1ApplicationWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ApplicationWatchEventWithDefaults() *V1alpha1ApplicationWatchEvent {
	this := V1alpha1ApplicationWatchEvent{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *V1alpha1ApplicationWatchEvent) GetApplication() V1alpha1Application {
	if o == nil || o.Application == nil {
		var ret V1alpha1Application
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationWatchEvent) GetApplicationOk() (*V1alpha1Application, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *V1alpha1ApplicationWatchEvent) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given V1alpha1Application and assigns it to the Application field.
func (o *V1alpha1ApplicationWatchEvent) SetApplication(v V1alpha1Application) {
	o.Application = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1alpha1ApplicationWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1alpha1ApplicationWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1alpha1ApplicationWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o V1alpha1ApplicationWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1ApplicationWatchEvent struct {
	value *V1alpha1ApplicationWatchEvent
	isSet bool
}

func (v NullableV1alpha1ApplicationWatchEvent) Get() *V1alpha1ApplicationWatchEvent {
	return v.value
}

func (v *NullableV1alpha1ApplicationWatchEvent) Set(val *V1alpha1ApplicationWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ApplicationWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ApplicationWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ApplicationWatchEvent(val *V1alpha1ApplicationWatchEvent) *NullableV1alpha1ApplicationWatchEvent {
	return &NullableV1alpha1ApplicationWatchEvent{value: val, isSet: true}
}

func (v NullableV1alpha1ApplicationWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ApplicationWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


