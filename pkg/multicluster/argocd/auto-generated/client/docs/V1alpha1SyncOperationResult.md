# V1alpha1SyncOperationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]V1alpha1ResourceResult**](V1alpha1ResourceResult.md) |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Revisions** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to [**V1alpha1ApplicationSource**](V1alpha1ApplicationSource.md) |  | [optional] 
**Sources** | Pointer to [**[]V1alpha1ApplicationSource**](V1alpha1ApplicationSource.md) |  | [optional] 

## Methods

### NewV1alpha1SyncOperationResult

`func NewV1alpha1SyncOperationResult() *V1alpha1SyncOperationResult`

NewV1alpha1SyncOperationResult instantiates a new V1alpha1SyncOperationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1SyncOperationResultWithDefaults

`func NewV1alpha1SyncOperationResultWithDefaults() *V1alpha1SyncOperationResult`

NewV1alpha1SyncOperationResultWithDefaults instantiates a new V1alpha1SyncOperationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *V1alpha1SyncOperationResult) GetResources() []V1alpha1ResourceResult`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V1alpha1SyncOperationResult) GetResourcesOk() (*[]V1alpha1ResourceResult, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V1alpha1SyncOperationResult) SetResources(v []V1alpha1ResourceResult)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V1alpha1SyncOperationResult) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRevision

`func (o *V1alpha1SyncOperationResult) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1alpha1SyncOperationResult) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1alpha1SyncOperationResult) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1alpha1SyncOperationResult) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRevisions

`func (o *V1alpha1SyncOperationResult) GetRevisions() []string`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *V1alpha1SyncOperationResult) GetRevisionsOk() (*[]string, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *V1alpha1SyncOperationResult) SetRevisions(v []string)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *V1alpha1SyncOperationResult) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### GetSource

`func (o *V1alpha1SyncOperationResult) GetSource() V1alpha1ApplicationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1alpha1SyncOperationResult) GetSourceOk() (*V1alpha1ApplicationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1alpha1SyncOperationResult) SetSource(v V1alpha1ApplicationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1alpha1SyncOperationResult) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSources

`func (o *V1alpha1SyncOperationResult) GetSources() []V1alpha1ApplicationSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *V1alpha1SyncOperationResult) GetSourcesOk() (*[]V1alpha1ApplicationSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *V1alpha1SyncOperationResult) SetSources(v []V1alpha1ApplicationSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *V1alpha1SyncOperationResult) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


