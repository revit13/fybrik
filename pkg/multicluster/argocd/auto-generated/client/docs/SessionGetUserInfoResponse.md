# SessionGetUserInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to **[]string** |  | [optional] 
**Iss** | Pointer to **string** |  | [optional] 
**LoggedIn** | Pointer to **bool** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionGetUserInfoResponse

`func NewSessionGetUserInfoResponse() *SessionGetUserInfoResponse`

NewSessionGetUserInfoResponse instantiates a new SessionGetUserInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionGetUserInfoResponseWithDefaults

`func NewSessionGetUserInfoResponseWithDefaults() *SessionGetUserInfoResponse`

NewSessionGetUserInfoResponseWithDefaults instantiates a new SessionGetUserInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *SessionGetUserInfoResponse) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *SessionGetUserInfoResponse) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *SessionGetUserInfoResponse) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *SessionGetUserInfoResponse) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIss

`func (o *SessionGetUserInfoResponse) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *SessionGetUserInfoResponse) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *SessionGetUserInfoResponse) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *SessionGetUserInfoResponse) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetLoggedIn

`func (o *SessionGetUserInfoResponse) GetLoggedIn() bool`

GetLoggedIn returns the LoggedIn field if non-nil, zero value otherwise.

### GetLoggedInOk

`func (o *SessionGetUserInfoResponse) GetLoggedInOk() (*bool, bool)`

GetLoggedInOk returns a tuple with the LoggedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedIn

`func (o *SessionGetUserInfoResponse) SetLoggedIn(v bool)`

SetLoggedIn sets LoggedIn field to given value.

### HasLoggedIn

`func (o *SessionGetUserInfoResponse) HasLoggedIn() bool`

HasLoggedIn returns a boolean if a field has been set.

### GetUsername

`func (o *SessionGetUserInfoResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SessionGetUserInfoResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SessionGetUserInfoResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SessionGetUserInfoResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


