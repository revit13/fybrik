# V1alpha1ResourceDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Diff** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Hook** | Pointer to **bool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**LiveState** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**NormalizedLiveState** | Pointer to **string** |  | [optional] 
**PredictedLiveState** | Pointer to **string** |  | [optional] 
**ResourceVersion** | Pointer to **string** |  | [optional] 
**TargetState** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ResourceDiff

`func NewV1alpha1ResourceDiff() *V1alpha1ResourceDiff`

NewV1alpha1ResourceDiff instantiates a new V1alpha1ResourceDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ResourceDiffWithDefaults

`func NewV1alpha1ResourceDiffWithDefaults() *V1alpha1ResourceDiff`

NewV1alpha1ResourceDiffWithDefaults instantiates a new V1alpha1ResourceDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiff

`func (o *V1alpha1ResourceDiff) GetDiff() string`

GetDiff returns the Diff field if non-nil, zero value otherwise.

### GetDiffOk

`func (o *V1alpha1ResourceDiff) GetDiffOk() (*string, bool)`

GetDiffOk returns a tuple with the Diff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiff

`func (o *V1alpha1ResourceDiff) SetDiff(v string)`

SetDiff sets Diff field to given value.

### HasDiff

`func (o *V1alpha1ResourceDiff) HasDiff() bool`

HasDiff returns a boolean if a field has been set.

### GetGroup

`func (o *V1alpha1ResourceDiff) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *V1alpha1ResourceDiff) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *V1alpha1ResourceDiff) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *V1alpha1ResourceDiff) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHook

`func (o *V1alpha1ResourceDiff) GetHook() bool`

GetHook returns the Hook field if non-nil, zero value otherwise.

### GetHookOk

`func (o *V1alpha1ResourceDiff) GetHookOk() (*bool, bool)`

GetHookOk returns a tuple with the Hook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHook

`func (o *V1alpha1ResourceDiff) SetHook(v bool)`

SetHook sets Hook field to given value.

### HasHook

`func (o *V1alpha1ResourceDiff) HasHook() bool`

HasHook returns a boolean if a field has been set.

### GetKind

`func (o *V1alpha1ResourceDiff) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1alpha1ResourceDiff) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1alpha1ResourceDiff) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1alpha1ResourceDiff) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLiveState

`func (o *V1alpha1ResourceDiff) GetLiveState() string`

GetLiveState returns the LiveState field if non-nil, zero value otherwise.

### GetLiveStateOk

`func (o *V1alpha1ResourceDiff) GetLiveStateOk() (*string, bool)`

GetLiveStateOk returns a tuple with the LiveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveState

`func (o *V1alpha1ResourceDiff) SetLiveState(v string)`

SetLiveState sets LiveState field to given value.

### HasLiveState

`func (o *V1alpha1ResourceDiff) HasLiveState() bool`

HasLiveState returns a boolean if a field has been set.

### GetModified

`func (o *V1alpha1ResourceDiff) GetModified() bool`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *V1alpha1ResourceDiff) GetModifiedOk() (*bool, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *V1alpha1ResourceDiff) SetModified(v bool)`

SetModified sets Modified field to given value.

### HasModified

`func (o *V1alpha1ResourceDiff) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *V1alpha1ResourceDiff) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha1ResourceDiff) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha1ResourceDiff) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1alpha1ResourceDiff) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *V1alpha1ResourceDiff) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1alpha1ResourceDiff) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1alpha1ResourceDiff) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1alpha1ResourceDiff) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNormalizedLiveState

`func (o *V1alpha1ResourceDiff) GetNormalizedLiveState() string`

GetNormalizedLiveState returns the NormalizedLiveState field if non-nil, zero value otherwise.

### GetNormalizedLiveStateOk

`func (o *V1alpha1ResourceDiff) GetNormalizedLiveStateOk() (*string, bool)`

GetNormalizedLiveStateOk returns a tuple with the NormalizedLiveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedLiveState

`func (o *V1alpha1ResourceDiff) SetNormalizedLiveState(v string)`

SetNormalizedLiveState sets NormalizedLiveState field to given value.

### HasNormalizedLiveState

`func (o *V1alpha1ResourceDiff) HasNormalizedLiveState() bool`

HasNormalizedLiveState returns a boolean if a field has been set.

### GetPredictedLiveState

`func (o *V1alpha1ResourceDiff) GetPredictedLiveState() string`

GetPredictedLiveState returns the PredictedLiveState field if non-nil, zero value otherwise.

### GetPredictedLiveStateOk

`func (o *V1alpha1ResourceDiff) GetPredictedLiveStateOk() (*string, bool)`

GetPredictedLiveStateOk returns a tuple with the PredictedLiveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedLiveState

`func (o *V1alpha1ResourceDiff) SetPredictedLiveState(v string)`

SetPredictedLiveState sets PredictedLiveState field to given value.

### HasPredictedLiveState

`func (o *V1alpha1ResourceDiff) HasPredictedLiveState() bool`

HasPredictedLiveState returns a boolean if a field has been set.

### GetResourceVersion

`func (o *V1alpha1ResourceDiff) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *V1alpha1ResourceDiff) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *V1alpha1ResourceDiff) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *V1alpha1ResourceDiff) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetTargetState

`func (o *V1alpha1ResourceDiff) GetTargetState() string`

GetTargetState returns the TargetState field if non-nil, zero value otherwise.

### GetTargetStateOk

`func (o *V1alpha1ResourceDiff) GetTargetStateOk() (*string, bool)`

GetTargetStateOk returns a tuple with the TargetState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetState

`func (o *V1alpha1ResourceDiff) SetTargetState(v string)`

SetTargetState sets TargetState field to given value.

### HasTargetState

`func (o *V1alpha1ResourceDiff) HasTargetState() bool`

HasTargetState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


