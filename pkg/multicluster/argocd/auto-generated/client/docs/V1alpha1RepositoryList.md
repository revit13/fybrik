# V1alpha1RepositoryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]V1alpha1Repository**](V1alpha1Repository.md) |  | [optional] 
**Metadata** | Pointer to [**V1ListMeta**](V1ListMeta.md) |  | [optional] 

## Methods

### NewV1alpha1RepositoryList

`func NewV1alpha1RepositoryList() *V1alpha1RepositoryList`

NewV1alpha1RepositoryList instantiates a new V1alpha1RepositoryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1RepositoryListWithDefaults

`func NewV1alpha1RepositoryListWithDefaults() *V1alpha1RepositoryList`

NewV1alpha1RepositoryListWithDefaults instantiates a new V1alpha1RepositoryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *V1alpha1RepositoryList) GetItems() []V1alpha1Repository`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1alpha1RepositoryList) GetItemsOk() (*[]V1alpha1Repository, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1alpha1RepositoryList) SetItems(v []V1alpha1Repository)`

SetItems sets Items field to given value.

### HasItems

`func (o *V1alpha1RepositoryList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMetadata

`func (o *V1alpha1RepositoryList) GetMetadata() V1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1alpha1RepositoryList) GetMetadataOk() (*V1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1alpha1RepositoryList) SetMetadata(v V1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1alpha1RepositoryList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


