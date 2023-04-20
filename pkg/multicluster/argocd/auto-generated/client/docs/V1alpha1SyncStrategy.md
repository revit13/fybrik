# V1alpha1SyncStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apply** | Pointer to [**V1alpha1SyncStrategyApply**](V1alpha1SyncStrategyApply.md) |  | [optional] 
**Hook** | Pointer to [**V1alpha1SyncStrategyHook**](V1alpha1SyncStrategyHook.md) |  | [optional] 

## Methods

### NewV1alpha1SyncStrategy

`func NewV1alpha1SyncStrategy() *V1alpha1SyncStrategy`

NewV1alpha1SyncStrategy instantiates a new V1alpha1SyncStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1SyncStrategyWithDefaults

`func NewV1alpha1SyncStrategyWithDefaults() *V1alpha1SyncStrategy`

NewV1alpha1SyncStrategyWithDefaults instantiates a new V1alpha1SyncStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApply

`func (o *V1alpha1SyncStrategy) GetApply() V1alpha1SyncStrategyApply`

GetApply returns the Apply field if non-nil, zero value otherwise.

### GetApplyOk

`func (o *V1alpha1SyncStrategy) GetApplyOk() (*V1alpha1SyncStrategyApply, bool)`

GetApplyOk returns a tuple with the Apply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApply

`func (o *V1alpha1SyncStrategy) SetApply(v V1alpha1SyncStrategyApply)`

SetApply sets Apply field to given value.

### HasApply

`func (o *V1alpha1SyncStrategy) HasApply() bool`

HasApply returns a boolean if a field has been set.

### GetHook

`func (o *V1alpha1SyncStrategy) GetHook() V1alpha1SyncStrategyHook`

GetHook returns the Hook field if non-nil, zero value otherwise.

### GetHookOk

`func (o *V1alpha1SyncStrategy) GetHookOk() (*V1alpha1SyncStrategyHook, bool)`

GetHookOk returns a tuple with the Hook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHook

`func (o *V1alpha1SyncStrategy) SetHook(v V1alpha1SyncStrategyHook)`

SetHook sets Hook field to given value.

### HasHook

`func (o *V1alpha1SyncStrategy) HasHook() bool`

HasHook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


