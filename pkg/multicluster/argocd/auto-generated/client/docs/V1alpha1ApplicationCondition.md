# V1alpha1ApplicationCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to [**V1Time**](V1Time.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ApplicationCondition

`func NewV1alpha1ApplicationCondition() *V1alpha1ApplicationCondition`

NewV1alpha1ApplicationCondition instantiates a new V1alpha1ApplicationCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationConditionWithDefaults

`func NewV1alpha1ApplicationConditionWithDefaults() *V1alpha1ApplicationCondition`

NewV1alpha1ApplicationConditionWithDefaults instantiates a new V1alpha1ApplicationCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *V1alpha1ApplicationCondition) GetLastTransitionTime() V1Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *V1alpha1ApplicationCondition) GetLastTransitionTimeOk() (*V1Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *V1alpha1ApplicationCondition) SetLastTransitionTime(v V1Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *V1alpha1ApplicationCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *V1alpha1ApplicationCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1alpha1ApplicationCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1alpha1ApplicationCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1alpha1ApplicationCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *V1alpha1ApplicationCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1alpha1ApplicationCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1alpha1ApplicationCondition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1alpha1ApplicationCondition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


