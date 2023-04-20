# V1alpha1ApplicationSetRolloutStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchExpressions** | Pointer to [**[]V1alpha1ApplicationMatchExpression**](V1alpha1ApplicationMatchExpression.md) |  | [optional] 
**MaxUpdate** | Pointer to [**IntstrIntOrString**](IntstrIntOrString.md) |  | [optional] 

## Methods

### NewV1alpha1ApplicationSetRolloutStep

`func NewV1alpha1ApplicationSetRolloutStep() *V1alpha1ApplicationSetRolloutStep`

NewV1alpha1ApplicationSetRolloutStep instantiates a new V1alpha1ApplicationSetRolloutStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSetRolloutStepWithDefaults

`func NewV1alpha1ApplicationSetRolloutStepWithDefaults() *V1alpha1ApplicationSetRolloutStep`

NewV1alpha1ApplicationSetRolloutStepWithDefaults instantiates a new V1alpha1ApplicationSetRolloutStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchExpressions

`func (o *V1alpha1ApplicationSetRolloutStep) GetMatchExpressions() []V1alpha1ApplicationMatchExpression`

GetMatchExpressions returns the MatchExpressions field if non-nil, zero value otherwise.

### GetMatchExpressionsOk

`func (o *V1alpha1ApplicationSetRolloutStep) GetMatchExpressionsOk() (*[]V1alpha1ApplicationMatchExpression, bool)`

GetMatchExpressionsOk returns a tuple with the MatchExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchExpressions

`func (o *V1alpha1ApplicationSetRolloutStep) SetMatchExpressions(v []V1alpha1ApplicationMatchExpression)`

SetMatchExpressions sets MatchExpressions field to given value.

### HasMatchExpressions

`func (o *V1alpha1ApplicationSetRolloutStep) HasMatchExpressions() bool`

HasMatchExpressions returns a boolean if a field has been set.

### GetMaxUpdate

`func (o *V1alpha1ApplicationSetRolloutStep) GetMaxUpdate() IntstrIntOrString`

GetMaxUpdate returns the MaxUpdate field if non-nil, zero value otherwise.

### GetMaxUpdateOk

`func (o *V1alpha1ApplicationSetRolloutStep) GetMaxUpdateOk() (*IntstrIntOrString, bool)`

GetMaxUpdateOk returns a tuple with the MaxUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpdate

`func (o *V1alpha1ApplicationSetRolloutStep) SetMaxUpdate(v IntstrIntOrString)`

SetMaxUpdate sets MaxUpdate field to given value.

### HasMaxUpdate

`func (o *V1alpha1ApplicationSetRolloutStep) HasMaxUpdate() bool`

HasMaxUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


