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

// ClusterSettingsPluginsResponse struct for ClusterSettingsPluginsResponse
type ClusterSettingsPluginsResponse struct {
	Plugins *[]ClusterPlugin `json:"plugins,omitempty"`
}

// NewClusterSettingsPluginsResponse instantiates a new ClusterSettingsPluginsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterSettingsPluginsResponse() *ClusterSettingsPluginsResponse {
	this := ClusterSettingsPluginsResponse{}
	return &this
}

// NewClusterSettingsPluginsResponseWithDefaults instantiates a new ClusterSettingsPluginsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterSettingsPluginsResponseWithDefaults() *ClusterSettingsPluginsResponse {
	this := ClusterSettingsPluginsResponse{}
	return &this
}

// GetPlugins returns the Plugins field value if set, zero value otherwise.
func (o *ClusterSettingsPluginsResponse) GetPlugins() []ClusterPlugin {
	if o == nil || o.Plugins == nil {
		var ret []ClusterPlugin
		return ret
	}
	return *o.Plugins
}

// GetPluginsOk returns a tuple with the Plugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSettingsPluginsResponse) GetPluginsOk() (*[]ClusterPlugin, bool) {
	if o == nil || o.Plugins == nil {
		return nil, false
	}
	return o.Plugins, true
}

// HasPlugins returns a boolean if a field has been set.
func (o *ClusterSettingsPluginsResponse) HasPlugins() bool {
	if o != nil && o.Plugins != nil {
		return true
	}

	return false
}

// SetPlugins gets a reference to the given []ClusterPlugin and assigns it to the Plugins field.
func (o *ClusterSettingsPluginsResponse) SetPlugins(v []ClusterPlugin) {
	o.Plugins = &v
}

func (o ClusterSettingsPluginsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Plugins != nil {
		toSerialize["plugins"] = o.Plugins
	}
	return json.Marshal(toSerialize)
}

type NullableClusterSettingsPluginsResponse struct {
	value *ClusterSettingsPluginsResponse
	isSet bool
}

func (v NullableClusterSettingsPluginsResponse) Get() *ClusterSettingsPluginsResponse {
	return v.value
}

func (v *NullableClusterSettingsPluginsResponse) Set(val *ClusterSettingsPluginsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterSettingsPluginsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterSettingsPluginsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterSettingsPluginsResponse(val *ClusterSettingsPluginsResponse) *NullableClusterSettingsPluginsResponse {
	return &NullableClusterSettingsPluginsResponse{value: val, isSet: true}
}

func (v NullableClusterSettingsPluginsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterSettingsPluginsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}