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

// V1alpha1BasicAuthBitbucketServer BasicAuthBitbucketServer defines the username/(password or personal access token) for Basic auth.
type V1alpha1BasicAuthBitbucketServer struct {
	PasswordRef *V1alpha1SecretRef `json:"passwordRef,omitempty"`
	Username    *string            `json:"username,omitempty"`
}

// NewV1alpha1BasicAuthBitbucketServer instantiates a new V1alpha1BasicAuthBitbucketServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1BasicAuthBitbucketServer() *V1alpha1BasicAuthBitbucketServer {
	this := V1alpha1BasicAuthBitbucketServer{}
	return &this
}

// NewV1alpha1BasicAuthBitbucketServerWithDefaults instantiates a new V1alpha1BasicAuthBitbucketServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1BasicAuthBitbucketServerWithDefaults() *V1alpha1BasicAuthBitbucketServer {
	this := V1alpha1BasicAuthBitbucketServer{}
	return &this
}

// GetPasswordRef returns the PasswordRef field value if set, zero value otherwise.
func (o *V1alpha1BasicAuthBitbucketServer) GetPasswordRef() V1alpha1SecretRef {
	if o == nil || o.PasswordRef == nil {
		var ret V1alpha1SecretRef
		return ret
	}
	return *o.PasswordRef
}

// GetPasswordRefOk returns a tuple with the PasswordRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1BasicAuthBitbucketServer) GetPasswordRefOk() (*V1alpha1SecretRef, bool) {
	if o == nil || o.PasswordRef == nil {
		return nil, false
	}
	return o.PasswordRef, true
}

// HasPasswordRef returns a boolean if a field has been set.
func (o *V1alpha1BasicAuthBitbucketServer) HasPasswordRef() bool {
	if o != nil && o.PasswordRef != nil {
		return true
	}

	return false
}

// SetPasswordRef gets a reference to the given V1alpha1SecretRef and assigns it to the PasswordRef field.
func (o *V1alpha1BasicAuthBitbucketServer) SetPasswordRef(v V1alpha1SecretRef) {
	o.PasswordRef = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *V1alpha1BasicAuthBitbucketServer) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1BasicAuthBitbucketServer) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *V1alpha1BasicAuthBitbucketServer) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *V1alpha1BasicAuthBitbucketServer) SetUsername(v string) {
	o.Username = &v
}

func (o V1alpha1BasicAuthBitbucketServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PasswordRef != nil {
		toSerialize["passwordRef"] = o.PasswordRef
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1BasicAuthBitbucketServer struct {
	value *V1alpha1BasicAuthBitbucketServer
	isSet bool
}

func (v NullableV1alpha1BasicAuthBitbucketServer) Get() *V1alpha1BasicAuthBitbucketServer {
	return v.value
}

func (v *NullableV1alpha1BasicAuthBitbucketServer) Set(val *V1alpha1BasicAuthBitbucketServer) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1BasicAuthBitbucketServer) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1BasicAuthBitbucketServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1BasicAuthBitbucketServer(val *V1alpha1BasicAuthBitbucketServer) *NullableV1alpha1BasicAuthBitbucketServer {
	return &NullableV1alpha1BasicAuthBitbucketServer{value: val, isSet: true}
}

func (v NullableV1alpha1BasicAuthBitbucketServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1BasicAuthBitbucketServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}