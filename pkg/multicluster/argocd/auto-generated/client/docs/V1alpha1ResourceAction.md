# V1alpha1ResourceAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Params** | Pointer to [**[]V1alpha1ResourceActionParam**](V1alpha1ResourceActionParam.md) |  | [optional] 

## Methods

### NewV1alpha1ResourceAction

`func NewV1alpha1ResourceAction() *V1alpha1ResourceAction`

NewV1alpha1ResourceAction instantiates a new V1alpha1ResourceAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ResourceActionWithDefaults

`func NewV1alpha1ResourceActionWithDefaults() *V1alpha1ResourceAction`

NewV1alpha1ResourceActionWithDefaults instantiates a new V1alpha1ResourceAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *V1alpha1ResourceAction) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *V1alpha1ResourceAction) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *V1alpha1ResourceAction) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *V1alpha1ResourceAction) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetName

`func (o *V1alpha1ResourceAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha1ResourceAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha1ResourceAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1alpha1ResourceAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParams

`func (o *V1alpha1ResourceAction) GetParams() []V1alpha1ResourceActionParam`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *V1alpha1ResourceAction) GetParamsOk() (*[]V1alpha1ResourceActionParam, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *V1alpha1ResourceAction) SetParams(v []V1alpha1ResourceActionParam)`

SetParams sets Params field to given value.

### HasParams

`func (o *V1alpha1ResourceAction) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


