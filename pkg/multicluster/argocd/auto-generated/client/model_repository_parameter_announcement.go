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

// RepositoryParameterAnnouncement struct for RepositoryParameterAnnouncement
type RepositoryParameterAnnouncement struct {
	// array is the default value of the parameter if the parameter is an array.
	Array *[]string `json:"array,omitempty"`
	// collectionType is the type of value this parameter holds - either a single value (a string) or a collection (array or map). If collectionType is set, only the field with that type will be used. If collectionType is not set, `string` is the default. If collectionType is set to an invalid value, a validation error is thrown.
	CollectionType *string `json:"collectionType,omitempty"`
	// itemType determines the primitive data type represented by the parameter. Parameters are always encoded as strings, but this field lets them be interpreted as other primitive types.
	ItemType *string `json:"itemType,omitempty"`
	// map is the default value of the parameter if the parameter is a map.
	Map *map[string]string `json:"map,omitempty"`
	// name is the name identifying a parameter.
	Name *string `json:"name,omitempty"`
	// required defines if this given parameter is mandatory.
	Required *bool `json:"required,omitempty"`
	// string is the default value of the parameter if the parameter is a string.
	String *string `json:"string,omitempty"`
	// title is a human-readable text of the parameter name.
	Title *string `json:"title,omitempty"`
	// tooltip is a human-readable description of the parameter.
	Tooltip *string `json:"tooltip,omitempty"`
}

// NewRepositoryParameterAnnouncement instantiates a new RepositoryParameterAnnouncement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryParameterAnnouncement() *RepositoryParameterAnnouncement {
	this := RepositoryParameterAnnouncement{}
	return &this
}

// NewRepositoryParameterAnnouncementWithDefaults instantiates a new RepositoryParameterAnnouncement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryParameterAnnouncementWithDefaults() *RepositoryParameterAnnouncement {
	this := RepositoryParameterAnnouncement{}
	return &this
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *RepositoryParameterAnnouncement) GetArray() []string {
	if o == nil || o.Array == nil {
		var ret []string
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryParameterAnnouncement) GetArrayOk() (*[]string, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *RepositoryParameterAnnouncement) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given []string and assigns it to the Array field.
func (o *RepositoryParameterAnnouncement) SetArray(v []string) {
	o.Array = &v
}

// GetCollectionType returns the CollectionType field value if set, zero value otherwise.
func (o *RepositoryParameterAnnouncement) GetCollectionType() string {
	if o == nil || o.CollectionType == nil {
		var ret string
		return ret
	}
	return *o.CollectionType
}

// GetCollectionTypeOk returns a tuple with the CollectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryParameterAnnouncement) GetCollectionTypeOk() (*string, bool) {
	if o == nil || o.CollectionType == nil {
		return nil, false
	}
	return o.CollectionType, true
}

// HasCollectionType returns a boolean if a field has been set.
func (o *RepositoryParameterAnnouncement) HasCollectionType() bool {
	if o != nil && o.CollectionType != nil {
		return true
	}

	return false
}

// SetCollectionType gets a reference to the given string and assigns it to the CollectionType field.
func (o *RepositoryParameterAnnouncement) SetCollectionType(v string) {
	o.CollectionType = &v
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *RepositoryParameterAnnouncement) GetItemType() string {
	if o == nil || o.ItemType == nil {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryParameterAnnouncement) GetItemTypeOk() (*string, bool) {
	if o == nil || o.ItemType == nil {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *RepositoryParameterAnnouncement) HasItemType() bool {
	if o != nil && o.ItemType != nil {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *RepositoryParameterAnnouncement) SetItemType(v string) {
	o.ItemType = &v
}

// GetMap returns the Map field value if set, zero value otherwise.
func (o *RepositoryParameterAnnouncement) GetMap() map[string]string {
	if o == nil || o.Map == nil {
		var ret map[string]string
		return ret
	}
	return *o.Map
}

// GetMapOk returns a tuple with the Map field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryParameterAnnouncement) GetMapOk() (*map[string]string, bool) {
	if o == nil || o.Map == nil {
		return nil, false
	}
	return o.Map, true
}

// HasMap returns a boolean if a field has been set.
func (o *RepositoryParameterAnnouncement) HasMap() bool {
	if o != nil && o.Map != nil {
		return true
	}

	return false
}

// SetMap gets a reference to the given map[string]string and assigns it to the Map field.
func (o *RepositoryParameterAnnouncement) SetMap(v map[string]string) {
	o.Map = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RepositoryParameterAnnouncement) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryParameterAnnouncement) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RepositoryParameterAnnouncement) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RepositoryParameterAnnouncement) SetName(v string) {
	o.Name = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *RepositoryParameterAnnouncement) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryParameterAnnouncement) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *RepositoryParameterAnnouncement) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *RepositoryParameterAnnouncement) SetRequired(v bool) {
	o.Required = &v
}

// GetString returns the String field value if set, zero value otherwise.
func (o *RepositoryParameterAnnouncement) GetString() string {
	if o == nil || o.String == nil {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryParameterAnnouncement) GetStringOk() (*string, bool) {
	if o == nil || o.String == nil {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *RepositoryParameterAnnouncement) HasString() bool {
	if o != nil && o.String != nil {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *RepositoryParameterAnnouncement) SetString(v string) {
	o.String = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *RepositoryParameterAnnouncement) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryParameterAnnouncement) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *RepositoryParameterAnnouncement) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *RepositoryParameterAnnouncement) SetTitle(v string) {
	o.Title = &v
}

// GetTooltip returns the Tooltip field value if set, zero value otherwise.
func (o *RepositoryParameterAnnouncement) GetTooltip() string {
	if o == nil || o.Tooltip == nil {
		var ret string
		return ret
	}
	return *o.Tooltip
}

// GetTooltipOk returns a tuple with the Tooltip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryParameterAnnouncement) GetTooltipOk() (*string, bool) {
	if o == nil || o.Tooltip == nil {
		return nil, false
	}
	return o.Tooltip, true
}

// HasTooltip returns a boolean if a field has been set.
func (o *RepositoryParameterAnnouncement) HasTooltip() bool {
	if o != nil && o.Tooltip != nil {
		return true
	}

	return false
}

// SetTooltip gets a reference to the given string and assigns it to the Tooltip field.
func (o *RepositoryParameterAnnouncement) SetTooltip(v string) {
	o.Tooltip = &v
}

func (o RepositoryParameterAnnouncement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Array != nil {
		toSerialize["array"] = o.Array
	}
	if o.CollectionType != nil {
		toSerialize["collectionType"] = o.CollectionType
	}
	if o.ItemType != nil {
		toSerialize["itemType"] = o.ItemType
	}
	if o.Map != nil {
		toSerialize["map"] = o.Map
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.String != nil {
		toSerialize["string"] = o.String
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Tooltip != nil {
		toSerialize["tooltip"] = o.Tooltip
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryParameterAnnouncement struct {
	value *RepositoryParameterAnnouncement
	isSet bool
}

func (v NullableRepositoryParameterAnnouncement) Get() *RepositoryParameterAnnouncement {
	return v.value
}

func (v *NullableRepositoryParameterAnnouncement) Set(val *RepositoryParameterAnnouncement) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryParameterAnnouncement) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryParameterAnnouncement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryParameterAnnouncement(val *RepositoryParameterAnnouncement) *NullableRepositoryParameterAnnouncement {
	return &NullableRepositoryParameterAnnouncement{value: val, isSet: true}
}

func (v NullableRepositoryParameterAnnouncement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryParameterAnnouncement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}