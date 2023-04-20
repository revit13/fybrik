# AccountAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tokens** | Pointer to [**[]AccountToken**](AccountToken.md) |  | [optional] 

## Methods

### NewAccountAccount

`func NewAccountAccount() *AccountAccount`

NewAccountAccount instantiates a new AccountAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAccountWithDefaults

`func NewAccountAccountWithDefaults() *AccountAccount`

NewAccountAccountWithDefaults instantiates a new AccountAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *AccountAccount) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AccountAccount) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AccountAccount) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *AccountAccount) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetEnabled

`func (o *AccountAccount) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AccountAccount) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AccountAccount) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AccountAccount) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *AccountAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTokens

`func (o *AccountAccount) GetTokens() []AccountToken`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *AccountAccount) GetTokensOk() (*[]AccountToken, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *AccountAccount) SetTokens(v []AccountToken)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *AccountAccount) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


