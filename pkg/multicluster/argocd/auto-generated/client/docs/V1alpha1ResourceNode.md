# V1alpha1ResourceNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**V1Time**](V1Time.md) |  | [optional] 
**Health** | Pointer to [**V1alpha1HealthStatus**](V1alpha1HealthStatus.md) |  | [optional] 
**Images** | Pointer to **[]string** |  | [optional] 
**Info** | Pointer to [**[]V1alpha1InfoItem**](V1alpha1InfoItem.md) |  | [optional] 
**NetworkingInfo** | Pointer to [**V1alpha1ResourceNetworkingInfo**](V1alpha1ResourceNetworkingInfo.md) |  | [optional] 
**ParentRefs** | Pointer to [**[]V1alpha1ResourceRef**](V1alpha1ResourceRef.md) |  | [optional] 
**ResourceRef** | Pointer to [**V1alpha1ResourceRef**](V1alpha1ResourceRef.md) |  | [optional] 
**ResourceVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ResourceNode

`func NewV1alpha1ResourceNode() *V1alpha1ResourceNode`

NewV1alpha1ResourceNode instantiates a new V1alpha1ResourceNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ResourceNodeWithDefaults

`func NewV1alpha1ResourceNodeWithDefaults() *V1alpha1ResourceNode`

NewV1alpha1ResourceNodeWithDefaults instantiates a new V1alpha1ResourceNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *V1alpha1ResourceNode) GetCreatedAt() V1Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V1alpha1ResourceNode) GetCreatedAtOk() (*V1Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V1alpha1ResourceNode) SetCreatedAt(v V1Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V1alpha1ResourceNode) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHealth

`func (o *V1alpha1ResourceNode) GetHealth() V1alpha1HealthStatus`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *V1alpha1ResourceNode) GetHealthOk() (*V1alpha1HealthStatus, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *V1alpha1ResourceNode) SetHealth(v V1alpha1HealthStatus)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *V1alpha1ResourceNode) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetImages

`func (o *V1alpha1ResourceNode) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *V1alpha1ResourceNode) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *V1alpha1ResourceNode) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *V1alpha1ResourceNode) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetInfo

`func (o *V1alpha1ResourceNode) GetInfo() []V1alpha1InfoItem`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *V1alpha1ResourceNode) GetInfoOk() (*[]V1alpha1InfoItem, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *V1alpha1ResourceNode) SetInfo(v []V1alpha1InfoItem)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *V1alpha1ResourceNode) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetNetworkingInfo

`func (o *V1alpha1ResourceNode) GetNetworkingInfo() V1alpha1ResourceNetworkingInfo`

GetNetworkingInfo returns the NetworkingInfo field if non-nil, zero value otherwise.

### GetNetworkingInfoOk

`func (o *V1alpha1ResourceNode) GetNetworkingInfoOk() (*V1alpha1ResourceNetworkingInfo, bool)`

GetNetworkingInfoOk returns a tuple with the NetworkingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingInfo

`func (o *V1alpha1ResourceNode) SetNetworkingInfo(v V1alpha1ResourceNetworkingInfo)`

SetNetworkingInfo sets NetworkingInfo field to given value.

### HasNetworkingInfo

`func (o *V1alpha1ResourceNode) HasNetworkingInfo() bool`

HasNetworkingInfo returns a boolean if a field has been set.

### GetParentRefs

`func (o *V1alpha1ResourceNode) GetParentRefs() []V1alpha1ResourceRef`

GetParentRefs returns the ParentRefs field if non-nil, zero value otherwise.

### GetParentRefsOk

`func (o *V1alpha1ResourceNode) GetParentRefsOk() (*[]V1alpha1ResourceRef, bool)`

GetParentRefsOk returns a tuple with the ParentRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRefs

`func (o *V1alpha1ResourceNode) SetParentRefs(v []V1alpha1ResourceRef)`

SetParentRefs sets ParentRefs field to given value.

### HasParentRefs

`func (o *V1alpha1ResourceNode) HasParentRefs() bool`

HasParentRefs returns a boolean if a field has been set.

### GetResourceRef

`func (o *V1alpha1ResourceNode) GetResourceRef() V1alpha1ResourceRef`

GetResourceRef returns the ResourceRef field if non-nil, zero value otherwise.

### GetResourceRefOk

`func (o *V1alpha1ResourceNode) GetResourceRefOk() (*V1alpha1ResourceRef, bool)`

GetResourceRefOk returns a tuple with the ResourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRef

`func (o *V1alpha1ResourceNode) SetResourceRef(v V1alpha1ResourceRef)`

SetResourceRef sets ResourceRef field to given value.

### HasResourceRef

`func (o *V1alpha1ResourceNode) HasResourceRef() bool`

HasResourceRef returns a boolean if a field has been set.

### GetResourceVersion

`func (o *V1alpha1ResourceNode) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *V1alpha1ResourceNode) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *V1alpha1ResourceNode) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *V1alpha1ResourceNode) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


