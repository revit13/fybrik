# V1alpha1ApplicationSetTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**V1alpha1ApplicationSetTemplateMeta**](V1alpha1ApplicationSetTemplateMeta.md) |  | [optional] 
**Spec** | Pointer to [**V1alpha1ApplicationSpec**](V1alpha1ApplicationSpec.md) |  | [optional] 

## Methods

### NewV1alpha1ApplicationSetTemplate

`func NewV1alpha1ApplicationSetTemplate() *V1alpha1ApplicationSetTemplate`

NewV1alpha1ApplicationSetTemplate instantiates a new V1alpha1ApplicationSetTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSetTemplateWithDefaults

`func NewV1alpha1ApplicationSetTemplateWithDefaults() *V1alpha1ApplicationSetTemplate`

NewV1alpha1ApplicationSetTemplateWithDefaults instantiates a new V1alpha1ApplicationSetTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *V1alpha1ApplicationSetTemplate) GetMetadata() V1alpha1ApplicationSetTemplateMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1alpha1ApplicationSetTemplate) GetMetadataOk() (*V1alpha1ApplicationSetTemplateMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1alpha1ApplicationSetTemplate) SetMetadata(v V1alpha1ApplicationSetTemplateMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1alpha1ApplicationSetTemplate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *V1alpha1ApplicationSetTemplate) GetSpec() V1alpha1ApplicationSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1alpha1ApplicationSetTemplate) GetSpecOk() (*V1alpha1ApplicationSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1alpha1ApplicationSetTemplate) SetSpec(v V1alpha1ApplicationSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *V1alpha1ApplicationSetTemplate) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


