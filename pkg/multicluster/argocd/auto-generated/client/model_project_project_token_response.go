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

// ProjectProjectTokenResponse ProjectTokenResponse wraps the created token or returns an empty string if deleted.
type ProjectProjectTokenResponse struct {
	Token *string `json:"token,omitempty"`
}

// NewProjectProjectTokenResponse instantiates a new ProjectProjectTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectProjectTokenResponse() *ProjectProjectTokenResponse {
	this := ProjectProjectTokenResponse{}
	return &this
}

// NewProjectProjectTokenResponseWithDefaults instantiates a new ProjectProjectTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectProjectTokenResponseWithDefaults() *ProjectProjectTokenResponse {
	this := ProjectProjectTokenResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ProjectProjectTokenResponse) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectTokenResponse) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ProjectProjectTokenResponse) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ProjectProjectTokenResponse) SetToken(v string) {
	o.Token = &v
}

func (o ProjectProjectTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableProjectProjectTokenResponse struct {
	value *ProjectProjectTokenResponse
	isSet bool
}

func (v NullableProjectProjectTokenResponse) Get() *ProjectProjectTokenResponse {
	return v.value
}

func (v *NullableProjectProjectTokenResponse) Set(val *ProjectProjectTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectProjectTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectProjectTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectProjectTokenResponse(val *ProjectProjectTokenResponse) *NullableProjectProjectTokenResponse {
	return &NullableProjectProjectTokenResponse{value: val, isSet: true}
}

func (v NullableProjectProjectTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectProjectTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


