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

// V1alpha1ResourceDiff struct for V1alpha1ResourceDiff
type V1alpha1ResourceDiff struct {
	Diff                *string `json:"diff,omitempty"`
	Group               *string `json:"group,omitempty"`
	Hook                *bool   `json:"hook,omitempty"`
	Kind                *string `json:"kind,omitempty"`
	LiveState           *string `json:"liveState,omitempty"`
	Modified            *bool   `json:"modified,omitempty"`
	Name                *string `json:"name,omitempty"`
	Namespace           *string `json:"namespace,omitempty"`
	NormalizedLiveState *string `json:"normalizedLiveState,omitempty"`
	PredictedLiveState  *string `json:"predictedLiveState,omitempty"`
	ResourceVersion     *string `json:"resourceVersion,omitempty"`
	TargetState         *string `json:"targetState,omitempty"`
}

// NewV1alpha1ResourceDiff instantiates a new V1alpha1ResourceDiff object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ResourceDiff() *V1alpha1ResourceDiff {
	this := V1alpha1ResourceDiff{}
	return &this
}

// NewV1alpha1ResourceDiffWithDefaults instantiates a new V1alpha1ResourceDiff object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ResourceDiffWithDefaults() *V1alpha1ResourceDiff {
	this := V1alpha1ResourceDiff{}
	return &this
}

// GetDiff returns the Diff field value if set, zero value otherwise.
func (o *V1alpha1ResourceDiff) GetDiff() string {
	if o == nil || o.Diff == nil {
		var ret string
		return ret
	}
	return *o.Diff
}

// GetDiffOk returns a tuple with the Diff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceDiff) GetDiffOk() (*string, bool) {
	if o == nil || o.Diff == nil {
		return nil, false
	}
	return o.Diff, true
}

// HasDiff returns a boolean if a field has been set.
func (o *V1alpha1ResourceDiff) HasDiff() bool {
	if o != nil && o.Diff != nil {
		return true
	}

	return false
}

// SetDiff gets a reference to the given string and assigns it to the Diff field.
func (o *V1alpha1ResourceDiff) SetDiff(v string) {
	o.Diff = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *V1alpha1ResourceDiff) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceDiff) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *V1alpha1ResourceDiff) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *V1alpha1ResourceDiff) SetGroup(v string) {
	o.Group = &v
}

// GetHook returns the Hook field value if set, zero value otherwise.
func (o *V1alpha1ResourceDiff) GetHook() bool {
	if o == nil || o.Hook == nil {
		var ret bool
		return ret
	}
	return *o.Hook
}

// GetHookOk returns a tuple with the Hook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceDiff) GetHookOk() (*bool, bool) {
	if o == nil || o.Hook == nil {
		return nil, false
	}
	return o.Hook, true
}

// HasHook returns a boolean if a field has been set.
func (o *V1alpha1ResourceDiff) HasHook() bool {
	if o != nil && o.Hook != nil {
		return true
	}

	return false
}

// SetHook gets a reference to the given bool and assigns it to the Hook field.
func (o *V1alpha1ResourceDiff) SetHook(v bool) {
	o.Hook = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1alpha1ResourceDiff) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceDiff) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1alpha1ResourceDiff) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1alpha1ResourceDiff) SetKind(v string) {
	o.Kind = &v
}

// GetLiveState returns the LiveState field value if set, zero value otherwise.
func (o *V1alpha1ResourceDiff) GetLiveState() string {
	if o == nil || o.LiveState == nil {
		var ret string
		return ret
	}
	return *o.LiveState
}

// GetLiveStateOk returns a tuple with the LiveState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceDiff) GetLiveStateOk() (*string, bool) {
	if o == nil || o.LiveState == nil {
		return nil, false
	}
	return o.LiveState, true
}

// HasLiveState returns a boolean if a field has been set.
func (o *V1alpha1ResourceDiff) HasLiveState() bool {
	if o != nil && o.LiveState != nil {
		return true
	}

	return false
}

// SetLiveState gets a reference to the given string and assigns it to the LiveState field.
func (o *V1alpha1ResourceDiff) SetLiveState(v string) {
	o.LiveState = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *V1alpha1ResourceDiff) GetModified() bool {
	if o == nil || o.Modified == nil {
		var ret bool
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceDiff) GetModifiedOk() (*bool, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *V1alpha1ResourceDiff) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given bool and assigns it to the Modified field.
func (o *V1alpha1ResourceDiff) SetModified(v bool) {
	o.Modified = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1alpha1ResourceDiff) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceDiff) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1alpha1ResourceDiff) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1alpha1ResourceDiff) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *V1alpha1ResourceDiff) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceDiff) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *V1alpha1ResourceDiff) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *V1alpha1ResourceDiff) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNormalizedLiveState returns the NormalizedLiveState field value if set, zero value otherwise.
