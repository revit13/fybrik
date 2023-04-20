# V1alpha1ApplicationSetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationStatus** | Pointer to [**[]V1alpha1ApplicationSetApplicationStatus**](V1alpha1ApplicationSetApplicationStatus.md) |  | [optional] 
**Conditions** | Pointer to [**[]V1alpha1ApplicationSetCondition**](V1alpha1ApplicationSetCondition.md) |  | [optional] 

## Methods

### NewV1alpha1ApplicationSetStatus

`func NewV1alpha1ApplicationSetStatus() *V1alpha1ApplicationSetStatus`

NewV1alpha1ApplicationSetStatus instantiates a new V1alpha1ApplicationSetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSetStatusWithDefaults

`func NewV1alpha1ApplicationSetStatusWithDefaults() *V1alpha1ApplicationSetStatus`

NewV1alpha1ApplicationSetStatusWithDefaults instantiates a new V1alpha1ApplicationSetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationStatus

`func (o *V1alpha1ApplicationSetStatus) GetApplicationStatus() []V1alpha1ApplicationSetApplicationStatus`

GetApplicationStatus returns the ApplicationStatus field if non-nil, zero value otherwise.

### GetApplicationStatusOk

`func (o *V1alpha1ApplicationSetStatus) GetApplicationStatusOk() (*[]V1alpha1ApplicationSetApplicationStatus, bool)`

GetApplicationStatusOk returns a tuple with the ApplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationStatus

`func (o *V1alpha1ApplicationSetStatus) SetApplicationStatus(v []V1alpha1ApplicationSetApplicationStatus)`

SetApplicationStatus sets ApplicationStatus field to given value.

### HasApplicationStatus

`func (o *V1alpha1ApplicationSetStatus) HasApplicationStatus() bool`

HasApplicationStatus returns a boolean if a field has been set.

### GetConditions

`func (o *V1alpha1ApplicationSetStatus) GetConditions() []V1alpha1ApplicationSetCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1alpha1ApplicationSetStatus) GetConditionsOk() (*[]V1alpha1ApplicationSetCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1alpha1ApplicationSetStatus) SetConditions(v []V1alpha1ApplicationSetCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1alpha1ApplicationSetStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


