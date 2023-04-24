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

// V1alpha1ListGenerator struct for V1alpha1ListGenerator
type V1alpha1ListGenerator struct {
	Elements *[]V1JSON `json:"elements,omitempty"`
	Template *V1alpha1ApplicationSetTemplate `json:"template,omitempty"`
}

// NewV1alpha1ListGenerator instantiates a new V1alpha1ListGenerator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ListGenerator() *V1alpha1ListGenerator {
	this := V1alpha1ListGenerator{}
	return &this
}

// NewV1alpha1ListGeneratorWithDefaults instantiates a new V1alpha1ListGenerator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ListGeneratorWithDefaults() *V1alpha1ListGenerator {
	this := V1alpha1ListGenerator{}
	return &this
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *V1alpha1ListGenerator) GetElements() []V1JSON {
	if o == nil || o.Elements == nil {
		var ret []V1JSON
		return ret
	}
	return *o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ListGenerator) GetElementsOk() (*[]V1JSON, bool) {
	if o == nil || o.Elements == nil {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *V1alpha1ListGenerator) HasElements() bool {
	if o != nil && o.Elements != nil {
		return true
	}

	return false
}

// SetElements gets a reference to the given []V1JSON and assigns it to the Elements field.
func (o *V1alpha1ListGenerator) SetElements(v []V1JSON) {
	o.Elements = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *V1alpha1ListGenerator) GetTemplate() V1alpha1ApplicationSetTemplate {
	if o == nil || o.Template == nil {
		var ret V1alpha1ApplicationSetTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ListGenerator) GetTemplateOk() (*V1alpha1ApplicationSetTemplate, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *V1alpha1ListGenerator) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given V1alpha1ApplicationSetTemplate and assigns it to the Template field.
func (o *V1alpha1ListGenerator) SetTemplate(v V1alpha1ApplicationSetTemplate) {
	o.Template = &v
}

func (o V1alpha1ListGenerator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Elements != nil {
		toSerialize["elements"] = o.Elements
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1ListGenerator struct {
	value *V1alpha1ListGenerator
	isSet bool
}

func (v NullableV1alpha1ListGenerator) Get() *V1alpha1ListGenerator {
	return v.value
}

func (v *NullableV1alpha1ListGenerator) Set(val *V1alpha1ListGenerator) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ListGenerator) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ListGenerator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ListGenerator(val *V1alpha1ListGenerator) *NullableV1alpha1ListGenerator {
	return &NullableV1alpha1ListGenerator{value: val, isSet: true}
}

func (v NullableV1alpha1ListGenerator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ListGenerator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


