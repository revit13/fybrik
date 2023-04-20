# V1EventSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**LastObservedTime** | Pointer to [**V1MicroTime**](V1MicroTime.md) |  | [optional] 

## Methods

### NewV1EventSeries

`func NewV1EventSeries() *V1EventSeries`

NewV1EventSeries instantiates a new V1EventSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EventSeriesWithDefaults

`func NewV1EventSeriesWithDefaults() *V1EventSeries`

NewV1EventSeriesWithDefaults instantiates a new V1EventSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *V1EventSeries) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V1EventSeries) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V1EventSeries) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V1EventSeries) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLastObservedTime

`func (o *V1EventSeries) GetLastObservedTime() V1MicroTime`

GetLastObservedTime returns the LastObservedTime field if non-nil, zero value otherwise.

### GetLastObservedTimeOk

`func (o *V1EventSeries) GetLastObservedTimeOk() (*V1MicroTime, bool)`

GetLastObservedTimeOk returns a tuple with the LastObservedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastObservedTime

`func (o *V1EventSeries) SetLastObservedTime(v V1MicroTime)`

SetLastObservedTime sets LastObservedTime field to given value.

### HasLastObservedTime

`func (o *V1EventSeries) HasLastObservedTime() bool`

HasLastObservedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


