# V1alpha1Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**[]V1alpha1Info**](V1alpha1Info.md) |  | [optional] 
**InitiatedBy** | Pointer to [**V1alpha1OperationInitiator**](V1alpha1OperationInitiator.md) |  | [optional] 
**Retry** | Pointer to [**V1alpha1RetryStrategy**](V1alpha1RetryStrategy.md) |  | [optional] 
**Sync** | Pointer to [**V1alpha1SyncOperation**](V1alpha1SyncOperation.md) |  | [optional] 

## Methods

### NewV1alpha1Operation

`func NewV1alpha1Operation() *V1alpha1Operation`

NewV1alpha1Operation instantiates a new V1alpha1Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1OperationWithDefaults

`func NewV1alpha1OperationWithDefaults() *V1alpha1Operation`

NewV1alpha1OperationWithDefaults instantiates a new V1alpha1Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *V1alpha1Operation) GetInfo() []V1alpha1Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *V1alpha1Operation) GetInfoOk() (*[]V1alpha1Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *V1alpha1Operation) SetInfo(v []V1alpha1Info)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *V1alpha1Operation) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetInitiatedBy

`func (o *V1alpha1Operation) GetInitiatedBy() V1alpha1OperationInitiator`

GetInitiatedBy returns the InitiatedBy field if non-nil, zero value otherwise.

### GetInitiatedByOk

`func (o *V1alpha1Operation) GetInitiatedByOk() (*V1alpha1OperationInitiator, bool)`

GetInitiatedByOk returns a tuple with the InitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedBy

`func (o *V1alpha1Operation) SetInitiatedBy(v V1alpha1OperationInitiator)`

SetInitiatedBy sets InitiatedBy field to given value.

### HasInitiatedBy

`func (o *V1alpha1Operation) HasInitiatedBy() bool`

HasInitiatedBy returns a boolean if a field has been set.

### GetRetry

`func (o *V1alpha1Operation) GetRetry() V1alpha1RetryStrategy`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *V1alpha1Operation) GetRetryOk() (*V1alpha1RetryStrategy, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *V1alpha1Operation) SetRetry(v V1alpha1RetryStrategy)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *V1alpha1Operation) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetSync

`func (o *V1alpha1Operation) GetSync() V1alpha1SyncOperation`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *V1alpha1Operation) GetSyncOk() (*V1alpha1SyncOperation, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *V1alpha1Operation) SetSync(v V1alpha1SyncOperation)`

SetSync sets Sync field to given value.

### HasSync

`func (o *V1alpha1Operation) HasSync() bool`

HasSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


