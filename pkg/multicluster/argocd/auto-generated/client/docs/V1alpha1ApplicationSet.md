# V1alpha1ApplicationSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**V1alpha1ApplicationSetSpec**](V1alpha1ApplicationSetSpec.md) |  | [optional] 
**Status** | Pointer to [**V1alpha1ApplicationSetStatus**](V1alpha1ApplicationSetStatus.md) |  | [optional] 

## Methods

### NewV1alpha1ApplicationSet

`func NewV1alpha1ApplicationSet() *V1alpha1ApplicationSet`

NewV1alpha1ApplicationSet instantiates a new V1alpha1ApplicationSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSetWithDefaults

`func NewV1alpha1ApplicationSetWithDefaults() *V1alpha1ApplicationSet`

NewV1alpha1ApplicationSetWithDefaults instantiates a new V1alpha1ApplicationSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *V1alpha1ApplicationSet) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1alpha1ApplicationSet) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1alpha1ApplicationSet) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1alpha1ApplicationSet) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *V1alpha1ApplicationSet) GetSpec() V1alpha1ApplicationSetSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1alpha1ApplicationSet) GetSpecOk() (*V1alpha1ApplicationSetSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1alpha1ApplicationSet) SetSpec(v V1alpha1ApplicationSetSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *V1alpha1ApplicationSet) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *V1alpha1ApplicationSet) GetStatus() V1alpha1ApplicationSetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1alpha1ApplicationSet) GetStatusOk() (*V1alpha1ApplicationSetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1alpha1ApplicationSet) SetStatus(v V1alpha1ApplicationSetStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1alpha1ApplicationSet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


