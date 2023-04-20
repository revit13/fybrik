# AccountUpdatePasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NewPassword** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountUpdatePasswordRequest

`func NewAccountUpdatePasswordRequest() *AccountUpdatePasswordRequest`

NewAccountUpdatePasswordRequest instantiates a new AccountUpdatePasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdatePasswordRequestWithDefaults

`func NewAccountUpdatePasswordRequestWithDefaults() *AccountUpdatePasswordRequest`

NewAccountUpdatePasswordRequestWithDefaults instantiates a new AccountUpdatePasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *AccountUpdatePasswordRequest) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *AccountUpdatePasswordRequest) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *AccountUpdatePasswordRequest) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *AccountUpdatePasswordRequest) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetName

`func (o *AccountUpdatePasswordRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountUpdatePasswordRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountUpdatePasswordRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountUpdatePasswordRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewPassword

`func (o *AccountUpdatePasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *AccountUpdatePasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *AccountUpdatePasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *AccountUpdatePasswordRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


