# V1alpha1HostInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ResourcesInfo** | Pointer to [**[]V1alpha1HostResourceInfo**](V1alpha1HostResourceInfo.md) |  | [optional] 
**SystemInfo** | Pointer to [**V1NodeSystemInfo**](V1NodeSystemInfo.md) |  | [optional] 

## Methods

### NewV1alpha1HostInfo

`func NewV1alpha1HostInfo() *V1alpha1HostInfo`

NewV1alpha1HostInfo instantiates a new V1alpha1HostInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1HostInfoWithDefaults

`func NewV1alpha1HostInfoWithDefaults() *V1alpha1HostInfo`

NewV1alpha1HostInfoWithDefaults instantiates a new V1alpha1HostInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1alpha1HostInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha1HostInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha1HostInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1alpha1HostInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourcesInfo

`func (o *V1alpha1HostInfo) GetResourcesInfo() []V1alpha1HostResourceInfo`

GetResourcesInfo returns the ResourcesInfo field if non-nil, zero value otherwise.

### GetResourcesInfoOk

`func (o *V1alpha1HostInfo) GetResourcesInfoOk() (*[]V1alpha1HostResourceInfo, bool)`

GetResourcesInfoOk returns a tuple with the ResourcesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesInfo

`func (o *V1alpha1HostInfo) SetResourcesInfo(v []V1alpha1HostResourceInfo)`

SetResourcesInfo sets ResourcesInfo field to given value.

### HasResourcesInfo

`func (o *V1alpha1HostInfo) HasResourcesInfo() bool`

HasResourcesInfo returns a boolean if a field has been set.

### GetSystemInfo

`func (o *V1alpha1HostInfo) GetSystemInfo() V1NodeSystemInfo`

GetSystemInfo returns the SystemInfo field if non-nil, zero value otherwise.

### GetSystemInfoOk

`func (o *V1alpha1HostInfo) GetSystemInfoOk() (*V1NodeSystemInfo, bool)`

GetSystemInfoOk returns a tuple with the SystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInfo

`func (o *V1alpha1HostInfo) SetSystemInfo(v V1NodeSystemInfo)`

SetSystemInfo sets SystemInfo field to given value.

### HasSystemInfo

`func (o *V1alpha1HostInfo) HasSystemInfo() bool`

HasSystemInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


