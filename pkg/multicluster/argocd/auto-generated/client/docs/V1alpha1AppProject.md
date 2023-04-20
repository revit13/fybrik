# V1alpha1AppProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**V1alpha1AppProjectSpec**](V1alpha1AppProjectSpec.md) |  | [optional] 
**Status** | Pointer to [**V1alpha1AppProjectStatus**](V1alpha1AppProjectStatus.md) |  | [optional] 

## Methods

### NewV1alpha1AppProject

`func NewV1alpha1AppProject() *V1alpha1AppProject`

NewV1alpha1AppProject instantiates a new V1alpha1AppProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1AppProjectWithDefaults

`func NewV1alpha1AppProjectWithDefaults() *V1alpha1AppProject`

NewV1alpha1AppProjectWithDefaults instantiates a new V1alpha1AppProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *V1alpha1AppProject) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1alpha1AppProject) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1alpha1AppProject) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1alpha1AppProject) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *V1alpha1AppProject) GetSpec() V1alpha1AppProjectSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1alpha1AppProject) GetSpecOk() (*V1alpha1AppProjectSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1alpha1AppProject) SetSpec(v V1alpha1AppProjectSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *V1alpha1AppProject) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *V1alpha1AppProject) GetStatus() V1alpha1AppProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1alpha1AppProject) GetStatusOk() (*V1alpha1AppProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1alpha1AppProject) SetStatus(v V1alpha1AppProjectStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1alpha1AppProject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


