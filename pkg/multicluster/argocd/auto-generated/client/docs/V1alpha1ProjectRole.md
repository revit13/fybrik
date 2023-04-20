# V1alpha1ProjectRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**JwtTokens** | Pointer to [**[]V1alpha1JWTToken**](V1alpha1JWTToken.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Policies** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1alpha1ProjectRole

`func NewV1alpha1ProjectRole() *V1alpha1ProjectRole`

NewV1alpha1ProjectRole instantiates a new V1alpha1ProjectRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ProjectRoleWithDefaults

`func NewV1alpha1ProjectRoleWithDefaults() *V1alpha1ProjectRole`

NewV1alpha1ProjectRoleWithDefaults instantiates a new V1alpha1ProjectRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V1alpha1ProjectRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1alpha1ProjectRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1alpha1ProjectRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1alpha1ProjectRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroups

`func (o *V1alpha1ProjectRole) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V1alpha1ProjectRole) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V1alpha1ProjectRole) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V1alpha1ProjectRole) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetJwtTokens

`func (o *V1alpha1ProjectRole) GetJwtTokens() []V1alpha1JWTToken`

GetJwtTokens returns the JwtTokens field if non-nil, zero value otherwise.

### GetJwtTokensOk

`func (o *V1alpha1ProjectRole) GetJwtTokensOk() (*[]V1alpha1JWTToken, bool)`

GetJwtTokensOk returns a tuple with the JwtTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTokens

`func (o *V1alpha1ProjectRole) SetJwtTokens(v []V1alpha1JWTToken)`

SetJwtTokens sets JwtTokens field to given value.

### HasJwtTokens

`func (o *V1alpha1ProjectRole) HasJwtTokens() bool`

HasJwtTokens returns a boolean if a field has been set.

### GetName

`func (o *V1alpha1ProjectRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha1ProjectRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha1ProjectRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1alpha1ProjectRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicies

`func (o *V1alpha1ProjectRole) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *V1alpha1ProjectRole) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *V1alpha1ProjectRole) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *V1alpha1ProjectRole) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


