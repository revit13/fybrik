# V1alpha1DuckTypeGenerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMapRef** | Pointer to **string** |  | [optional] 
**LabelSelector** | Pointer to [**V1LabelSelector**](V1LabelSelector.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RequeueAfterSeconds** | Pointer to **string** |  | [optional] 
**Template** | Pointer to [**V1alpha1ApplicationSetTemplate**](V1alpha1ApplicationSetTemplate.md) |  | [optional] 
**Values** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewV1alpha1DuckTypeGenerator

`func NewV1alpha1DuckTypeGenerator() *V1alpha1DuckTypeGenerator`

NewV1alpha1DuckTypeGenerator instantiates a new V1alpha1DuckTypeGenerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1DuckTypeGeneratorWithDefaults

`func NewV1alpha1DuckTypeGeneratorWithDefaults() *V1alpha1DuckTypeGenerator`

NewV1alpha1DuckTypeGeneratorWithDefaults instantiates a new V1alpha1DuckTypeGenerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMapRef

`func (o *V1alpha1DuckTypeGenerator) GetConfigMapRef() string`

GetConfigMapRef returns the ConfigMapRef field if non-nil, zero value otherwise.

### GetConfigMapRefOk

`func (o *V1alpha1DuckTypeGenerator) GetConfigMapRefOk() (*string, bool)`

GetConfigMapRefOk returns a tuple with the ConfigMapRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMapRef

`func (o *V1alpha1DuckTypeGenerator) SetConfigMapRef(v string)`

SetConfigMapRef sets ConfigMapRef field to given value.

### HasConfigMapRef

`func (o *V1alpha1DuckTypeGenerator) HasConfigMapRef() bool`

HasConfigMapRef returns a boolean if a field has been set.

### GetLabelSelector

`func (o *V1alpha1DuckTypeGenerator) GetLabelSelector() V1LabelSelector`

GetLabelSelector returns the LabelSelector field if non-nil, zero value otherwise.

### GetLabelSelectorOk

`func (o *V1alpha1DuckTypeGenerator) GetLabelSelectorOk() (*V1LabelSelector, bool)`

GetLabelSelectorOk returns a tuple with the LabelSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSelector

`func (o *V1alpha1DuckTypeGenerator) SetLabelSelector(v V1LabelSelector)`

SetLabelSelector sets LabelSelector field to given value.

### HasLabelSelector

`func (o *V1alpha1DuckTypeGenerator) HasLabelSelector() bool`

HasLabelSelector returns a boolean if a field has been set.

### GetName

`func (o *V1alpha1DuckTypeGenerator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha1DuckTypeGenerator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha1DuckTypeGenerator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1alpha1DuckTypeGenerator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequeueAfterSeconds

`func (o *V1alpha1DuckTypeGenerator) GetRequeueAfterSeconds() string`

GetRequeueAfterSeconds returns the RequeueAfterSeconds field if non-nil, zero value otherwise.

### GetRequeueAfterSecondsOk

`func (o *V1alpha1DuckTypeGenerator) GetRequeueAfterSecondsOk() (*string, bool)`

GetRequeueAfterSecondsOk returns a tuple with the RequeueAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequeueAfterSeconds

`func (o *V1alpha1DuckTypeGenerator) SetRequeueAfterSeconds(v string)`

SetRequeueAfterSeconds sets RequeueAfterSeconds field to given value.

### HasRequeueAfterSeconds

`func (o *V1alpha1DuckTypeGenerator) HasRequeueAfterSeconds() bool`

HasRequeueAfterSeconds returns a boolean if a field has been set.

### GetTemplate

`func (o *V1alpha1DuckTypeGenerator) GetTemplate() V1alpha1ApplicationSetTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1alpha1DuckTypeGenerator) GetTemplateOk() (*V1alpha1ApplicationSetTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1alpha1DuckTypeGenerator) SetTemplate(v V1alpha1ApplicationSetTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *V1alpha1DuckTypeGenerator) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetValues

`func (o *V1alpha1DuckTypeGenerator) GetValues() map[string]string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *V1alpha1DuckTypeGenerator) GetValuesOk() (*map[string]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *V1alpha1DuckTypeGenerator) SetValues(v map[string]string)`

SetValues sets Values field to given value.

### HasValues

`func (o *V1alpha1DuckTypeGenerator) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


