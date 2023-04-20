# ApplicationLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**Last** | Pointer to **bool** |  | [optional] 
**PodName** | Pointer to **string** |  | [optional] 
**TimeStamp** | Pointer to [**V1Time**](V1Time.md) |  | [optional] 
**TimeStampStr** | Pointer to **string** |  | [optional] 

## Methods

### NewApplicationLogEntry

`func NewApplicationLogEntry() *ApplicationLogEntry`

NewApplicationLogEntry instantiates a new ApplicationLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationLogEntryWithDefaults

`func NewApplicationLogEntryWithDefaults() *ApplicationLogEntry`

NewApplicationLogEntryWithDefaults instantiates a new ApplicationLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ApplicationLogEntry) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ApplicationLogEntry) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ApplicationLogEntry) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ApplicationLogEntry) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetLast

`func (o *ApplicationLogEntry) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *ApplicationLogEntry) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *ApplicationLogEntry) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *ApplicationLogEntry) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetPodName

`func (o *ApplicationLogEntry) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *ApplicationLogEntry) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *ApplicationLogEntry) SetPodName(v string)`

SetPodName sets PodName field to given value.

### HasPodName

`func (o *ApplicationLogEntry) HasPodName() bool`

HasPodName returns a boolean if a field has been set.

### GetTimeStamp

`func (o *ApplicationLogEntry) GetTimeStamp() V1Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *ApplicationLogEntry) GetTimeStampOk() (*V1Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *ApplicationLogEntry) SetTimeStamp(v V1Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *ApplicationLogEntry) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetTimeStampStr

`func (o *ApplicationLogEntry) GetTimeStampStr() string`

GetTimeStampStr returns the TimeStampStr field if non-nil, zero value otherwise.

### GetTimeStampStrOk

`func (o *ApplicationLogEntry) GetTimeStampStrOk() (*string, bool)`

GetTimeStampStrOk returns a tuple with the TimeStampStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStampStr

`func (o *ApplicationLogEntry) SetTimeStampStr(v string)`

SetTimeStampStr sets TimeStampStr field to given value.

### HasTimeStampStr

`func (o *ApplicationLogEntry) HasTimeStampStr() bool`

HasTimeStampStr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


