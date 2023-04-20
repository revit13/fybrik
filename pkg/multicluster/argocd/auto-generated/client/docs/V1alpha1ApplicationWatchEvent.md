# V1alpha1ApplicationWatchEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**V1alpha1Application**](V1alpha1Application.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ApplicationWatchEvent

`func NewV1alpha1ApplicationWatchEvent() *V1alpha1ApplicationWatchEvent`

NewV1alpha1ApplicationWatchEvent instantiates a new V1alpha1ApplicationWatchEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationWatchEventWithDefaults

`func NewV1alpha1ApplicationWatchEventWithDefaults() *V1alpha1ApplicationWatchEvent`

NewV1alpha1ApplicationWatchEventWithDefaults instantiates a new V1alpha1ApplicationWatchEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *V1alpha1ApplicationWatchEvent) GetApplication() V1alpha1Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *V1alpha1ApplicationWatchEvent) GetApplicationOk() (*V1alpha1Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *V1alpha1ApplicationWatchEvent) SetApplication(v V1alpha1Application)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *V1alpha1ApplicationWatchEvent) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetType

`func (o *V1alpha1ApplicationWatchEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1alpha1ApplicationWatchEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1alpha1ApplicationWatchEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1alpha1ApplicationWatchEvent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


