# V1alpha1ClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersions** | Pointer to **[]string** |  | [optional] 
**ApplicationsCount** | Pointer to **string** |  | [optional] 
**CacheInfo** | Pointer to [**V1alpha1ClusterCacheInfo**](V1alpha1ClusterCacheInfo.md) |  | [optional] 
**ConnectionState** | Pointer to [**V1alpha1ConnectionState**](V1alpha1ConnectionState.md) |  | [optional] 
**ServerVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ClusterInfo

`func NewV1alpha1ClusterInfo() *V1alpha1ClusterInfo`

NewV1alpha1ClusterInfo instantiates a new V1alpha1ClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ClusterInfoWithDefaults

`func NewV1alpha1ClusterInfoWithDefaults() *V1alpha1ClusterInfo`

NewV1alpha1ClusterInfoWithDefaults instantiates a new V1alpha1ClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersions

`func (o *V1alpha1ClusterInfo) GetApiVersions() []string`

GetApiVersions returns the ApiVersions field if non-nil, zero value otherwise.

### GetApiVersionsOk

`func (o *V1alpha1ClusterInfo) GetApiVersionsOk() (*[]string, bool)`

GetApiVersionsOk returns a tuple with the ApiVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersions

`func (o *V1alpha1ClusterInfo) SetApiVersions(v []string)`

SetApiVersions sets ApiVersions field to given value.

### HasApiVersions

`func (o *V1alpha1ClusterInfo) HasApiVersions() bool`

HasApiVersions returns a boolean if a field has been set.

### GetApplicationsCount

`func (o *V1alpha1ClusterInfo) GetApplicationsCount() string`

GetApplicationsCount returns the ApplicationsCount field if non-nil, zero value otherwise.

### GetApplicationsCountOk

`func (o *V1alpha1ClusterInfo) GetApplicationsCountOk() (*string, bool)`

GetApplicationsCountOk returns a tuple with the ApplicationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationsCount

`func (o *V1alpha1ClusterInfo) SetApplicationsCount(v string)`

SetApplicationsCount sets ApplicationsCount field to given value.

### HasApplicationsCount

`func (o *V1alpha1ClusterInfo) HasApplicationsCount() bool`

HasApplicationsCount returns a boolean if a field has been set.

### GetCacheInfo

`func (o *V1alpha1ClusterInfo) GetCacheInfo() V1alpha1ClusterCacheInfo`

GetCacheInfo returns the CacheInfo field if non-nil, zero value otherwise.

### GetCacheInfoOk

`func (o *V1alpha1ClusterInfo) GetCacheInfoOk() (*V1alpha1ClusterCacheInfo, bool)`

GetCacheInfoOk returns a tuple with the CacheInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheInfo

`func (o *V1alpha1ClusterInfo) SetCacheInfo(v V1alpha1ClusterCacheInfo)`

SetCacheInfo sets CacheInfo field to given value.

### HasCacheInfo

`func (o *V1alpha1ClusterInfo) HasCacheInfo() bool`

HasCacheInfo returns a boolean if a field has been set.

### GetConnectionState

`func (o *V1alpha1ClusterInfo) GetConnectionState() V1alpha1ConnectionState`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *V1alpha1ClusterInfo) GetConnectionStateOk() (*V1alpha1ConnectionState, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *V1alpha1ClusterInfo) SetConnectionState(v V1alpha1ConnectionState)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *V1alpha1ClusterInfo) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetServerVersion

`func (o *V1alpha1ClusterInfo) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *V1alpha1ClusterInfo) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *V1alpha1ClusterInfo) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *V1alpha1ClusterInfo) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


