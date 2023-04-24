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

// V1alpha1HostInfo struct for V1alpha1HostInfo
type V1alpha1HostInfo struct {
	Name *string `json:"name,omitempty"`
	ResourcesInfo *[]V1alpha1HostResourceInfo `json:"resourcesInfo,omitempty"`
	SystemInfo *V1NodeSystemInfo `json:"systemInfo,omitempty"`
}

// NewV1alpha1HostInfo instantiates a new V1alpha1HostInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1HostInfo() *V1alpha1HostInfo {
	this := V1alpha1HostInfo{}
	return &this
}

// NewV1alpha1HostInfoWithDefaults instantiates a new V1alpha1HostInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1HostInfoWithDefaults() *V1alpha1HostInfo {
	this := V1alpha1HostInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1alpha1HostInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1HostInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1alpha1HostInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1alpha1HostInfo) SetName(v string) {
	o.Name = &v
}

// GetResourcesInfo returns the ResourcesInfo field value if set, zero value otherwise.
func (o *V1alpha1HostInfo) GetResourcesInfo() []V1alpha1HostResourceInfo {
	if o == nil || o.ResourcesInfo == nil {
		var ret []V1alpha1HostResourceInfo
		return ret
	}
	return *o.ResourcesInfo
}

// GetResourcesInfoOk returns a tuple with the ResourcesInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1HostInfo) GetResourcesInfoOk() (*[]V1alpha1HostResourceInfo, bool) {
	if o == nil || o.ResourcesInfo == nil {
		return nil, false
	}
	return o.ResourcesInfo, true
}

// HasResourcesInfo returns a boolean if a field has been set.
func (o *V1alpha1HostInfo) HasResourcesInfo() bool {
	if o != nil && o.ResourcesInfo != nil {
		return true
	}

	return false
}

// SetResourcesInfo gets a reference to the given []V1alpha1HostResourceInfo and assigns it to the ResourcesInfo field.
func (o *V1alpha1HostInfo) SetResourcesInfo(v []V1alpha1HostResourceInfo) {
	o.ResourcesInfo = &v
}

// GetSystemInfo returns the SystemInfo field value if set, zero value otherwise.
func (o *V1alpha1HostInfo) GetSystemInfo() V1NodeSystemInfo {
	if o == nil || o.SystemInfo == nil {
		var ret V1NodeSystemInfo
		return ret
	}
	return *o.SystemInfo
}

// GetSystemInfoOk returns a tuple with the SystemInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1HostInfo) GetSystemInfoOk() (*V1NodeSystemInfo, bool) {
	if o == nil || o.SystemInfo == nil {
		return nil, false
	}
	return o.SystemInfo, true
}

// HasSystemInfo returns a boolean if a field has been set.
func (o *V1alpha1HostInfo) HasSystemInfo() bool {
	if o != nil && o.SystemInfo != nil {
		return true
	}

	return false
}

// SetSystemInfo gets a reference to the given V1NodeSystemInfo and assigns it to the SystemInfo field.
func (o *V1alpha1HostInfo) SetSystemInfo(v V1NodeSystemInfo) {
	o.SystemInfo = &v
}

func (o V1alpha1HostInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ResourcesInfo != nil {
		toSerialize["resourcesInfo"] = o.ResourcesInfo
	}
	if o.SystemInfo != nil {
		toSerialize["systemInfo"] = o.SystemInfo
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1HostInfo struct {
	value *V1alpha1HostInfo
	isSet bool
}

func (v NullableV1alpha1HostInfo) Get() *V1alpha1HostInfo {
	return v.value
}

func (v *NullableV1alpha1HostInfo) Set(val *V1alpha1HostInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1HostInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1HostInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1HostInfo(val *V1alpha1HostInfo) *NullableV1alpha1HostInfo {
	return &NullableV1alpha1HostInfo{value: val, isSet: true}
}

func (v NullableV1alpha1HostInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1HostInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


