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

// ApplicationApplicationResourceResponse struct for ApplicationApplicationResourceResponse
type ApplicationApplicationResourceResponse struct {
	Manifest *string `json:"manifest,omitempty"`
}

// NewApplicationApplicationResourceResponse instantiates a new ApplicationApplicationResourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationApplicationResourceResponse() *ApplicationApplicationResourceResponse {
	this := ApplicationApplicationResourceResponse{}
	return &this
}

// NewApplicationApplicationResourceResponseWithDefaults instantiates a new ApplicationApplicationResourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationApplicationResourceResponseWithDefaults() *ApplicationApplicationResourceResponse {
	this := ApplicationApplicationResourceResponse{}
	return &this
}

// GetManifest returns the Manifest field value if set, zero value otherwise.
func (o *ApplicationApplicationResourceResponse) GetManifest() string {
	if o == nil || o.Manifest == nil {
		var ret string
		return ret
	}
	return *o.Manifest
}

// GetManifestOk returns a tuple with the Manifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApplicationResourceResponse) GetManifestOk() (*string, bool) {
	if o == nil || o.Manifest == nil {
		return nil, false
	}
	return o.Manifest, true
}

// HasManifest returns a boolean if a field has been set.
func (o *ApplicationApplicationResourceResponse) HasManifest() bool {
	if o != nil && o.Manifest != nil {
		return true
	}

	return false
}

// SetManifest gets a reference to the given string and assigns it to the Manifest field.
func (o *ApplicationApplicationResourceResponse) SetManifest(v string) {
	o.Manifest = &v
}

func (o ApplicationApplicationResourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Manifest != nil {
		toSerialize["manifest"] = o.Manifest
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationApplicationResourceResponse struct {
	value *ApplicationApplicationResourceResponse
	isSet bool
}

func (v NullableApplicationApplicationResourceResponse) Get() *ApplicationApplicationResourceResponse {
	return v.value
}

func (v *NullableApplicationApplicationResourceResponse) Set(val *ApplicationApplicationResourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationApplicationResourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationApplicationResourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationApplicationResourceResponse(val *ApplicationApplicationResourceResponse) *NullableApplicationApplicationResourceResponse {
	return &NullableApplicationApplicationResourceResponse{value: val, isSet: true}
}

func (v NullableApplicationApplicationResourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationApplicationResourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}