# V1alpha1SyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparedTo** | Pointer to [**V1alpha1ComparedTo**](V1alpha1ComparedTo.md) |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Revisions** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1SyncStatus

`func NewV1alpha1SyncStatus() *V1alpha1SyncStatus`

NewV1alpha1SyncStatus instantiates a new V1alpha1SyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1SyncStatusWithDefaults

`func NewV1alpha1SyncStatusWithDefaults() *V1alpha1SyncStatus`

NewV1alpha1SyncStatusWithDefaults instantiates a new V1alpha1SyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparedTo

`func (o *V1alpha1SyncStatus) GetComparedTo() V1alpha1ComparedTo`

GetComparedTo returns the ComparedTo field if non-nil, zero value otherwise.

### GetComparedToOk

`func (o *V1alpha1SyncStatus) GetComparedToOk() (*V1alpha1ComparedTo, bool)`

GetComparedToOk returns a tuple with the ComparedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparedTo

`func (o *V1alpha1SyncStatus) SetComparedTo(v V1alpha1ComparedTo)`

SetComparedTo sets ComparedTo field to given value.

### HasComparedTo

`func (o *V1alpha1SyncStatus) HasComparedTo() bool`

HasComparedTo returns a boolean if a field has been set.

### GetRevision

`func (o *V1alpha1SyncStatus) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1alpha1SyncStatus) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1alpha1SyncStatus) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1alpha1SyncStatus) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRevisions

`func (o *V1alpha1SyncStatus) GetRevisions() []string`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *V1alpha1SyncStatus) GetRevisionsOk() (*[]string, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *V1alpha1SyncStatus) SetRevisions(v []string)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *V1alpha1SyncStatus) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### GetStatus

`func (o *V1alpha1SyncStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1alpha1SyncStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1alpha1SyncStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1alpha1SyncStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


