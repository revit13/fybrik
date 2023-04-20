# V1alpha1ApplicationSetStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RollingSync** | Pointer to [**V1alpha1ApplicationSetRolloutStrategy**](V1alpha1ApplicationSetRolloutStrategy.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ApplicationSetStrategy

`func NewV1alpha1ApplicationSetStrategy() *V1alpha1ApplicationSetStrategy`

NewV1alpha1ApplicationSetStrategy instantiates a new V1alpha1ApplicationSetStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSetStrategyWithDefaults

`func NewV1alpha1ApplicationSetStrategyWithDefaults() *V1alpha1ApplicationSetStrategy`

NewV1alpha1ApplicationSetStrategyWithDefaults instantiates a new V1alpha1ApplicationSetStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollingSync

`func (o *V1alpha1ApplicationSetStrategy) GetRollingSync() V1alpha1ApplicationSetRolloutStrategy`

GetRollingSync returns the RollingSync field if non-nil, zero value otherwise.

### GetRollingSyncOk

`func (o *V1alpha1ApplicationSetStrategy) GetRollingSyncOk() (*V1alpha1ApplicationSetRolloutStrategy, bool)`

GetRollingSyncOk returns a tuple with the RollingSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingSync

`func (o *V1alpha1ApplicationSetStrategy) SetRollingSync(v V1alpha1ApplicationSetRolloutStrategy)`

SetRollingSync sets RollingSync field to given value.

### HasRollingSync

`func (o *V1alpha1ApplicationSetStrategy) HasRollingSync() bool`

HasRollingSync returns a boolean if a field has been set.

### GetType

`func (o *V1alpha1ApplicationSetStrategy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1alpha1ApplicationSetStrategy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1alpha1ApplicationSetStrategy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1alpha1ApplicationSetStrategy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


