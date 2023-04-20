# V1Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**EventTime** | Pointer to [**V1MicroTime**](V1MicroTime.md) |  | [optional] 
**FirstTimestamp** | Pointer to **string** |  | [optional] 
**InvolvedObject** | Pointer to [**V1ObjectReference**](V1ObjectReference.md) |  | [optional] 
**LastTimestamp** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Related** | Pointer to [**V1ObjectReference**](V1ObjectReference.md) |  | [optional] 
**ReportingComponent** | Pointer to **string** |  | [optional] 
**ReportingInstance** | Pointer to **string** |  | [optional] 
**Series** | Pointer to [**V1EventSeries**](V1EventSeries.md) |  | [optional] 
**Source** | Pointer to [**V1EventSource**](V1EventSource.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewV1Event

`func NewV1Event() *V1Event`

NewV1Event instantiates a new V1Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EventWithDefaults

`func NewV1EventWithDefaults() *V1Event`

NewV1EventWithDefaults instantiates a new V1Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *V1Event) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V1Event) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V1Event) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *V1Event) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCount

`func (o *V1Event) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V1Event) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V1Event) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V1Event) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetEventTime

`func (o *V1Event) GetEventTime() V1MicroTime`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *V1Event) GetEventTimeOk() (*V1MicroTime, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *V1Event) SetEventTime(v V1MicroTime)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *V1Event) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetFirstTimestamp

`func (o *V1Event) GetFirstTimestamp() string`

GetFirstTimestamp returns the FirstTimestamp field if non-nil, zero value otherwise.

### GetFirstTimestampOk

`func (o *V1Event) GetFirstTimestampOk() (*string, bool)`

GetFirstTimestampOk returns a tuple with the FirstTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTimestamp

`func (o *V1Event) SetFirstTimestamp(v string)`

SetFirstTimestamp sets FirstTimestamp field to given value.

### HasFirstTimestamp

`func (o *V1Event) HasFirstTimestamp() bool`

HasFirstTimestamp returns a boolean if a field has been set.

### GetInvolvedObject

`func (o *V1Event) GetInvolvedObject() V1ObjectReference`

GetInvolvedObject returns the InvolvedObject field if non-nil, zero value otherwise.

### GetInvolvedObjectOk

`func (o *V1Event) GetInvolvedObjectOk() (*V1ObjectReference, bool)`

GetInvolvedObjectOk returns a tuple with the InvolvedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvolvedObject

`func (o *V1Event) SetInvolvedObject(v V1ObjectReference)`

SetInvolvedObject sets InvolvedObject field to given value.

### HasInvolvedObject

`func (o *V1Event) HasInvolvedObject() bool`

HasInvolvedObject returns a boolean if a field has been set.

### GetLastTimestamp

`func (o *V1Event) GetLastTimestamp() string`

GetLastTimestamp returns the LastTimestamp field if non-nil, zero value otherwise.

### GetLastTimestampOk

`func (o *V1Event) GetLastTimestampOk() (*string, bool)`

GetLastTimestampOk returns a tuple with the LastTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimestamp

`func (o *V1Event) SetLastTimestamp(v string)`

SetLastTimestamp sets LastTimestamp field to given value.

### HasLastTimestamp

`func (o *V1Event) HasLastTimestamp() bool`

HasLastTimestamp returns a boolean if a field has been set.

### GetMessage

`func (o *V1Event) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1Event) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1Event) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1Event) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *V1Event) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1Event) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1Event) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1Event) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReason

`func (o *V1Event) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1Event) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1Event) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1Event) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRelated

`func (o *V1Event) GetRelated() V1ObjectReference`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *V1Event) GetRelatedOk() (*V1ObjectReference, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *V1Event) SetRelated(v V1ObjectReference)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *V1Event) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetReportingComponent

`func (o *V1Event) GetReportingComponent() string`

GetReportingComponent returns the ReportingComponent field if non-nil, zero value otherwise.

### GetReportingComponentOk

`func (o *V1Event) GetReportingComponentOk() (*string, bool)`

GetReportingComponentOk returns a tuple with the ReportingComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingComponent

`func (o *V1Event) SetReportingComponent(v string)`

SetReportingComponent sets ReportingComponent field to given value.

### HasReportingComponent

`func (o *V1Event) HasReportingComponent() bool`

HasReportingComponent returns a boolean if a field has been set.

### GetReportingInstance

`func (o *V1Event) GetReportingInstance() string`

GetReportingInstance returns the ReportingInstance field if non-nil, zero value otherwise.

### GetReportingInstanceOk

`func (o *V1Event) GetReportingInstanceOk() (*string, bool)`

GetReportingInstanceOk returns a tuple with the ReportingInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingInstance

`func (o *V1Event) SetReportingInstance(v string)`

SetReportingInstance sets ReportingInstance field to given value.

### HasReportingInstance

`func (o *V1Event) HasReportingInstance() bool`

HasReportingInstance returns a boolean if a field has been set.

### GetSeries

`func (o *V1Event) GetSeries() V1EventSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *V1Event) GetSeriesOk() (*V1EventSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *V1Event) SetSeries(v V1EventSeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *V1Event) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetSource

`func (o *V1Event) GetSource() V1EventSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1Event) GetSourceOk() (*V1EventSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1Event) SetSource(v V1EventSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1Event) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *V1Event) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1Event) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1Event) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1Event) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


