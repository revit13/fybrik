# V1alpha1SyncStrategyApply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Force** | Pointer to **bool** | Force indicates whether or not to supply the --force flag to &#x60;kubectl apply&#x60;. The --force flag deletes and re-create the resource, when PATCH encounters conflict and has retried for 5 times. | [optional] 

## Methods

### NewV1alpha1SyncStrategyApply

`func NewV1alpha1SyncStrategyApply() *V1alpha1SyncStrategyApply`

NewV1alpha1SyncStrategyApply instantiates a new V1alpha1SyncStrategyApply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1SyncStrategyApplyWithDefaults

`func NewV1alpha1SyncStrategyApplyWithDefaults() *V1alpha1SyncStrategyApply`

NewV1alpha1SyncStrategyApplyWithDefaults instantiates a new V1alpha1SyncStrategyApply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForce

`func (o *V1alpha1SyncStrategyApply) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *V1alpha1SyncStrategyApply) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *V1alpha1SyncStrategyApply) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *V1alpha1SyncStrategyApply) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


