# V1alpha1ApplicationSetApplicationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** |  | [optional] 
**LastTransitionTime** | Pointer to [**V1Time**](V1Time.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Step** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ApplicationSetApplicationStatus

`func NewV1alpha1ApplicationSetApplicationStatus() *V1alpha1ApplicationSetApplicationStatus`

NewV1alpha1ApplicationSetApplicationStatus instantiates a new V1alpha1ApplicationSetApplicationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSetApplicationStatusWithDefaults

`func NewV1alpha1ApplicationSetApplicationStatusWithDefaults() *V1alpha1ApplicationSetApplicationStatus`

NewV1alpha1ApplicationSetApplicationStatusWithDefaults instantiates a new V1alpha1ApplicationSetApplicationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *V1alpha1ApplicationSetApplicationStatus) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *V1alpha1ApplicationSetApplicationStatus) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *V1alpha1ApplicationSetApplicationStatus) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *V1alpha1ApplicationSetApplicationStatus) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *V1alpha1ApplicationSetApplicationStatus) GetLastTransitionTime() V1Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *V1alpha1ApplicationSetApplicationStatus) GetLastTransitionTimeOk() (*V1Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *V1alpha1ApplicationSetApplicationStatus) SetLastTransitionTime(v V1Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *V1alpha1ApplicationSetApplicationStatus) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *V1alpha1ApplicationSetApplicationStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1alpha1ApplicationSetApplicationStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1alpha1ApplicationSetApplicationStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1alpha1ApplicationSetApplicationStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *V1alpha1ApplicationSetApplicationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1alpha1ApplicationSetApplicationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1alpha1ApplicationSetApplicationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1alpha1ApplicationSetApplicationStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStep

`func (o *V1alpha1ApplicationSetApplicationStatus) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *V1alpha1ApplicationSetApplicationStatus) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *V1alpha1ApplicationSetApplicationStatus) SetStep(v string)`

SetStep sets Step field to given value.

### HasStep

`func (o *V1alpha1ApplicationSetApplicationStatus) HasStep() bool`

HasStep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


