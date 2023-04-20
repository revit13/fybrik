# V1alpha1ResourceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** |  | [optional] 
**HookPhase** | Pointer to **string** | HookPhase contains the state of any operation associated with this resource OR hook This can also contain values for non-hook resources. | [optional] 
**HookType** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SyncPhase** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ResourceResult

`func NewV1alpha1ResourceResult() *V1alpha1ResourceResult`

NewV1alpha1ResourceResult instantiates a new V1alpha1ResourceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ResourceResultWithDefaults

`func NewV1alpha1ResourceResultWithDefaults() *V1alpha1ResourceResult`

NewV1alpha1ResourceResultWithDefaults instantiates a new V1alpha1ResourceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *V1alpha1ResourceResult) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *V1alpha1ResourceResult) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *V1alpha1ResourceResult) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *V1alpha1ResourceResult) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHookPhase

`func (o *V1alpha1ResourceResult) GetHookPhase() string`

GetHookPhase returns the HookPhase field if non-nil, zero value otherwise.

### GetHookPhaseOk

`func (o *V1alpha1ResourceResult) GetHookPhaseOk() (*string, bool)`

GetHookPhaseOk returns a tuple with the HookPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookPhase

`func (o *V1alpha1ResourceResult) SetHookPhase(v string)`

SetHookPhase sets HookPhase field to given value.

### HasHookPhase

`func (o *V1alpha1ResourceResult) HasHookPhase() bool`

HasHookPhase returns a boolean if a field has been set.

### GetHookType

`func (o *V1alpha1ResourceResult) GetHookType() string`

GetHookType returns the HookType field if non-nil, zero value otherwise.

### GetHookTypeOk

`func (o *V1alpha1ResourceResult) GetHookTypeOk() (*string, bool)`

GetHookTypeOk returns a tuple with the HookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookType

`func (o *V1alpha1ResourceResult) SetHookType(v string)`

SetHookType sets HookType field to given value.

### HasHookType

`func (o *V1alpha1ResourceResult) HasHookType() bool`

HasHookType returns a boolean if a field has been set.

### GetKind

`func (o *V1alpha1ResourceResult) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1alpha1ResourceResult) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1alpha1ResourceResult) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1alpha1ResourceResult) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMessage

`func (o *V1alpha1ResourceResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1alpha1ResourceResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1alpha1ResourceResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1alpha1ResourceResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *V1alpha1ResourceResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha1ResourceResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha1ResourceResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1alpha1ResourceResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *V1alpha1ResourceResult) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1alpha1ResourceResult) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1alpha1ResourceResult) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1alpha1ResourceResult) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetStatus

`func (o *V1alpha1ResourceResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1alpha1ResourceResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1alpha1ResourceResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1alpha1ResourceResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSyncPhase

`func (o *V1alpha1ResourceResult) GetSyncPhase() string`

GetSyncPhase returns the SyncPhase field if non-nil, zero value otherwise.

### GetSyncPhaseOk

`func (o *V1alpha1ResourceResult) GetSyncPhaseOk() (*string, bool)`

GetSyncPhaseOk returns a tuple with the SyncPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPhase

`func (o *V1alpha1ResourceResult) SetSyncPhase(v string)`

SetSyncPhase sets SyncPhase field to given value.

### HasSyncPhase

`func (o *V1alpha1ResourceResult) HasSyncPhase() bool`

HasSyncPhase returns a boolean if a field has been set.

### GetVersion

`func (o *V1alpha1ResourceResult) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V1alpha1ResourceResult) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V1alpha1ResourceResult) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V1alpha1ResourceResult) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


