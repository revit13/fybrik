# V1alpha1OrphanedResourcesMonitorSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ignore** | Pointer to [**[]V1alpha1OrphanedResourceKey**](V1alpha1OrphanedResourceKey.md) |  | [optional] 
**Warn** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1alpha1OrphanedResourcesMonitorSettings

`func NewV1alpha1OrphanedResourcesMonitorSettings() *V1alpha1OrphanedResourcesMonitorSettings`

NewV1alpha1OrphanedResourcesMonitorSettings instantiates a new V1alpha1OrphanedResourcesMonitorSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1OrphanedResourcesMonitorSettingsWithDefaults

`func NewV1alpha1OrphanedResourcesMonitorSettingsWithDefaults() *V1alpha1OrphanedResourcesMonitorSettings`

NewV1alpha1OrphanedResourcesMonitorSettingsWithDefaults instantiates a new V1alpha1OrphanedResourcesMonitorSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnore

`func (o *V1alpha1OrphanedResourcesMonitorSettings) GetIgnore() []V1alpha1OrphanedResourceKey`

GetIgnore returns the Ignore field if non-nil, zero value otherwise.

### GetIgnoreOk

`func (o *V1alpha1OrphanedResourcesMonitorSettings) GetIgnoreOk() (*[]V1alpha1OrphanedResourceKey, bool)`

GetIgnoreOk returns a tuple with the Ignore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnore

`func (o *V1alpha1OrphanedResourcesMonitorSettings) SetIgnore(v []V1alpha1OrphanedResourceKey)`

SetIgnore sets Ignore field to given value.

### HasIgnore

`func (o *V1alpha1OrphanedResourcesMonitorSettings) HasIgnore() bool`

HasIgnore returns a boolean if a field has been set.

### GetWarn

`func (o *V1alpha1OrphanedResourcesMonitorSettings) GetWarn() bool`

GetWarn returns the Warn field if non-nil, zero value otherwise.

### GetWarnOk

`func (o *V1alpha1OrphanedResourcesMonitorSettings) GetWarnOk() (*bool, bool)`

GetWarnOk returns a tuple with the Warn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarn

`func (o *V1alpha1OrphanedResourcesMonitorSettings) SetWarn(v bool)`

SetWarn sets Warn field to given value.

### HasWarn

`func (o *V1alpha1OrphanedResourcesMonitorSettings) HasWarn() bool`

HasWarn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


