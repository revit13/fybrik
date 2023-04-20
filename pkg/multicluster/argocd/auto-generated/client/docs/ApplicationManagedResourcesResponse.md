# ApplicationManagedResourcesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]V1alpha1ResourceDiff**](V1alpha1ResourceDiff.md) |  | [optional] 

## Methods

### NewApplicationManagedResourcesResponse

`func NewApplicationManagedResourcesResponse() *ApplicationManagedResourcesResponse`

NewApplicationManagedResourcesResponse instantiates a new ApplicationManagedResourcesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationManagedResourcesResponseWithDefaults

`func NewApplicationManagedResourcesResponseWithDefaults() *ApplicationManagedResourcesResponse`

NewApplicationManagedResourcesResponseWithDefaults instantiates a new ApplicationManagedResourcesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ApplicationManagedResourcesResponse) GetItems() []V1alpha1ResourceDiff`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApplicationManagedResourcesResponse) GetItemsOk() (*[]V1alpha1ResourceDiff, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApplicationManagedResourcesResponse) SetItems(v []V1alpha1ResourceDiff)`

SetItems sets Items field to given value.

### HasItems

`func (o *ApplicationManagedResourcesResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


