# V1alpha1ConnectionState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptedAt** | Pointer to [**V1Time**](V1Time.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ConnectionState

`func NewV1alpha1ConnectionState() *V1alpha1ConnectionState`

NewV1alpha1ConnectionState instantiates a new V1alpha1ConnectionState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ConnectionStateWithDefaults

`func NewV1alpha1ConnectionStateWithDefaults() *V1alpha1ConnectionState`

NewV1alpha1ConnectionStateWithDefaults instantiates a new V1alpha1ConnectionState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptedAt

`func (o *V1alpha1ConnectionState) GetAttemptedAt() V1Time`

GetAttemptedAt returns the AttemptedAt field if non-nil, zero value otherwise.

### GetAttemptedAtOk

`func (o *V1alpha1ConnectionState) GetAttemptedAtOk() (*V1Time, bool)`

GetAttemptedAtOk returns a tuple with the AttemptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedAt

`func (o *V1alpha1ConnectionState) SetAttemptedAt(v V1Time)`

SetAttemptedAt sets AttemptedAt field to given value.

### HasAttemptedAt

`func (o *V1alpha1ConnectionState) HasAttemptedAt() bool`

HasAttemptedAt returns a boolean if a field has been set.

### GetMessage

`func (o *V1alpha1ConnectionState) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1alpha1ConnectionState) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1alpha1ConnectionState) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1alpha1ConnectionState) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *V1alpha1ConnectionState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1alpha1ConnectionState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1alpha1ConnectionState) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1alpha1ConnectionState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


