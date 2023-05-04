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

// V1alpha1ApplicationSetTemplate struct for V1alpha1ApplicationSetTemplate
type V1alpha1ApplicationSetTemplate struct {
	Metadata *V1alpha1ApplicationSetTemplateMeta `json:"metadata,omitempty"`
	Spec     *V1alpha1ApplicationSpec            `json:"spec,omitempty"`
}

// NewV1alpha1ApplicationSetTemplate instantiates a new V1alpha1ApplicationSetTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ApplicationSetTemplate() *V1alpha1ApplicationSetTemplate {
	this := V1alpha1ApplicationSetTemplate{}
	return &this
}

// NewV1alpha1ApplicationSetTemplateWithDefaults instantiates a new V1alpha1ApplicationSetTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ApplicationSetTemplateWithDefaults() *V1alpha1ApplicationSetTemplate {
	this := V1alpha1ApplicationSetTemplate{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetTemplate) GetMetadata() V1alpha1ApplicationSetTemplateMeta {
	if o == nil || o.Metadata == nil {
		var ret V1alpha1ApplicationSetTemplateMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetTemplate) GetMetadataOk() (*V1alpha1ApplicationSetTemplateMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetTemplate) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1alpha1ApplicationSetTemplateMeta and assigns it to the Metadata field.
func (o *V1alpha1ApplicationSetTemplate) SetMetadata(v V1alpha1ApplicationSetTemplateMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetTemplate) GetSpec() V1alpha1ApplicationSpec {
	if o == nil || o.Spec == nil {
		var ret V1alpha1ApplicationSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetTemplate) GetSpecOk() (*V1alpha1ApplicationSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetTemplate) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given V1alpha1ApplicationSpec and assigns it to the Spec field.
func (o *V1alpha1ApplicationSetTemplate) SetSpec(v V1alpha1ApplicationSpec) {
	o.Spec = &v
}

func (o V1alpha1ApplicationSetTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1ApplicationSetTemplate struct {
	value *V1alpha1ApplicationSetTemplate
	isSet bool
}

func (v NullableV1alpha1ApplicationSetTemplate) Get() *V1alpha1ApplicationSetTemplate {
	return v.value
}

func (v *NullableV1alpha1ApplicationSetTemplate) Set(val *V1alpha1ApplicationSetTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ApplicationSetTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ApplicationSetTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ApplicationSetTemplate(val *V1alpha1ApplicationSetTemplate) *NullableV1alpha1ApplicationSetTemplate {
	return &NullableV1alpha1ApplicationSetTemplate{value: val, isSet: true}
}

func (v NullableV1alpha1ApplicationSetTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ApplicationSetTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
