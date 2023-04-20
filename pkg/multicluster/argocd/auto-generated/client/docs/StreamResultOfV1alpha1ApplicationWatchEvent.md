# StreamResultOfV1alpha1ApplicationWatchEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**RuntimeStreamError**](RuntimeStreamError.md) |  | [optional] 
**Result** | Pointer to [**V1alpha1ApplicationWatchEvent**](V1alpha1ApplicationWatchEvent.md) |  | [optional] 

## Methods

### NewStreamResultOfV1alpha1ApplicationWatchEvent

`func NewStreamResultOfV1alpha1ApplicationWatchEvent() *StreamResultOfV1alpha1ApplicationWatchEvent`

NewStreamResultOfV1alpha1ApplicationWatchEvent instantiates a new StreamResultOfV1alpha1ApplicationWatchEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamResultOfV1alpha1ApplicationWatchEventWithDefaults

`func NewStreamResultOfV1alpha1ApplicationWatchEventWithDefaults() *StreamResultOfV1alpha1ApplicationWatchEvent`

NewStreamResultOfV1alpha1ApplicationWatchEventWithDefaults instantiates a new StreamResultOfV1alpha1ApplicationWatchEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StreamResultOfV1alpha1ApplicationWatchEvent) GetError() RuntimeStreamError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamResultOfV1alpha1ApplicationWatchEvent) GetErrorOk() (*RuntimeStreamError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamResultOfV1alpha1ApplicationWatchEvent) SetError(v RuntimeStreamError)`

SetError sets Error field to given value.

### HasError

`func (o *StreamResultOfV1alpha1ApplicationWatchEvent) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResult

`func (o *StreamResultOfV1alpha1ApplicationWatchEvent) GetResult() V1alpha1ApplicationWatchEvent`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StreamResultOfV1alpha1ApplicationWatchEvent) GetResultOk() (*V1alpha1ApplicationWatchEvent, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StreamResultOfV1alpha1ApplicationWatchEvent) SetResult(v V1alpha1ApplicationWatchEvent)`

SetResult sets Result field to given value.

### HasResult

`func (o *StreamResultOfV1alpha1ApplicationWatchEvent) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


