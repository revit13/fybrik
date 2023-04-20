# StreamResultOfApplicationLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**RuntimeStreamError**](RuntimeStreamError.md) |  | [optional] 
**Result** | Pointer to [**ApplicationLogEntry**](ApplicationLogEntry.md) |  | [optional] 

## Methods

### NewStreamResultOfApplicationLogEntry

`func NewStreamResultOfApplicationLogEntry() *StreamResultOfApplicationLogEntry`

NewStreamResultOfApplicationLogEntry instantiates a new StreamResultOfApplicationLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamResultOfApplicationLogEntryWithDefaults

`func NewStreamResultOfApplicationLogEntryWithDefaults() *StreamResultOfApplicationLogEntry`

NewStreamResultOfApplicationLogEntryWithDefaults instantiates a new StreamResultOfApplicationLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StreamResultOfApplicationLogEntry) GetError() RuntimeStreamError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamResultOfApplicationLogEntry) GetErrorOk() (*RuntimeStreamError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamResultOfApplicationLogEntry) SetError(v RuntimeStreamError)`

SetError sets Error field to given value.

### HasError

`func (o *StreamResultOfApplicationLogEntry) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResult

`func (o *StreamResultOfApplicationLogEntry) GetResult() ApplicationLogEntry`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StreamResultOfApplicationLogEntry) GetResultOk() (*ApplicationLogEntry, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StreamResultOfApplicationLogEntry) SetResult(v ApplicationLogEntry)`

SetResult sets Result field to given value.

### HasResult

`func (o *StreamResultOfApplicationLogEntry) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


