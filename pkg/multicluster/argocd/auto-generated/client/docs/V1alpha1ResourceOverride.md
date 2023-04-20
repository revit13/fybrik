# V1alpha1ResourceOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **string** |  | [optional] 
**HealthLua** | Pointer to **string** |  | [optional] 
**IgnoreDifferences** | Pointer to [**V1alpha1OverrideIgnoreDiff**](V1alpha1OverrideIgnoreDiff.md) |  | [optional] 
**KnownTypeFields** | Pointer to [**[]V1alpha1KnownTypeField**](V1alpha1KnownTypeField.md) |  | [optional] 
**UseOpenLibs** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1alpha1ResourceOverride

`func NewV1alpha1ResourceOverride() *V1alpha1ResourceOverride`

NewV1alpha1ResourceOverride instantiates a new V1alpha1ResourceOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ResourceOverrideWithDefaults

`func NewV1alpha1ResourceOverrideWithDefaults() *V1alpha1ResourceOverride`

NewV1alpha1ResourceOverrideWithDefaults instantiates a new V1alpha1ResourceOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *V1alpha1ResourceOverride) GetActions() string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *V1alpha1ResourceOverride) GetActionsOk() (*string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *V1alpha1ResourceOverride) SetActions(v string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *V1alpha1ResourceOverride) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetHealthLua

`func (o *V1alpha1ResourceOverride) GetHealthLua() string`

GetHealthLua returns the HealthLua field if non-nil, zero value otherwise.

### GetHealthLuaOk

`func (o *V1alpha1ResourceOverride) GetHealthLuaOk() (*string, bool)`

GetHealthLuaOk returns a tuple with the HealthLua field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthLua

`func (o *V1alpha1ResourceOverride) SetHealthLua(v string)`

SetHealthLua sets HealthLua field to given value.

### HasHealthLua

`func (o *V1alpha1ResourceOverride) HasHealthLua() bool`

HasHealthLua returns a boolean if a field has been set.

### GetIgnoreDifferences

`func (o *V1alpha1ResourceOverride) GetIgnoreDifferences() V1alpha1OverrideIgnoreDiff`

GetIgnoreDifferences returns the IgnoreDifferences field if non-nil, zero value otherwise.

### GetIgnoreDifferencesOk

`func (o *V1alpha1ResourceOverride) GetIgnoreDifferencesOk() (*V1alpha1OverrideIgnoreDiff, bool)`

GetIgnoreDifferencesOk returns a tuple with the IgnoreDifferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDifferences

`func (o *V1alpha1ResourceOverride) SetIgnoreDifferences(v V1alpha1OverrideIgnoreDiff)`

SetIgnoreDifferences sets IgnoreDifferences field to given value.

### HasIgnoreDifferences

`func (o *V1alpha1ResourceOverride) HasIgnoreDifferences() bool`

HasIgnoreDifferences returns a boolean if a field has been set.

### GetKnownTypeFields

`func (o *V1alpha1ResourceOverride) GetKnownTypeFields() []V1alpha1KnownTypeField`

GetKnownTypeFields returns the KnownTypeFields field if non-nil, zero value otherwise.

### GetKnownTypeFieldsOk

`func (o *V1alpha1ResourceOverride) GetKnownTypeFieldsOk() (*[]V1alpha1KnownTypeField, bool)`

GetKnownTypeFieldsOk returns a tuple with the KnownTypeFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownTypeFields

`func (o *V1alpha1ResourceOverride) SetKnownTypeFields(v []V1alpha1KnownTypeField)`

SetKnownTypeFields sets KnownTypeFields field to given value.

### HasKnownTypeFields

`func (o *V1alpha1ResourceOverride) HasKnownTypeFields() bool`

HasKnownTypeFields returns a boolean if a field has been set.

### GetUseOpenLibs

`func (o *V1alpha1ResourceOverride) GetUseOpenLibs() bool`

GetUseOpenLibs returns the UseOpenLibs field if non-nil, zero value otherwise.

### GetUseOpenLibsOk

`func (o *V1alpha1ResourceOverride) GetUseOpenLibsOk() (*bool, bool)`

GetUseOpenLibsOk returns a tuple with the UseOpenLibs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOpenLibs

`func (o *V1alpha1ResourceOverride) SetUseOpenLibs(v bool)`

SetUseOpenLibs sets UseOpenLibs field to given value.

### HasUseOpenLibs

`func (o *V1alpha1ResourceOverride) HasUseOpenLibs() bool`

HasUseOpenLibs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


