# V1alpha1OperationInitiator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automated** | Pointer to **bool** | Automated is set to true if operation was initiated automatically by the application controller. | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1OperationInitiator

`func NewV1alpha1OperationInitiator() *V1alpha1OperationInitiator`

NewV1alpha1OperationInitiator instantiates a new V1alpha1OperationInitiator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1OperationInitiatorWithDefaults

`func NewV1alpha1OperationInitiatorWithDefaults() *V1alpha1OperationInitiator`

NewV1alpha1OperationInitiatorWithDefaults instantiates a new V1alpha1OperationInitiator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomated

`func (o *V1alpha1OperationInitiator) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *V1alpha1OperationInitiator) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *V1alpha1OperationInitiator) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *V1alpha1OperationInitiator) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetUsername

`func (o *V1alpha1OperationInitiator) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *V1alpha1OperationInitiator) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *V1alpha1OperationInitiator) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *V1alpha1OperationInitiator) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


