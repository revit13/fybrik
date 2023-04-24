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

// V1alpha1JWTTokens struct for V1alpha1JWTTokens
type V1alpha1JWTTokens struct {
	Items *[]V1alpha1JWTToken `json:"items,omitempty"`
}

// NewV1alpha1JWTTokens instantiates a new V1alpha1JWTTokens object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1JWTTokens() *V1alpha1JWTTokens {
	this := V1alpha1JWTTokens{}
	return &this
}

// NewV1alpha1JWTTokensWithDefaults instantiates a new V1alpha1JWTTokens object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1JWTTokensWithDefaults() *V1alpha1JWTTokens {
	this := V1alpha1JWTTokens{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *V1alpha1JWTTokens) GetItems() []V1alpha1JWTToken {
	if o == nil || o.Items == nil {
		var ret []V1alpha1JWTToken
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1JWTTokens) GetItemsOk() (*[]V1alpha1JWTToken, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *V1alpha1JWTTokens) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []V1alpha1JWTToken and assigns it to the Items field.
func (o *V1alpha1JWTTokens) SetItems(v []V1alpha1JWTToken) {
	o.Items = &v
}

func (o V1alpha1JWTTokens) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1JWTTokens struct {
	value *V1alpha1JWTTokens
	isSet bool
}

func (v NullableV1alpha1JWTTokens) Get() *V1alpha1JWTTokens {
	return v.value
}

func (v *NullableV1alpha1JWTTokens) Set(val *V1alpha1JWTTokens) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1JWTTokens) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1JWTTokens) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1JWTTokens(val *V1alpha1JWTTokens) *NullableV1alpha1JWTTokens {
	return &NullableV1alpha1JWTTokens{value: val, isSet: true}
}

func (v NullableV1alpha1JWTTokens) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1JWTTokens) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


