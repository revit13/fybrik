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

// V1alpha1SyncOperationResult struct for V1alpha1SyncOperationResult
type V1alpha1SyncOperationResult struct {
	Resources *[]V1alpha1ResourceResult    `json:"resources,omitempty"`
	Revision  *string                      `json:"revision,omitempty"`
	Revisions *[]string                    `json:"revisions,omitempty"`
	Source    *V1alpha1ApplicationSource   `json:"source,omitempty"`
	Sources   *[]V1alpha1ApplicationSource `json:"sources,omitempty"`
}

// NewV1alpha1SyncOperationResult instantiates a new V1alpha1SyncOperationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1SyncOperationResult() *V1alpha1SyncOperationResult {
	this := V1alpha1SyncOperationResult{}
	return &this
}

// NewV1alpha1SyncOperationResultWithDefaults instantiates a new V1alpha1SyncOperationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1SyncOperationResultWithDefaults() *V1alpha1SyncOperationResult {
	this := V1alpha1SyncOperationResult{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *V1alpha1SyncOperationResult) GetResources() []V1alpha1ResourceResult {
	if o == nil || o.Resources == nil {
		var ret []V1alpha1ResourceResult
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperationResult) GetResourcesOk() (*[]V1alpha1ResourceResult, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *V1alpha1SyncOperationResult) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []V1alpha1ResourceResult and assigns it to the Resources field.
func (o *V1alpha1SyncOperationResult) SetResources(v []V1alpha1ResourceResult) {
	o.Resources = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *V1alpha1SyncOperationResult) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperationResult) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *V1alpha1SyncOperationResult) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *V1alpha1SyncOperationResult) SetRevision(v string) {
	o.Revision = &v
}

// GetRevisions returns the Revisions field value if set, zero value otherwise.
func (o *V1alpha1SyncOperationResult) GetRevisions() []string {
	if o == nil || o.Revisions == nil {
		var ret []string
		return ret
	}
	return *o.Revisions
}

// GetRevisionsOk returns a tuple with the Revisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperationResult) GetRevisionsOk() (*[]string, bool) {
	if o == nil || o.Revisions == nil {
		return nil, false
	}
	return o.Revisions, true
}

// HasRevisions returns a boolean if a field has been set.
func (o *V1alpha1SyncOperationResult) HasRevisions() bool {
	if o != nil && o.Revisions != nil {
		return true
	}

	return false
}

// SetRevisions gets a reference to the given []string and assigns it to the Revisions field.
func (o *V1alpha1SyncOperationResult) SetRevisions(v []string) {
	o.Revisions = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *V1alpha1SyncOperationResult) GetSource() V1alpha1ApplicationSource {
	if o == nil || o.Source == nil {
		var ret V1alpha1ApplicationSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperationResult) GetSourceOk() (*V1alpha1ApplicationSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *V1alpha1SyncOperationResult) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given V1alpha1ApplicationSource and assigns it to the Source field.
func (o *V1alpha1SyncOperationResult) SetSource(v V1alpha1ApplicationSource) {
	o.Source = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *V1alpha1SyncOperationResult) GetSources() []V1alpha1ApplicationSource {
	if o == nil || o.Sources == nil {
		var ret []V1alpha1ApplicationSource
		return ret
	}
	return *o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperationResult) GetSourcesOk() (*[]V1alpha1ApplicationSource, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *V1alpha1SyncOperationResult) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []V1alpha1ApplicationSource and assigns it to the Sources field.
func (o *V1alpha1SyncOperationResult) SetSources(v []V1alpha1ApplicationSource) {
	o.Sources = &v
}

func (o V1alpha1SyncOperationResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.Revisions != nil {
		toSerialize["revisions"] = o.Revisions
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1SyncOperationResult struct {
	value *V1alpha1SyncOperationResult
	isSet bool
}

func (v NullableV1alpha1SyncOperationResult) Get() *V1alpha1SyncOperationResult {
	return v.value
}

func (v *NullableV1alpha1SyncOperationResult) Set(val *V1alpha1SyncOperationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1SyncOperationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1SyncOperationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1SyncOperationResult(val *V1alpha1SyncOperationResult) *NullableV1alpha1SyncOperationResult {
	return &NullableV1alpha1SyncOperationResult{value: val, isSet: true}
}

func (v NullableV1alpha1SyncOperationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1SyncOperationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}