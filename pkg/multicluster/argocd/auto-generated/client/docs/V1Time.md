# V1Time

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nanos** | Pointer to **int32** | Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. | [optional] 
**Seconds** | Pointer to **string** | Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. | [optional] 

## Methods

### NewV1Time

`func NewV1Time() *V1Time`

NewV1Time instantiates a new V1Time object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TimeWithDefaults

`func NewV1TimeWithDefaults() *V1Time`

NewV1TimeWithDefaults instantiates a new V1Time object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNanos

`func (o *V1Time) GetNanos() int32`

GetNanos returns the Nanos field if non-nil, zero value otherwise.

### GetNanosOk

`func (o *V1Time) GetNanosOk() (*int32, bool)`

GetNanosOk returns a tuple with the Nanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNanos

`func (o *V1Time) SetNanos(v int32)`

SetNanos sets Nanos field to given value.

### HasNanos

`func (o *V1Time) HasNanos() bool`

HasNanos returns a boolean if a field has been set.

### GetSeconds

`func (o *V1Time) GetSeconds() string`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *V1Time) GetSecondsOk() (*string, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *V1Time) SetSeconds(v string)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *V1Time) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


