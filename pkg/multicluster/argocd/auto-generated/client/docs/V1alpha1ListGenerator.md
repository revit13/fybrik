# V1alpha1ListGenerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elements** | Pointer to [**[]V1JSON**](V1JSON.md) |  | [optional] 
**Template** | Pointer to [**V1alpha1ApplicationSetTemplate**](V1alpha1ApplicationSetTemplate.md) |  | [optional] 

## Methods

### NewV1alpha1ListGenerator

`func NewV1alpha1ListGenerator() *V1alpha1ListGenerator`

NewV1alpha1ListGenerator instantiates a new V1alpha1ListGenerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ListGeneratorWithDefaults

`func NewV1alpha1ListGeneratorWithDefaults() *V1alpha1ListGenerator`

NewV1alpha1ListGeneratorWithDefaults instantiates a new V1alpha1ListGenerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElements

`func (o *V1alpha1ListGenerator) GetElements() []V1JSON`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *V1alpha1ListGenerator) GetElementsOk() (*[]V1JSON, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *V1alpha1ListGenerator) SetElements(v []V1JSON)`

SetElements sets Elements field to given value.

### HasElements

`func (o *V1alpha1ListGenerator) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetTemplate

`func (o *V1alpha1ListGenerator) GetTemplate() V1alpha1ApplicationSetTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1alpha1ListGenerator) GetTemplateOk() (*V1alpha1ApplicationSetTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1alpha1ListGenerator) SetTemplate(v V1alpha1ApplicationSetTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *V1alpha1ListGenerator) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


