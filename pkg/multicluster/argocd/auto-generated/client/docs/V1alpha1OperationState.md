# V1alpha1OperationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinishedAt** | Pointer to [**V1Time**](V1Time.md) |  | [optional] 
**Message** | Pointer to **string** | Message holds any pertinent messages when attempting to perform operation (typically errors). | [optional] 
**Operation** | Pointer to [**V1alpha1Operation**](V1alpha1Operation.md) |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**RetryCount** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to [**V1Time**](V1Time.md) |  | [optional] 
**SyncResult** | Pointer to [**V1alpha1SyncOperationResult**](V1alpha1SyncOperationResult.md) |  | [optional] 

## Methods

### NewV1alpha1OperationState

`func NewV1alpha1OperationState() *V1alpha1OperationState`

NewV1alpha1OperationState instantiates a new V1alpha1OperationState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1OperationStateWithDefaults

`func NewV1alpha1OperationStateWithDefaults() *V1alpha1OperationState`

NewV1alpha1OperationStateWithDefaults instantiates a new V1alpha1OperationState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinishedAt

`func (o *V1alpha1OperationState) GetFinishedAt() V1Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *V1alpha1OperationState) GetFinishedAtOk() (*V1Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *V1alpha1OperationState) SetFinishedAt(v V1Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *V1alpha1OperationState) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetMessage

`func (o *V1alpha1OperationState) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1alpha1OperationState) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1alpha1OperationState) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1alpha1OperationState) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOperation

`func (o *V1alpha1OperationState) GetOperation() V1alpha1Operation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *V1alpha1OperationState) GetOperationOk() (*V1alpha1Operation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *V1alpha1OperationState) SetOperation(v V1alpha1Operation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *V1alpha1OperationState) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPhase

`func (o *V1alpha1OperationState) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V1alpha1OperationState) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V1alpha1OperationState) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *V1alpha1OperationState) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetRetryCount

`func (o *V1alpha1OperationState) GetRetryCount() string`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *V1alpha1OperationState) GetRetryCountOk() (*string, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *V1alpha1OperationState) SetRetryCount(v string)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *V1alpha1OperationState) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetStartedAt

`func (o *V1alpha1OperationState) GetStartedAt() V1Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *V1alpha1OperationState) GetStartedAtOk() (*V1Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *V1alpha1OperationState) SetStartedAt(v V1Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *V1alpha1OperationState) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetSyncResult

`func (o *V1alpha1OperationState) GetSyncResult() V1alpha1SyncOperationResult`

GetSyncResult returns the SyncResult field if non-nil, zero value otherwise.

### GetSyncResultOk

`func (o *V1alpha1OperationState) GetSyncResultOk() (*V1alpha1SyncOperationResult, bool)`

GetSyncResultOk returns a tuple with the SyncResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncResult

`func (o *V1alpha1OperationState) SetSyncResult(v V1alpha1SyncOperationResult)`

SetSyncResult sets SyncResult field to given value.

### HasSyncResult

`func (o *V1alpha1OperationState) HasSyncResult() bool`

HasSyncResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


