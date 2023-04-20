# V1alpha1ClusterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]V1alpha1Cluster**](V1alpha1Cluster.md) |  | [optional] 
**Metadata** | Pointer to [**V1ListMeta**](V1ListMeta.md) |  | [optional] 

## Methods

### NewV1alpha1ClusterList

`func NewV1alpha1ClusterList() *V1alpha1ClusterList`

NewV1alpha1ClusterList instantiates a new V1alpha1ClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ClusterListWithDefaults

`func NewV1alpha1ClusterListWithDefaults() *V1alpha1ClusterList`

NewV1alpha1ClusterListWithDefaults instantiates a new V1alpha1ClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *V1alpha1ClusterList) GetItems() []V1alpha1Cluster`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1alpha1ClusterList) GetItemsOk() (*[]V1alpha1Cluster, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1alpha1ClusterList) SetItems(v []V1alpha1Cluster)`

SetItems sets Items field to given value.

### HasItems

`func (o *V1alpha1ClusterList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMetadata

`func (o *V1alpha1ClusterList) GetMetadata() V1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1alpha1ClusterList) GetMetadataOk() (*V1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1alpha1ClusterList) SetMetadata(v V1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1alpha1ClusterList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


