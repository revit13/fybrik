# V1alpha1ResourceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** |  | [optional] 
**Health** | Pointer to [**V1alpha1HealthStatus**](V1alpha1HealthStatus.md) |  | [optional] 
**Hook** | Pointer to **bool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**RequiresPruning** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SyncWave** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ResourceStatus

`func NewV1alpha1ResourceStatus() *V1alpha1ResourceStatus`

NewV1alpha1ResourceStatus instantiates a new V1alpha1ResourceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ResourceStatusWithDefaults

`func NewV1alpha1ResourceStatusWithDefaults() *V1alpha1ResourceStatus`

NewV1alpha1ResourceStatusWithDefaults instantiates a new V1alpha1ResourceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *V1alpha1ResourceStatus) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *V1alpha1ResourceStatus) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *V1alpha1ResourceStatus) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *V1alpha1ResourceStatus) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHealth

`func (o *V1alpha1ResourceStatus) GetHealth() V1alpha1HealthStatus`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *V1alpha1ResourceStatus) GetHealthOk() (*V1alpha1HealthStatus, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *V1alpha1ResourceStatus) SetHealth(v V1alpha1HealthStatus)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *V1alpha1ResourceStatus) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHook

`func (o *V1alpha1ResourceStatus) GetHook() bool`

GetHook returns the Hook field if non-nil, zero value otherwise.

### GetHookOk

`func (o *V1alpha1ResourceStatus) GetHookOk() (*bool, bool)`

GetHookOk returns a tuple with the Hook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHook

`func (o *V1alpha1ResourceStatus) SetHook(v bool)`

SetHook sets Hook field to given value.

### HasHook

`func (o *V1alpha1ResourceStatus) HasHook() bool`

HasHook returns a boolean if a field has been set.

### GetKind

`func (o *V1alpha1ResourceStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1alpha1ResourceStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1alpha1ResourceStatus) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1alpha1ResourceStatus) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *V1alpha1ResourceStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha1ResourceStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha1ResourceStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1alpha1ResourceStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *V1alpha1ResourceStatus) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1alpha1ResourceStatus) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1alpha1ResourceStatus) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1alpha1ResourceStatus) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRequiresPruning

`func (o *V1alpha1ResourceStatus) GetRequiresPruning() bool`

GetRequiresPruning returns the RequiresPruning field if non-nil, zero value otherwise.

### GetRequiresPruningOk

`func (o *V1alpha1ResourceStatus) GetRequiresPruningOk() (*bool, bool)`

GetRequiresPruningOk returns a tuple with the RequiresPruning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPruning

`func (o *V1alpha1ResourceStatus) SetRequiresPruning(v bool)`

SetRequiresPruning sets RequiresPruning field to given value.

### HasRequiresPruning

`func (o *V1alpha1ResourceStatus) HasRequiresPruning() bool`

HasRequiresPruning returns a boolean if a field has been set.

### GetStatus

`func (o *V1alpha1ResourceStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1alpha1ResourceStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1alpha1ResourceStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1alpha1ResourceStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSyncWave

`func (o *V1alpha1ResourceStatus) GetSyncWave() string`

GetSyncWave returns the SyncWave field if non-nil, zero value otherwise.

### GetSyncWaveOk

`func (o *V1alpha1ResourceStatus) GetSyncWaveOk() (*string, bool)`

GetSyncWaveOk returns a tuple with the SyncWave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncWave

`func (o *V1alpha1ResourceStatus) SetSyncWave(v string)`

SetSyncWave sets SyncWave field to given value.

### HasSyncWave

`func (o *V1alpha1ResourceStatus) HasSyncWave() bool`

HasSyncWave returns a boolean if a field has been set.

### GetVersion

`func (o *V1alpha1ResourceStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V1alpha1ResourceStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V1alpha1ResourceStatus) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V1alpha1ResourceStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


