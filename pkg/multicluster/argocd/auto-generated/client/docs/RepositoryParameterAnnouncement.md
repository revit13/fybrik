# RepositoryParameterAnnouncement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Array** | Pointer to **[]string** | array is the default value of the parameter if the parameter is an array. | [optional] 
**CollectionType** | Pointer to **string** | collectionType is the type of value this parameter holds - either a single value (a string) or a collection (array or map). If collectionType is set, only the field with that type will be used. If collectionType is not set, &#x60;string&#x60; is the default. If collectionType is set to an invalid value, a validation error is thrown. | [optional] 
**ItemType** | Pointer to **string** | itemType determines the primitive data type represented by the parameter. Parameters are always encoded as strings, but this field lets them be interpreted as other primitive types. | [optional] 
**Map** | Pointer to **map[string]string** | map is the default value of the parameter if the parameter is a map. | [optional] 
**Name** | Pointer to **string** | name is the name identifying a parameter. | [optional] 
**Required** | Pointer to **bool** | required defines if this given parameter is mandatory. | [optional] 
**String** | Pointer to **string** | string is the default value of the parameter if the parameter is a string. | [optional] 
**Title** | Pointer to **string** | title is a human-readable text of the parameter name. | [optional] 
**Tooltip** | Pointer to **string** | tooltip is a human-readable description of the parameter. | [optional] 

## Methods

### NewRepositoryParameterAnnouncement

`func NewRepositoryParameterAnnouncement() *RepositoryParameterAnnouncement`

NewRepositoryParameterAnnouncement instantiates a new RepositoryParameterAnnouncement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryParameterAnnouncementWithDefaults

`func NewRepositoryParameterAnnouncementWithDefaults() *RepositoryParameterAnnouncement`

NewRepositoryParameterAnnouncementWithDefaults instantiates a new RepositoryParameterAnnouncement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArray

`func (o *RepositoryParameterAnnouncement) GetArray() []string`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *RepositoryParameterAnnouncement) GetArrayOk() (*[]string, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *RepositoryParameterAnnouncement) SetArray(v []string)`

SetArray sets Array field to given value.

### HasArray

`func (o *RepositoryParameterAnnouncement) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetCollectionType

`func (o *RepositoryParameterAnnouncement) GetCollectionType() string`

GetCollectionType returns the CollectionType field if non-nil, zero value otherwise.

### GetCollectionTypeOk

`func (o *RepositoryParameterAnnouncement) GetCollectionTypeOk() (*string, bool)`

GetCollectionTypeOk returns a tuple with the CollectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionType

`func (o *RepositoryParameterAnnouncement) SetCollectionType(v string)`

SetCollectionType sets CollectionType field to given value.

### HasCollectionType

`func (o *RepositoryParameterAnnouncement) HasCollectionType() bool`

HasCollectionType returns a boolean if a field has been set.

### GetItemType

`func (o *RepositoryParameterAnnouncement) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *RepositoryParameterAnnouncement) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *RepositoryParameterAnnouncement) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *RepositoryParameterAnnouncement) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetMap

`func (o *RepositoryParameterAnnouncement) GetMap() map[string]string`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *RepositoryParameterAnnouncement) GetMapOk() (*map[string]string, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *RepositoryParameterAnnouncement) SetMap(v map[string]string)`

SetMap sets Map field to given value.

### HasMap

`func (o *RepositoryParameterAnnouncement) HasMap() bool`

HasMap returns a boolean if a field has been set.

### GetName

`func (o *RepositoryParameterAnnouncement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryParameterAnnouncement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryParameterAnnouncement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RepositoryParameterAnnouncement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *RepositoryParameterAnnouncement) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *RepositoryParameterAnnouncement) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *RepositoryParameterAnnouncement) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *RepositoryParameterAnnouncement) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetString

`func (o *RepositoryParameterAnnouncement) GetString() string`

GetString returns the String field if non-nil, zero value otherwise.

### GetStringOk

`func (o *RepositoryParameterAnnouncement) GetStringOk() (*string, bool)`

GetStringOk returns a tuple with the String field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetString

`func (o *RepositoryParameterAnnouncement) SetString(v string)`

SetString sets String field to given value.

### HasString

`func (o *RepositoryParameterAnnouncement) HasString() bool`

HasString returns a boolean if a field has been set.

### GetTitle

`func (o *RepositoryParameterAnnouncement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RepositoryParameterAnnouncement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RepositoryParameterAnnouncement) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RepositoryParameterAnnouncement) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTooltip

`func (o *RepositoryParameterAnnouncement) GetTooltip() string`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *RepositoryParameterAnnouncement) GetTooltipOk() (*string, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *RepositoryParameterAnnouncement) SetTooltip(v string)`

SetTooltip sets Tooltip field to given value.

### HasTooltip

`func (o *RepositoryParameterAnnouncement) HasTooltip() bool`

HasTooltip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


