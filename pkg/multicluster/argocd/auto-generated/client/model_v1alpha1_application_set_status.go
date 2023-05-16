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

// V1alpha1ApplicationSetStatus struct for V1alpha1ApplicationSetStatus
type V1alpha1ApplicationSetStatus struct {
	ApplicationStatus *[]V1alpha1ApplicationSetApplicationStatus `json:"applicationStatus,omitempty"`
	Conditions        *[]V1alpha1ApplicationSetCondition         `json:"conditions,omitempty"`
}

// NewV1alpha1ApplicationSetStatus instantiates a new V1alpha1ApplicationSetStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ApplicationSetStatus() *V1alpha1ApplicationSetStatus {
	this := V1alpha1ApplicationSetStatus{}
	return &this
}

// NewV1alpha1ApplicationSetStatusWithDefaults instantiates a new V1alpha1ApplicationSetStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ApplicationSetStatusWithDefaults() *V1alpha1ApplicationSetStatus {
	this := V1alpha1ApplicationSetStatus{}
	return &this
}

// GetApplicationStatus returns the ApplicationStatus field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetStatus) GetApplicationStatus() []V1alpha1ApplicationSetApplicationStatus {
	if o == nil || o.ApplicationStatus == nil {
		var ret []V1alpha1ApplicationSetApplicationStatus
		return ret
	}
	return *o.ApplicationStatus
}

// GetApplicationStatusOk returns a tuple with the ApplicationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetStatus) GetApplicationStatusOk() (*[]V1alpha1ApplicationSetApplicationStatus, bool) {
	if o == nil || o.ApplicationStatus == nil {
		return nil, false
	}
	return o.ApplicationStatus, true
}

// HasApplicationStatus returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetStatus) HasApplicationStatus() bool {
	if o != nil && o.ApplicationStatus != nil {
		return true
	}

	return false
}

// SetApplicationStatus gets a reference to the given []V1alpha1ApplicationSetApplicationStatus and assigns it to the ApplicationStatus field.
func (o *V1alpha1ApplicationSetStatus) SetApplicationStatus(v []V1alpha1ApplicationSetApplicationStatus) {
	o.ApplicationStatus = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetStatus) GetConditions() []V1alpha1ApplicationSetCondition {
	if o == nil || o.Conditions == nil {
		var ret []V1alpha1ApplicationSetCondition
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetStatus) GetConditionsOk() (*[]V1alpha1ApplicationSetCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetStatus) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []V1alpha1ApplicationSetCondition and assigns it to the Conditions field.
func (o *V1alpha1ApplicationSetStatus) SetConditions(v []V1alpha1ApplicationSetCondition) {
	o.Conditions = &v
}

func (o V1alpha1ApplicationSetStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationStatus != nil {
		toSerialize["applicationStatus"] = o.ApplicationStatus
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1ApplicationSetStatus struct {
	value *V1alpha1ApplicationSetStatus
	isSet bool
}

func (v NullableV1alpha1ApplicationSetStatus) Get() *V1alpha1ApplicationSetStatus {
	return v.value
}

func (v *NullableV1alpha1ApplicationSetStatus) Set(val *V1alpha1ApplicationSetStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ApplicationSetStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ApplicationSetStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ApplicationSetStatus(val *V1alpha1ApplicationSetStatus) *NullableV1alpha1ApplicationSetStatus {
	return &NullableV1alpha1ApplicationSetStatus{value: val, isSet: true}
}

func (v NullableV1alpha1ApplicationSetStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ApplicationSetStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}