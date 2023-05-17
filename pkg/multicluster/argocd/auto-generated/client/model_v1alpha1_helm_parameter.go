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

// V1alpha1HelmParameter struct for V1alpha1HelmParameter
type V1alpha1HelmParameter struct {
	ForceString *bool   `json:"forceString,omitempty"`
	Name        *string `json:"name,omitempty"`
	Value       *string `json:"value,omitempty"`
}

// NewV1alpha1HelmParameter instantiates a new V1alpha1HelmParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1HelmParameter() *V1alpha1HelmParameter {
	this := V1alpha1HelmParameter{}
	return &this
}

// NewV1alpha1HelmParameterWithDefaults instantiates a new V1alpha1HelmParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1HelmParameterWithDefaults() *V1alpha1HelmParameter {
	this := V1alpha1HelmParameter{}
	return &this
}

// GetForceString returns the ForceString field value if set, zero value otherwise.
func (o *V1alpha1HelmParameter) GetForceString() bool {
	if o == nil || o.ForceString == nil {
		var ret bool
		return ret
	}
	return *o.ForceString
}

// GetForceStringOk returns a tuple with the ForceString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1HelmParameter) GetForceStringOk() (*bool, bool) {
	if o == nil || o.ForceString == nil {
		return nil, false
	}
	return o.ForceString, true
}

// HasForceString returns a boolean if a field has been set.
func (o *V1alpha1HelmParameter) HasForceString() bool {
	if o != nil && o.ForceString != nil {
		return true
	}

	return false
}

// SetForceString gets a reference to the given bool and assigns it to the ForceString field.
func (o *V1alpha1HelmParameter) SetForceString(v bool) {
	o.ForceString = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1alpha1HelmParameter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1HelmParameter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1alpha1HelmParameter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1alpha1HelmParameter) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *V1alpha1HelmParameter) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1HelmParameter) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *V1alpha1HelmParameter) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *V1alpha1HelmParameter) SetValue(v string) {
	o.Value = &v
}

func (o V1alpha1HelmParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ForceString != nil {
		toSerialize["forceString"] = o.ForceString
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1HelmParameter struct {
	value *V1alpha1HelmParameter
	isSet bool
}

func (v NullableV1alpha1HelmParameter) Get() *V1alpha1HelmParameter {
	return v.value
}

func (v *NullableV1alpha1HelmParameter) Set(val *V1alpha1HelmParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1HelmParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1HelmParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1HelmParameter(val *V1alpha1HelmParameter) *NullableV1alpha1HelmParameter {
	return &NullableV1alpha1HelmParameter{value: val, isSet: true}
}

func (v NullableV1alpha1HelmParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1HelmParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
