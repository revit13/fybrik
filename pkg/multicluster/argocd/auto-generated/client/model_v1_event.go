/*
Consolidate Services

Description of all APIs

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
)

// V1Event Event is a report of an event somewhere in the cluster.  Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
type V1Event struct {
	Action             *string            `json:"action,omitempty"`
	Count              *int32             `json:"count,omitempty"`
	EventTime          *V1MicroTime       `json:"eventTime,omitempty"`
	FirstTimestamp     *string            `json:"firstTimestamp,omitempty"`
	InvolvedObject     *V1ObjectReference `json:"involvedObject,omitempty"`
	LastTimestamp      *string            `json:"lastTimestamp,omitempty"`
	Message            *string            `json:"message,omitempty"`
	Metadata           *V1ObjectMeta      `json:"metadata,omitempty"`
	Reason             *string            `json:"reason,omitempty"`
	Related            *V1ObjectReference `json:"related,omitempty"`
	ReportingComponent *string            `json:"reportingComponent,omitempty"`
	ReportingInstance  *string            `json:"reportingInstance,omitempty"`
	Series             *V1EventSeries     `json:"series,omitempty"`
	Source             *V1EventSource     `json:"source,omitempty"`
	Type               *string            `json:"type,omitempty"`
}

// NewV1Event instantiates a new V1Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Event() *V1Event {
	this := V1Event{}
	return &this
}

// NewV1EventWithDefaults instantiates a new V1Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1EventWithDefaults() *V1Event {
	this := V1Event{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *V1Event) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *V1Event) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *V1Event) SetAction(v string) {
	o.Action = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V1Event) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V1Event) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *V1Event) SetCount(v int32) {
	o.Count = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *V1Event) GetEventTime() V1MicroTime {
	if o == nil || o.EventTime == nil {
		var ret V1MicroTime
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetEventTimeOk() (*V1MicroTime, bool) {
	if o == nil || o.EventTime == nil {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *V1Event) HasEventTime() bool {
	if o != nil && o.EventTime != nil {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given V1MicroTime and assigns it to the EventTime field.
func (o *V1Event) SetEventTime(v V1MicroTime) {
	o.EventTime = &v
}

// GetFirstTimestamp returns the FirstTimestamp field value if set, zero value otherwise.
func (o *V1Event) GetFirstTimestamp() string {
	if o == nil || o.FirstTimestamp == nil {
		var ret string
		return ret
	}
	return *o.FirstTimestamp
}

// GetFirstTimestampOk returns a tuple with the FirstTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetFirstTimestampOk() (*string, bool) {
	if o == nil || o.FirstTimestamp == nil {
		return nil, false
	}
	return o.FirstTimestamp, true
}

// HasFirstTimestamp returns a boolean if a field has been set.
func (o *V1Event) HasFirstTimestamp() bool {
	if o != nil && o.FirstTimestamp != nil {
		return true
	}

	return false
}

// SetFirstTimestamp gets a reference to the given string and assigns it to the FirstTimestamp field.
func (o *V1Event) SetFirstTimestamp(v string) {
	o.FirstTimestamp = &v
}

// GetInvolvedObject returns the InvolvedObject field value if set, zero value otherwise.
func (o *V1Event) GetInvolvedObject() V1ObjectReference {
	if o == nil || o.InvolvedObject == nil {
		var ret V1ObjectReference
		return ret
	}
	return *o.InvolvedObject
}

// GetInvolvedObjectOk returns a tuple with the InvolvedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetInvolvedObjectOk() (*V1ObjectReference, bool) {
	if o == nil || o.InvolvedObject == nil {
		return nil, false
	}
	return o.InvolvedObject, true
}

// HasInvolvedObject returns a boolean if a field has been set.
func (o *V1Event) HasInvolvedObject() bool {
	if o != nil && o.InvolvedObject != nil {
		return true
	}

	return false
}

// SetInvolvedObject gets a reference to the given V1ObjectReference and assigns it to the InvolvedObject field.
func (o *V1Event) SetInvolvedObject(v V1ObjectReference) {
	o.InvolvedObject = &v
}

// GetLastTimestamp returns the LastTimestamp field value if set, zero value otherwise.
func (o *V1Event) GetLastTimestamp() string {
	if o == nil || o.LastTimestamp == nil {
		var ret string
		return ret
	}
	return *o.LastTimestamp
}

// GetLastTimestampOk returns a tuple with the LastTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetLastTimestampOk() (*string, bool) {
	if o == nil || o.LastTimestamp == nil {
		return nil, false
	}
	return o.LastTimestamp, true
}

// HasLastTimestamp returns a boolean if a field has been set.
func (o *V1Event) HasLastTimestamp() bool {
	if o != nil && o.LastTimestamp != nil {
		return true
	}

	return false
}

// SetLastTimestamp gets a reference to the given string and assigns it to the LastTimestamp field.
func (o *V1Event) SetLastTimestamp(v string) {
	o.LastTimestamp = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1Event) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1Event) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1Event) SetMessage(v string) {
	o.Message = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1Event) GetMetadata() V1ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret V1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetMetadataOk() (*V1ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1Event) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ObjectMeta and assigns it to the Metadata field.
func (o *V1Event) SetMetadata(v V1ObjectMeta) {
	o.Metadata = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *V1Event) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *V1Event) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *V1Event) SetReason(v string) {
	o.Reason = &v
}

// GetRelated returns the Related field value if set, zero value otherwise.
func (o *V1Event) GetRelated() V1ObjectReference {
	if o == nil || o.Related == nil {
		var ret V1ObjectReference
		return ret
	}
	return *o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetRelatedOk() (*V1ObjectReference, bool) {
	if o == nil || o.Related == nil {
		return nil, false
	}
	return o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *V1Event) HasRelated() bool {
	if o != nil && o.Related != nil {
		return true
	}

	return false
}

// SetRelated gets a reference to the given V1ObjectReference and assigns it to the Related field.
func (o *V1Event) SetRelated(v V1ObjectReference) {
	o.Related = &v
}

// GetReportingComponent returns the ReportingComponent field value if set, zero value otherwise.
func (o *V1Event) GetReportingComponent() string {
	if o == nil || o.ReportingComponent == nil {
		var ret string
		return ret
	}
	return *o.ReportingComponent
}

// GetReportingComponentOk returns a tuple with the ReportingComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetReportingComponentOk() (*string, bool) {
	if o == nil || o.ReportingComponent == nil {
		return nil, false
	}
	return o.ReportingComponent, true
}

// HasReportingComponent returns a boolean if a field has been set.
func (o *V1Event) HasReportingComponent() bool {
	if o != nil && o.ReportingComponent != nil {
		return true
	}

	return false
}

// SetReportingComponent gets a reference to the given string and assigns it to the ReportingComponent field.
func (o *V1Event) SetReportingComponent(v string) {
	o.ReportingComponent = &v
}

// GetReportingInstance returns the ReportingInstance field value if set, zero value otherwise.
func (o *V1Event) GetReportingInstance() string {
	if o == nil || o.ReportingInstance == nil {
		var ret string
		return ret
	}
	return *o.ReportingInstance
}

// GetReportingInstanceOk returns a tuple with the ReportingInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetReportingInstanceOk() (*string, bool) {
	if o == nil || o.ReportingInstance == nil {
		return nil, false
	}
	return o.ReportingInstance, true
}

// HasReportingInstance returns a boolean if a field has been set.
func (o *V1Event) HasReportingInstance() bool {
	if o != nil && o.ReportingInstance != nil {
		return true
	}

	return false
}

// SetReportingInstance gets a reference to the given string and assigns it to the ReportingInstance field.
func (o *V1Event) SetReportingInstance(v string) {
	o.ReportingInstance = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *V1Event) GetSeries() V1EventSeries {
	if o == nil || o.Series == nil {
		var ret V1EventSeries
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetSeriesOk() (*V1EventSeries, bool) {
	if o == nil || o.Series == nil {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *V1Event) HasSeries() bool {
	if o != nil && o.Series != nil {
		return true
	}

	return false
}

// SetSeries gets a reference to the given V1EventSeries and assigns it to the Series field.
func (o *V1Event) SetSeries(v V1EventSeries) {
	o.Series = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *V1Event) GetSource() V1EventSource {
	if o == nil || o.Source == nil {
		var ret V1EventSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetSourceOk() (*V1EventSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *V1Event) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given V1EventSource and assigns it to the Source field.
func (o *V1Event) SetSource(v V1EventSource) {
	o.Source = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1Event) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1Event) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1Event) SetType(v string) {
	o.Type = &v
}

func (o V1Event) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.EventTime != nil {
		toSerialize["eventTime"] = o.EventTime
	}
	if o.FirstTimestamp != nil {
		toSerialize["firstTimestamp"] = o.FirstTimestamp
	}
	if o.InvolvedObject != nil {
		toSerialize["involvedObject"] = o.InvolvedObject
	}
	if o.LastTimestamp != nil {
		toSerialize["lastTimestamp"] = o.LastTimestamp
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Related != nil {
		toSerialize["related"] = o.Related
	}
	if o.ReportingComponent != nil {
		toSerialize["reportingComponent"] = o.ReportingComponent
	}
	if o.ReportingInstance != nil {
		toSerialize["reportingInstance"] = o.ReportingInstance
	}
	if o.Series != nil {
		toSerialize["series"] = o.Series
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableV1Event struct {
	value *V1Event
	isSet bool
}

func (v NullableV1Event) Get() *V1Event {
	return v.value
}

func (v *NullableV1Event) Set(val *V1Event) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Event) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Event) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Event(val *V1Event) *NullableV1Event {
	return &NullableV1Event{value: val, isSet: true}
}

func (v NullableV1Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Event) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}