func (o *V1alpha1ResourceDiff) GetNormalizedLiveState() string {
	if o == nil || o.NormalizedLiveState == nil {
		var ret string
		return ret
	}
	return *o.NormalizedLiveState
}

// GetNormalizedLiveStateOk returns a tuple with the NormalizedLiveState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceDiff) GetNormalizedLiveStateOk() (*string, bool) {
	if o == nil || o.NormalizedLiveState == nil {
		return nil, false
	}
	return o.NormalizedLiveState, true
}

// HasNormalizedLiveState returns a boolean if a field has been set.
func (o *V1alpha1ResourceDiff) HasNormalizedLiveState() bool {
	if o != nil && o.NormalizedLiveState != nil {
		return true
	}

	return false
}

// SetNormalizedLiveState gets a reference to the given string and assigns it to the NormalizedLiveState field.
func (o *V1alpha1ResourceDiff) SetNormalizedLiveState(v string) {
	o.NormalizedLiveState = &v
}

// GetPredictedLiveState returns the PredictedLiveState field value if set, zero value otherwise.
func (o *V1alpha1ResourceDiff) GetPredictedLiveState() string {
	if o == nil || o.PredictedLiveState == nil {
		var ret string
		return ret
	}
	return *o.PredictedLiveState
}

// GetPredictedLiveStateOk returns a tuple with the PredictedLiveState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceDiff) GetPredictedLiveStateOk() (*string, bool) {
	if o == nil || o.PredictedLiveState == nil {
		return nil, false
	}
	return o.PredictedLiveState, true
}

// HasPredictedLiveState returns a boolean if a field has been set.
func (o *V1alpha1ResourceDiff) HasPredictedLiveState() bool {
	if o != nil && o.PredictedLiveState != nil {
		return true
	}

	return false
}

// SetPredictedLiveState gets a reference to the given string and assigns it to the PredictedLiveState field.
func (o *V1alpha1ResourceDiff) SetPredictedLiveState(v string) {
	o.PredictedLiveState = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *V1alpha1ResourceDiff) GetResourceVersion() string {
	if o == nil || o.ResourceVersion == nil {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceDiff) GetResourceVersionOk() (*string, bool) {
	if o == nil || o.ResourceVersion == nil {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *V1alpha1ResourceDiff) HasResourceVersion() bool {
	if o != nil && o.ResourceVersion != nil {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *V1alpha1ResourceDiff) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}

// GetTargetState returns the TargetState field value if set, zero value otherwise.
func (o *V1alpha1ResourceDiff) GetTargetState() string {
	if o == nil || o.TargetState == nil {
		var ret string
		return ret
	}
	return *o.TargetState
}

// GetTargetStateOk returns a tuple with the TargetState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceDiff) GetTargetStateOk() (*string, bool) {
	if o == nil || o.TargetState == nil {
		return nil, false
	}
	return o.TargetState, true
}

// HasTargetState returns a boolean if a field has been set.
func (o *V1alpha1ResourceDiff) HasTargetState() bool {
	if o != nil && o.TargetState != nil {
		return true
	}

	return false
}

// SetTargetState gets a reference to the given string and assigns it to the TargetState field.
func (o *V1alpha1ResourceDiff) SetTargetState(v string) {
	o.TargetState = &v
}

func (o V1alpha1ResourceDiff) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Diff != nil {
		toSerialize["diff"] = o.Diff
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Hook != nil {
		toSerialize["hook"] = o.Hook
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.LiveState != nil {
		toSerialize["liveState"] = o.LiveState
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NormalizedLiveState != nil {
		toSerialize["normalizedLiveState"] = o.NormalizedLiveState
	}
	if o.PredictedLiveState != nil {
		toSerialize["predictedLiveState"] = o.PredictedLiveState
	}
	if o.ResourceVersion != nil {
		toSerialize["resourceVersion"] = o.ResourceVersion
	}
	if o.TargetState != nil {
		toSerialize["targetState"] = o.TargetState
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1ResourceDiff struct {
	value *V1alpha1ResourceDiff
	isSet bool
}

func (v NullableV1alpha1ResourceDiff) Get() *V1alpha1ResourceDiff {
	return v.value
}

func (v *NullableV1alpha1ResourceDiff) Set(val *V1alpha1ResourceDiff) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ResourceDiff) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ResourceDiff) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ResourceDiff(val *V1alpha1ResourceDiff) *NullableV1alpha1ResourceDiff {
	return &NullableV1alpha1ResourceDiff{value: val, isSet: true}
}

func (v NullableV1alpha1ResourceDiff) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ResourceDiff) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}