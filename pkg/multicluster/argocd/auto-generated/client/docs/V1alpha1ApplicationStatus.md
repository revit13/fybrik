# V1alpha1ApplicationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]V1alpha1ApplicationCondition**](V1alpha1ApplicationCondition.md) |  | [optional] 
**Health** | Pointer to [**V1alpha1HealthStatus**](V1alpha1HealthStatus.md) |  | [optional] 
**History** | Pointer to [**[]V1alpha1RevisionHistory**](V1alpha1RevisionHistory.md) |  | [optional] 
**ObservedAt** | Pointer to [**V1Time**](V1Time.md) |  | [optional] 
**OperationState** | Pointer to [**V1alpha1OperationState**](V1alpha1OperationState.md) |  | [optional] 
**ReconciledAt** | Pointer to [**V1Time**](V1Time.md) |  | [optional] 
**ResourceHealthSource** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**[]V1alpha1ResourceStatus**](V1alpha1ResourceStatus.md) |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**SourceTypes** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to [**V1alpha1ApplicationSummary**](V1alpha1ApplicationSummary.md) |  | [optional] 
**Sync** | Pointer to [**V1alpha1SyncStatus**](V1alpha1SyncStatus.md) |  | [optional] 

## Methods

### NewV1alpha1ApplicationStatus

`func NewV1alpha1ApplicationStatus() *V1alpha1ApplicationStatus`

NewV1alpha1ApplicationStatus instantiates a new V1alpha1ApplicationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationStatusWithDefaults

`func NewV1alpha1ApplicationStatusWithDefaults() *V1alpha1ApplicationStatus`

NewV1alpha1ApplicationStatusWithDefaults instantiates a new V1alpha1ApplicationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *V1alpha1ApplicationStatus) GetConditions() []V1alpha1ApplicationCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1alpha1ApplicationStatus) GetConditionsOk() (*[]V1alpha1ApplicationCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1alpha1ApplicationStatus) SetConditions(v []V1alpha1ApplicationCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1alpha1ApplicationStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetHealth

`func (o *V1alpha1ApplicationStatus) GetHealth() V1alpha1HealthStatus`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *V1alpha1ApplicationStatus) GetHealthOk() (*V1alpha1HealthStatus, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *V1alpha1ApplicationStatus) SetHealth(v V1alpha1HealthStatus)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *V1alpha1ApplicationStatus) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *V1alpha1ApplicationStatus) GetHistory() []V1alpha1RevisionHistory`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *V1alpha1ApplicationStatus) GetHistoryOk() (*[]V1alpha1RevisionHistory, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *V1alpha1ApplicationStatus) SetHistory(v []V1alpha1RevisionHistory)`

SetHistory sets History field to given value.

### HasHistory

`func (o *V1alpha1ApplicationStatus) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetObservedAt

`func (o *V1alpha1ApplicationStatus) GetObservedAt() V1Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *V1alpha1ApplicationStatus) GetObservedAtOk() (*V1Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *V1alpha1ApplicationStatus) SetObservedAt(v V1Time)`

SetObservedAt sets ObservedAt field to given value.

### HasObservedAt

`func (o *V1alpha1ApplicationStatus) HasObservedAt() bool`

HasObservedAt returns a boolean if a field has been set.

### GetOperationState

`func (o *V1alpha1ApplicationStatus) GetOperationState() V1alpha1OperationState`

GetOperationState returns the OperationState field if non-nil, zero value otherwise.

### GetOperationStateOk

`func (o *V1alpha1ApplicationStatus) GetOperationStateOk() (*V1alpha1OperationState, bool)`

GetOperationStateOk returns a tuple with the OperationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationState

`func (o *V1alpha1ApplicationStatus) SetOperationState(v V1alpha1OperationState)`

SetOperationState sets OperationState field to given value.

### HasOperationState

`func (o *V1alpha1ApplicationStatus) HasOperationState() bool`

HasOperationState returns a boolean if a field has been set.

### GetReconciledAt

`func (o *V1alpha1ApplicationStatus) GetReconciledAt() V1Time`

GetReconciledAt returns the ReconciledAt field if non-nil, zero value otherwise.

### GetReconciledAtOk

`func (o *V1alpha1ApplicationStatus) GetReconciledAtOk() (*V1Time, bool)`

GetReconciledAtOk returns a tuple with the ReconciledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciledAt

`func (o *V1alpha1ApplicationStatus) SetReconciledAt(v V1Time)`

SetReconciledAt sets ReconciledAt field to given value.

### HasReconciledAt

`func (o *V1alpha1ApplicationStatus) HasReconciledAt() bool`

HasReconciledAt returns a boolean if a field has been set.

### GetResourceHealthSource

`func (o *V1alpha1ApplicationStatus) GetResourceHealthSource() string`

GetResourceHealthSource returns the ResourceHealthSource field if non-nil, zero value otherwise.

### GetResourceHealthSourceOk

`func (o *V1alpha1ApplicationStatus) GetResourceHealthSourceOk() (*string, bool)`

GetResourceHealthSourceOk returns a tuple with the ResourceHealthSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceHealthSource

`func (o *V1alpha1ApplicationStatus) SetResourceHealthSource(v string)`

SetResourceHealthSource sets ResourceHealthSource field to given value.

### HasResourceHealthSource

`func (o *V1alpha1ApplicationStatus) HasResourceHealthSource() bool`

HasResourceHealthSource returns a boolean if a field has been set.

### GetResources

`func (o *V1alpha1ApplicationStatus) GetResources() []V1alpha1ResourceStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V1alpha1ApplicationStatus) GetResourcesOk() (*[]V1alpha1ResourceStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V1alpha1ApplicationStatus) SetResources(v []V1alpha1ResourceStatus)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V1alpha1ApplicationStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSourceType

`func (o *V1alpha1ApplicationStatus) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *V1alpha1ApplicationStatus) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *V1alpha1ApplicationStatus) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *V1alpha1ApplicationStatus) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSourceTypes

`func (o *V1alpha1ApplicationStatus) GetSourceTypes() []string`

GetSourceTypes returns the SourceTypes field if non-nil, zero value otherwise.

### GetSourceTypesOk

`func (o *V1alpha1ApplicationStatus) GetSourceTypesOk() (*[]string, bool)`

GetSourceTypesOk returns a tuple with the SourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTypes

`func (o *V1alpha1ApplicationStatus) SetSourceTypes(v []string)`

SetSourceTypes sets SourceTypes field to given value.

### HasSourceTypes

`func (o *V1alpha1ApplicationStatus) HasSourceTypes() bool`

HasSourceTypes returns a boolean if a field has been set.

### GetSummary

`func (o *V1alpha1ApplicationStatus) GetSummary() V1alpha1ApplicationSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *V1alpha1ApplicationStatus) GetSummaryOk() (*V1alpha1ApplicationSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *V1alpha1ApplicationStatus) SetSummary(v V1alpha1ApplicationSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *V1alpha1ApplicationStatus) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetSync

`func (o *V1alpha1ApplicationStatus) GetSync() V1alpha1SyncStatus`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *V1alpha1ApplicationStatus) GetSyncOk() (*V1alpha1SyncStatus, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *V1alpha1ApplicationStatus) SetSync(v V1alpha1SyncStatus)`

SetSync sets Sync field to given value.

### HasSync

`func (o *V1alpha1ApplicationStatus) HasSync() bool`

HasSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


