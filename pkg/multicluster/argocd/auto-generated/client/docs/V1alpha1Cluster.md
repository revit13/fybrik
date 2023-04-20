# V1alpha1Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**ClusterResources** | Pointer to **bool** | Indicates if cluster level resources should be managed. This setting is used only if cluster is connected in a namespaced mode. | [optional] 
**Config** | Pointer to [**V1alpha1ClusterConfig**](V1alpha1ClusterConfig.md) |  | [optional] 
**ConnectionState** | Pointer to [**V1alpha1ConnectionState**](V1alpha1ConnectionState.md) |  | [optional] 
**Info** | Pointer to [**V1alpha1ClusterInfo**](V1alpha1ClusterInfo.md) |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespaces** | Pointer to **[]string** | Holds list of namespaces which are accessible in that cluster. Cluster level resources will be ignored if namespace list is not empty. | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**RefreshRequestedAt** | Pointer to **string** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**ServerVersion** | Pointer to **string** |  | [optional] 
**Shard** | Pointer to **string** | Shard contains optional shard number. Calculated on the fly by the application controller if not specified. | [optional] 

## Methods

### NewV1alpha1Cluster

`func NewV1alpha1Cluster() *V1alpha1Cluster`

NewV1alpha1Cluster instantiates a new V1alpha1Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ClusterWithDefaults

`func NewV1alpha1ClusterWithDefaults() *V1alpha1Cluster`

NewV1alpha1ClusterWithDefaults instantiates a new V1alpha1Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *V1alpha1Cluster) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *V1alpha1Cluster) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *V1alpha1Cluster) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *V1alpha1Cluster) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetClusterResources

`func (o *V1alpha1Cluster) GetClusterResources() bool`

GetClusterResources returns the ClusterResources field if non-nil, zero value otherwise.

### GetClusterResourcesOk

`func (o *V1alpha1Cluster) GetClusterResourcesOk() (*bool, bool)`

GetClusterResourcesOk returns a tuple with the ClusterResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterResources

`func (o *V1alpha1Cluster) SetClusterResources(v bool)`

SetClusterResources sets ClusterResources field to given value.

### HasClusterResources

`func (o *V1alpha1Cluster) HasClusterResources() bool`

HasClusterResources returns a boolean if a field has been set.

### GetConfig

`func (o *V1alpha1Cluster) GetConfig() V1alpha1ClusterConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1alpha1Cluster) GetConfigOk() (*V1alpha1ClusterConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1alpha1Cluster) SetConfig(v V1alpha1ClusterConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V1alpha1Cluster) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConnectionState

`func (o *V1alpha1Cluster) GetConnectionState() V1alpha1ConnectionState`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *V1alpha1Cluster) GetConnectionStateOk() (*V1alpha1ConnectionState, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *V1alpha1Cluster) SetConnectionState(v V1alpha1ConnectionState)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *V1alpha1Cluster) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetInfo

`func (o *V1alpha1Cluster) GetInfo() V1alpha1ClusterInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *V1alpha1Cluster) GetInfoOk() (*V1alpha1ClusterInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *V1alpha1Cluster) SetInfo(v V1alpha1ClusterInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *V1alpha1Cluster) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetLabels

`func (o *V1alpha1Cluster) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1alpha1Cluster) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1alpha1Cluster) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1alpha1Cluster) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *V1alpha1Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha1Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha1Cluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1alpha1Cluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaces

`func (o *V1alpha1Cluster) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *V1alpha1Cluster) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *V1alpha1Cluster) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *V1alpha1Cluster) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetProject

`func (o *V1alpha1Cluster) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *V1alpha1Cluster) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *V1alpha1Cluster) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *V1alpha1Cluster) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRefreshRequestedAt

`func (o *V1alpha1Cluster) GetRefreshRequestedAt() string`

GetRefreshRequestedAt returns the RefreshRequestedAt field if non-nil, zero value otherwise.

### GetRefreshRequestedAtOk

`func (o *V1alpha1Cluster) GetRefreshRequestedAtOk() (*string, bool)`

GetRefreshRequestedAtOk returns a tuple with the RefreshRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshRequestedAt

`func (o *V1alpha1Cluster) SetRefreshRequestedAt(v string)`

SetRefreshRequestedAt sets RefreshRequestedAt field to given value.

### HasRefreshRequestedAt

`func (o *V1alpha1Cluster) HasRefreshRequestedAt() bool`

HasRefreshRequestedAt returns a boolean if a field has been set.

### GetServer

`func (o *V1alpha1Cluster) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *V1alpha1Cluster) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *V1alpha1Cluster) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *V1alpha1Cluster) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetServerVersion

`func (o *V1alpha1Cluster) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *V1alpha1Cluster) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *V1alpha1Cluster) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *V1alpha1Cluster) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetShard

`func (o *V1alpha1Cluster) GetShard() string`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *V1alpha1Cluster) GetShardOk() (*string, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *V1alpha1Cluster) SetShard(v string)`

SetShard sets Shard field to given value.

### HasShard

`func (o *V1alpha1Cluster) HasShard() bool`

HasShard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


