# V1alpha1MergeGenerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Generators** | Pointer to [**[]V1alpha1ApplicationSetNestedGenerator**](V1alpha1ApplicationSetNestedGenerator.md) |  | [optional] 
**MergeKeys** | Pointer to **[]string** |  | [optional] 
**Template** | Pointer to [**V1alpha1ApplicationSetTemplate**](V1alpha1ApplicationSetTemplate.md) |  | [optional] 

## Methods

### NewV1alpha1MergeGenerator

`func NewV1alpha1MergeGenerator() *V1alpha1MergeGenerator`

NewV1alpha1MergeGenerator instantiates a new V1alpha1MergeGenerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1MergeGeneratorWithDefaults

`func NewV1alpha1MergeGeneratorWithDefaults() *V1alpha1MergeGenerator`

NewV1alpha1MergeGeneratorWithDefaults instantiates a new V1alpha1MergeGenerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerators

`func (o *V1alpha1MergeGenerator) GetGenerators() []V1alpha1ApplicationSetNestedGenerator`

GetGenerators returns the Generators field if non-nil, zero value otherwise.

### GetGeneratorsOk

`func (o *V1alpha1MergeGenerator) GetGeneratorsOk() (*[]V1alpha1ApplicationSetNestedGenerator, bool)`

GetGeneratorsOk returns a tuple with the Generators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerators

`func (o *V1alpha1MergeGenerator) SetGenerators(v []V1alpha1ApplicationSetNestedGenerator)`

SetGenerators sets Generators field to given value.

### HasGenerators

`func (o *V1alpha1MergeGenerator) HasGenerators() bool`

HasGenerators returns a boolean if a field has been set.

### GetMergeKeys

`func (o *V1alpha1MergeGenerator) GetMergeKeys() []string`

GetMergeKeys returns the MergeKeys field if non-nil, zero value otherwise.

### GetMergeKeysOk

`func (o *V1alpha1MergeGenerator) GetMergeKeysOk() (*[]string, bool)`

GetMergeKeysOk returns a tuple with the MergeKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeKeys

`func (o *V1alpha1MergeGenerator) SetMergeKeys(v []string)`

SetMergeKeys sets MergeKeys field to given value.

### HasMergeKeys

`func (o *V1alpha1MergeGenerator) HasMergeKeys() bool`

HasMergeKeys returns a boolean if a field has been set.

### GetTemplate

`func (o *V1alpha1MergeGenerator) GetTemplate() V1alpha1ApplicationSetTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1alpha1MergeGenerator) GetTemplateOk() (*V1alpha1ApplicationSetTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1alpha1MergeGenerator) SetTemplate(v V1alpha1ApplicationSetTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *V1alpha1MergeGenerator) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


