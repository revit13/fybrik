# V1alpha1ClusterGenerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selector** | Pointer to [**V1LabelSelector**](V1LabelSelector.md) |  | [optional] 
**Template** | Pointer to [**V1alpha1ApplicationSetTemplate**](V1alpha1ApplicationSetTemplate.md) |  | [optional] 
**Values** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewV1alpha1ClusterGenerator

`func NewV1alpha1ClusterGenerator() *V1alpha1ClusterGenerator`

NewV1alpha1ClusterGenerator instantiates a new V1alpha1ClusterGenerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ClusterGeneratorWithDefaults

`func NewV1alpha1ClusterGeneratorWithDefaults() *V1alpha1ClusterGenerator`

NewV1alpha1ClusterGeneratorWithDefaults instantiates a new V1alpha1ClusterGenerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelector

`func (o *V1alpha1ClusterGenerator) GetSelector() V1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1alpha1ClusterGenerator) GetSelectorOk() (*V1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1alpha1ClusterGenerator) SetSelector(v V1LabelSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *V1alpha1ClusterGenerator) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetTemplate

`func (o *V1alpha1ClusterGenerator) GetTemplate() V1alpha1ApplicationSetTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1alpha1ClusterGenerator) GetTemplateOk() (*V1alpha1ApplicationSetTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1alpha1ClusterGenerator) SetTemplate(v V1alpha1ApplicationSetTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *V1alpha1ClusterGenerator) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetValues

`func (o *V1alpha1ClusterGenerator) GetValues() map[string]string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *V1alpha1ClusterGenerator) GetValuesOk() (*map[string]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *V1alpha1ClusterGenerator) SetValues(v map[string]string)`

SetValues sets Values field to given value.

### HasValues

`func (o *V1alpha1ClusterGenerator) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


