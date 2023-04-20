# V1alpha1ApplicationSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Generators** | Pointer to [**[]V1alpha1ApplicationSetGenerator**](V1alpha1ApplicationSetGenerator.md) |  | [optional] 
**GoTemplate** | Pointer to **bool** |  | [optional] 
**Strategy** | Pointer to [**V1alpha1ApplicationSetStrategy**](V1alpha1ApplicationSetStrategy.md) |  | [optional] 
**SyncPolicy** | Pointer to [**V1alpha1ApplicationSetSyncPolicy**](V1alpha1ApplicationSetSyncPolicy.md) |  | [optional] 
**Template** | Pointer to [**V1alpha1ApplicationSetTemplate**](V1alpha1ApplicationSetTemplate.md) |  | [optional] 

## Methods

### NewV1alpha1ApplicationSetSpec

`func NewV1alpha1ApplicationSetSpec() *V1alpha1ApplicationSetSpec`

NewV1alpha1ApplicationSetSpec instantiates a new V1alpha1ApplicationSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSetSpecWithDefaults

`func NewV1alpha1ApplicationSetSpecWithDefaults() *V1alpha1ApplicationSetSpec`

NewV1alpha1ApplicationSetSpecWithDefaults instantiates a new V1alpha1ApplicationSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerators

`func (o *V1alpha1ApplicationSetSpec) GetGenerators() []V1alpha1ApplicationSetGenerator`

GetGenerators returns the Generators field if non-nil, zero value otherwise.

### GetGeneratorsOk

`func (o *V1alpha1ApplicationSetSpec) GetGeneratorsOk() (*[]V1alpha1ApplicationSetGenerator, bool)`

GetGeneratorsOk returns a tuple with the Generators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerators

`func (o *V1alpha1ApplicationSetSpec) SetGenerators(v []V1alpha1ApplicationSetGenerator)`

SetGenerators sets Generators field to given value.

### HasGenerators

`func (o *V1alpha1ApplicationSetSpec) HasGenerators() bool`

HasGenerators returns a boolean if a field has been set.

### GetGoTemplate

`func (o *V1alpha1ApplicationSetSpec) GetGoTemplate() bool`

GetGoTemplate returns the GoTemplate field if non-nil, zero value otherwise.

### GetGoTemplateOk

`func (o *V1alpha1ApplicationSetSpec) GetGoTemplateOk() (*bool, bool)`

GetGoTemplateOk returns a tuple with the GoTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoTemplate

`func (o *V1alpha1ApplicationSetSpec) SetGoTemplate(v bool)`

SetGoTemplate sets GoTemplate field to given value.

### HasGoTemplate

`func (o *V1alpha1ApplicationSetSpec) HasGoTemplate() bool`

HasGoTemplate returns a boolean if a field has been set.

### GetStrategy

`func (o *V1alpha1ApplicationSetSpec) GetStrategy() V1alpha1ApplicationSetStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *V1alpha1ApplicationSetSpec) GetStrategyOk() (*V1alpha1ApplicationSetStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *V1alpha1ApplicationSetSpec) SetStrategy(v V1alpha1ApplicationSetStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *V1alpha1ApplicationSetSpec) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetSyncPolicy

`func (o *V1alpha1ApplicationSetSpec) GetSyncPolicy() V1alpha1ApplicationSetSyncPolicy`

GetSyncPolicy returns the SyncPolicy field if non-nil, zero value otherwise.

### GetSyncPolicyOk

`func (o *V1alpha1ApplicationSetSpec) GetSyncPolicyOk() (*V1alpha1ApplicationSetSyncPolicy, bool)`

GetSyncPolicyOk returns a tuple with the SyncPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPolicy

`func (o *V1alpha1ApplicationSetSpec) SetSyncPolicy(v V1alpha1ApplicationSetSyncPolicy)`

SetSyncPolicy sets SyncPolicy field to given value.

### HasSyncPolicy

`func (o *V1alpha1ApplicationSetSpec) HasSyncPolicy() bool`

HasSyncPolicy returns a boolean if a field has been set.

### GetTemplate

`func (o *V1alpha1ApplicationSetSpec) GetTemplate() V1alpha1ApplicationSetTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1alpha1ApplicationSetSpec) GetTemplateOk() (*V1alpha1ApplicationSetTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1alpha1ApplicationSetSpec) SetTemplate(v V1alpha1ApplicationSetTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *V1alpha1ApplicationSetSpec) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


