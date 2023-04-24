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

// V1alpha1ApplicationSetRolloutStep struct for V1alpha1ApplicationSetRolloutStep
type V1alpha1ApplicationSetRolloutStep struct {
	MatchExpressions *[]V1alpha1ApplicationMatchExpression `json:"matchExpressions,omitempty"`
	MaxUpdate *IntstrIntOrString `json:"maxUpdate,omitempty"`
}

// NewV1alpha1ApplicationSetRolloutStep instantiates a new V1alpha1ApplicationSetRolloutStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ApplicationSetRolloutStep() *V1alpha1ApplicationSetRolloutStep {
	this := V1alpha1ApplicationSetRolloutStep{}
	return &this
}

// NewV1alpha1ApplicationSetRolloutStepWithDefaults instantiates a new V1alpha1ApplicationSetRolloutStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ApplicationSetRolloutStepWithDefaults() *V1alpha1ApplicationSetRolloutStep {
	this := V1alpha1ApplicationSetRolloutStep{}
	return &this
}

// GetMatchExpressions returns the MatchExpressions field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetRolloutStep) GetMatchExpressions() []V1alpha1ApplicationMatchExpression {
	if o == nil || o.MatchExpressions == nil {
		var ret []V1alpha1ApplicationMatchExpression
		return ret
	}
	return *o.MatchExpressions
}

// GetMatchExpressionsOk returns a tuple with the MatchExpressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetRolloutStep) GetMatchExpressionsOk() (*[]V1alpha1ApplicationMatchExpression, bool) {
	if o == nil || o.MatchExpressions == nil {
		return nil, false
	}
	return o.MatchExpressions, true
}

// HasMatchExpressions returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetRolloutStep) HasMatchExpressions() bool {
	if o != nil && o.MatchExpressions != nil {
		return true
	}

	return false
}

// SetMatchExpressions gets a reference to the given []V1alpha1ApplicationMatchExpression and assigns it to the MatchExpressions field.
func (o *V1alpha1ApplicationSetRolloutStep) SetMatchExpressions(v []V1alpha1ApplicationMatchExpression) {
	o.MatchExpressions = &v
}

// GetMaxUpdate returns the MaxUpdate field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetRolloutStep) GetMaxUpdate() IntstrIntOrString {
	if o == nil || o.MaxUpdate == nil {
		var ret IntstrIntOrString
		return ret
	}
	return *o.MaxUpdate
}

// GetMaxUpdateOk returns a tuple with the MaxUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetRolloutStep) GetMaxUpdateOk() (*IntstrIntOrString, bool) {
	if o == nil || o.MaxUpdate == nil {
		return nil, false
	}
	return o.MaxUpdate, true
}

// HasMaxUpdate returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetRolloutStep) HasMaxUpdate() bool {
	if o != nil && o.MaxUpdate != nil {
		return true
	}

	return false
}

// SetMaxUpdate gets a reference to the given IntstrIntOrString and assigns it to the MaxUpdate field.
func (o *V1alpha1ApplicationSetRolloutStep) SetMaxUpdate(v IntstrIntOrString) {
	o.MaxUpdate = &v
}

func (o V1alpha1ApplicationSetRolloutStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MatchExpressions != nil {
		toSerialize["matchExpressions"] = o.MatchExpressions
	}
	if o.MaxUpdate != nil {
		toSerialize["maxUpdate"] = o.MaxUpdate
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1ApplicationSetRolloutStep struct {
	value *V1alpha1ApplicationSetRolloutStep
	isSet bool
}

func (v NullableV1alpha1ApplicationSetRolloutStep) Get() *V1alpha1ApplicationSetRolloutStep {
	return v.value
}

func (v *NullableV1alpha1ApplicationSetRolloutStep) Set(val *V1alpha1ApplicationSetRolloutStep) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ApplicationSetRolloutStep) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ApplicationSetRolloutStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ApplicationSetRolloutStep(val *V1alpha1ApplicationSetRolloutStep) *NullableV1alpha1ApplicationSetRolloutStep {
	return &NullableV1alpha1ApplicationSetRolloutStep{value: val, isSet: true}
}

func (v NullableV1alpha1ApplicationSetRolloutStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ApplicationSetRolloutStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


