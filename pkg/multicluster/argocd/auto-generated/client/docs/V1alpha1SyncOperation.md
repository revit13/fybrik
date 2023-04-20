# V1alpha1SyncOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** |  | [optional] 
**Manifests** | Pointer to **[]string** |  | [optional] 
**Prune** | Pointer to **bool** |  | [optional] 
**Resources** | Pointer to [**[]V1alpha1SyncOperationResource**](V1alpha1SyncOperationResource.md) |  | [optional] 
**Revision** | Pointer to **string** | Revision is the revision (Git) or chart version (Helm) which to sync the application to If omitted, will use the revision specified in app spec. | [optional] 
**Revisions** | Pointer to **[]string** | Revisions is the list of revision (Git) or chart version (Helm) which to sync each source in sources field for the application to If omitted, will use the revision specified in app spec. | [optional] 
**Source** | Pointer to [**V1alpha1ApplicationSource**](V1alpha1ApplicationSource.md) |  | [optional] 
**Sources** | Pointer to [**[]V1alpha1ApplicationSource**](V1alpha1ApplicationSource.md) |  | [optional] 
**SyncOptions** | Pointer to **[]string** |  | [optional] 
**SyncStrategy** | Pointer to [**V1alpha1SyncStrategy**](V1alpha1SyncStrategy.md) |  | [optional] 

## Methods

### NewV1alpha1SyncOperation

`func NewV1alpha1SyncOperation() *V1alpha1SyncOperation`

NewV1alpha1SyncOperation instantiates a new V1alpha1SyncOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1SyncOperationWithDefaults

`func NewV1alpha1SyncOperationWithDefaults() *V1alpha1SyncOperation`

NewV1alpha1SyncOperationWithDefaults instantiates a new V1alpha1SyncOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *V1alpha1SyncOperation) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *V1alpha1SyncOperation) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *V1alpha1SyncOperation) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *V1alpha1SyncOperation) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetManifests

`func (o *V1alpha1SyncOperation) GetManifests() []string`

GetManifests returns the Manifests field if non-nil, zero value otherwise.

### GetManifestsOk

`func (o *V1alpha1SyncOperation) GetManifestsOk() (*[]string, bool)`

GetManifestsOk returns a tuple with the Manifests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifests

`func (o *V1alpha1SyncOperation) SetManifests(v []string)`

SetManifests sets Manifests field to given value.

### HasManifests

`func (o *V1alpha1SyncOperation) HasManifests() bool`

HasManifests returns a boolean if a field has been set.

### GetPrune

`func (o *V1alpha1SyncOperation) GetPrune() bool`

GetPrune returns the Prune field if non-nil, zero value otherwise.

### GetPruneOk

`func (o *V1alpha1SyncOperation) GetPruneOk() (*bool, bool)`

GetPruneOk returns a tuple with the Prune field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrune

`func (o *V1alpha1SyncOperation) SetPrune(v bool)`

SetPrune sets Prune field to given value.

### HasPrune

`func (o *V1alpha1SyncOperation) HasPrune() bool`

HasPrune returns a boolean if a field has been set.

### GetResources

`func (o *V1alpha1SyncOperation) GetResources() []V1alpha1SyncOperationResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V1alpha1SyncOperation) GetResourcesOk() (*[]V1alpha1SyncOperationResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V1alpha1SyncOperation) SetResources(v []V1alpha1SyncOperationResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V1alpha1SyncOperation) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRevision

`func (o *V1alpha1SyncOperation) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1alpha1SyncOperation) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1alpha1SyncOperation) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1alpha1SyncOperation) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRevisions

`func (o *V1alpha1SyncOperation) GetRevisions() []string`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *V1alpha1SyncOperation) GetRevisionsOk() (*[]string, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *V1alpha1SyncOperation) SetRevisions(v []string)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *V1alpha1SyncOperation) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### GetSource

`func (o *V1alpha1SyncOperation) GetSource() V1alpha1ApplicationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1alpha1SyncOperation) GetSourceOk() (*V1alpha1ApplicationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1alpha1SyncOperation) SetSource(v V1alpha1ApplicationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1alpha1SyncOperation) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSources

`func (o *V1alpha1SyncOperation) GetSources() []V1alpha1ApplicationSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *V1alpha1SyncOperation) GetSourcesOk() (*[]V1alpha1ApplicationSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *V1alpha1SyncOperation) SetSources(v []V1alpha1ApplicationSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *V1alpha1SyncOperation) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSyncOptions

`func (o *V1alpha1SyncOperation) GetSyncOptions() []string`

GetSyncOptions returns the SyncOptions field if non-nil, zero value otherwise.

### GetSyncOptionsOk

`func (o *V1alpha1SyncOperation) GetSyncOptionsOk() (*[]string, bool)`

GetSyncOptionsOk returns a tuple with the SyncOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncOptions

`func (o *V1alpha1SyncOperation) SetSyncOptions(v []string)`

SetSyncOptions sets SyncOptions field to given value.

### HasSyncOptions

`func (o *V1alpha1SyncOperation) HasSyncOptions() bool`

HasSyncOptions returns a boolean if a field has been set.

### GetSyncStrategy

`func (o *V1alpha1SyncOperation) GetSyncStrategy() V1alpha1SyncStrategy`

GetSyncStrategy returns the SyncStrategy field if non-nil, zero value otherwise.

### GetSyncStrategyOk

`func (o *V1alpha1SyncOperation) GetSyncStrategyOk() (*V1alpha1SyncStrategy, bool)`

GetSyncStrategyOk returns a tuple with the SyncStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStrategy

`func (o *V1alpha1SyncOperation) SetSyncStrategy(v V1alpha1SyncStrategy)`

SetSyncStrategy sets SyncStrategy field to given value.

### HasSyncStrategy

`func (o *V1alpha1SyncOperation) HasSyncStrategy() bool`

HasSyncStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


