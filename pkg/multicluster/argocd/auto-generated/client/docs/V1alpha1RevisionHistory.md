# V1alpha1RevisionHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeployStartedAt** | Pointer to **string** |  | [optional] 
**DeployedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Revisions** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to [**V1alpha1ApplicationSource**](V1alpha1ApplicationSource.md) |  | [optional] 
**Sources** | Pointer to [**[]V1alpha1ApplicationSource**](V1alpha1ApplicationSource.md) |  | [optional] 

## Methods

### NewV1alpha1RevisionHistory

`func NewV1alpha1RevisionHistory() *V1alpha1RevisionHistory`

NewV1alpha1RevisionHistory instantiates a new V1alpha1RevisionHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1RevisionHistoryWithDefaults

`func NewV1alpha1RevisionHistoryWithDefaults() *V1alpha1RevisionHistory`

NewV1alpha1RevisionHistoryWithDefaults instantiates a new V1alpha1RevisionHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployStartedAt

`func (o *V1alpha1RevisionHistory) GetDeployStartedAt() string`

GetDeployStartedAt returns the DeployStartedAt field if non-nil, zero value otherwise.

### GetDeployStartedAtOk

`func (o *V1alpha1RevisionHistory) GetDeployStartedAtOk() (*string, bool)`

GetDeployStartedAtOk returns a tuple with the DeployStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStartedAt

`func (o *V1alpha1RevisionHistory) SetDeployStartedAt(v string)`

SetDeployStartedAt sets DeployStartedAt field to given value.

### HasDeployStartedAt

`func (o *V1alpha1RevisionHistory) HasDeployStartedAt() bool`

HasDeployStartedAt returns a boolean if a field has been set.

### GetDeployedAt

`func (o *V1alpha1RevisionHistory) GetDeployedAt() string`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *V1alpha1RevisionHistory) GetDeployedAtOk() (*string, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *V1alpha1RevisionHistory) SetDeployedAt(v string)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *V1alpha1RevisionHistory) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### GetId

`func (o *V1alpha1RevisionHistory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1alpha1RevisionHistory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1alpha1RevisionHistory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V1alpha1RevisionHistory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRevision

`func (o *V1alpha1RevisionHistory) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1alpha1RevisionHistory) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1alpha1RevisionHistory) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1alpha1RevisionHistory) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRevisions

`func (o *V1alpha1RevisionHistory) GetRevisions() []string`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *V1alpha1RevisionHistory) GetRevisionsOk() (*[]string, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *V1alpha1RevisionHistory) SetRevisions(v []string)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *V1alpha1RevisionHistory) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### GetSource

`func (o *V1alpha1RevisionHistory) GetSource() V1alpha1ApplicationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1alpha1RevisionHistory) GetSourceOk() (*V1alpha1ApplicationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1alpha1RevisionHistory) SetSource(v V1alpha1ApplicationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1alpha1RevisionHistory) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSources

`func (o *V1alpha1RevisionHistory) GetSources() []V1alpha1ApplicationSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *V1alpha1RevisionHistory) GetSourcesOk() (*[]V1alpha1ApplicationSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *V1alpha1RevisionHistory) SetSources(v []V1alpha1ApplicationSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *V1alpha1RevisionHistory) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


