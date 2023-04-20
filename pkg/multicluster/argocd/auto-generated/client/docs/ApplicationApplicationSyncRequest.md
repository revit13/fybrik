# ApplicationApplicationSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppNamespace** | Pointer to **string** |  | [optional] 
**DryRun** | Pointer to **bool** |  | [optional] 
**Infos** | Pointer to [**[]V1alpha1Info**](V1alpha1Info.md) |  | [optional] 
**Manifests** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Prune** | Pointer to **bool** |  | [optional] 
**Resources** | Pointer to [**[]V1alpha1SyncOperationResource**](V1alpha1SyncOperationResource.md) |  | [optional] 
**RetryStrategy** | Pointer to [**V1alpha1RetryStrategy**](V1alpha1RetryStrategy.md) |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Strategy** | Pointer to [**V1alpha1SyncStrategy**](V1alpha1SyncStrategy.md) |  | [optional] 
**SyncOptions** | Pointer to [**ApplicationSyncOptions**](ApplicationSyncOptions.md) |  | [optional] 

## Methods

### NewApplicationApplicationSyncRequest

`func NewApplicationApplicationSyncRequest() *ApplicationApplicationSyncRequest`

NewApplicationApplicationSyncRequest instantiates a new ApplicationApplicationSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationApplicationSyncRequestWithDefaults

`func NewApplicationApplicationSyncRequestWithDefaults() *ApplicationApplicationSyncRequest`

NewApplicationApplicationSyncRequestWithDefaults instantiates a new ApplicationApplicationSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppNamespace

`func (o *ApplicationApplicationSyncRequest) GetAppNamespace() string`

GetAppNamespace returns the AppNamespace field if non-nil, zero value otherwise.

### GetAppNamespaceOk

`func (o *ApplicationApplicationSyncRequest) GetAppNamespaceOk() (*string, bool)`

GetAppNamespaceOk returns a tuple with the AppNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppNamespace

`func (o *ApplicationApplicationSyncRequest) SetAppNamespace(v string)`

SetAppNamespace sets AppNamespace field to given value.

### HasAppNamespace

`func (o *ApplicationApplicationSyncRequest) HasAppNamespace() bool`

HasAppNamespace returns a boolean if a field has been set.

### GetDryRun

`func (o *ApplicationApplicationSyncRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ApplicationApplicationSyncRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ApplicationApplicationSyncRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ApplicationApplicationSyncRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetInfos

`func (o *ApplicationApplicationSyncRequest) GetInfos() []V1alpha1Info`

GetInfos returns the Infos field if non-nil, zero value otherwise.

### GetInfosOk

`func (o *ApplicationApplicationSyncRequest) GetInfosOk() (*[]V1alpha1Info, bool)`

GetInfosOk returns a tuple with the Infos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfos

`func (o *ApplicationApplicationSyncRequest) SetInfos(v []V1alpha1Info)`

SetInfos sets Infos field to given value.

### HasInfos

`func (o *ApplicationApplicationSyncRequest) HasInfos() bool`

HasInfos returns a boolean if a field has been set.

### GetManifests

`func (o *ApplicationApplicationSyncRequest) GetManifests() []string`

GetManifests returns the Manifests field if non-nil, zero value otherwise.

### GetManifestsOk

`func (o *ApplicationApplicationSyncRequest) GetManifestsOk() (*[]string, bool)`

GetManifestsOk returns a tuple with the Manifests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifests

`func (o *ApplicationApplicationSyncRequest) SetManifests(v []string)`

SetManifests sets Manifests field to given value.

### HasManifests

`func (o *ApplicationApplicationSyncRequest) HasManifests() bool`

HasManifests returns a boolean if a field has been set.

### GetName

`func (o *ApplicationApplicationSyncRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationApplicationSyncRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationApplicationSyncRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationApplicationSyncRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrune

`func (o *ApplicationApplicationSyncRequest) GetPrune() bool`

GetPrune returns the Prune field if non-nil, zero value otherwise.

### GetPruneOk

`func (o *ApplicationApplicationSyncRequest) GetPruneOk() (*bool, bool)`

GetPruneOk returns a tuple with the Prune field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrune

`func (o *ApplicationApplicationSyncRequest) SetPrune(v bool)`

SetPrune sets Prune field to given value.

### HasPrune

`func (o *ApplicationApplicationSyncRequest) HasPrune() bool`

HasPrune returns a boolean if a field has been set.

### GetResources

`func (o *ApplicationApplicationSyncRequest) GetResources() []V1alpha1SyncOperationResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ApplicationApplicationSyncRequest) GetResourcesOk() (*[]V1alpha1SyncOperationResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ApplicationApplicationSyncRequest) SetResources(v []V1alpha1SyncOperationResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ApplicationApplicationSyncRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRetryStrategy

`func (o *ApplicationApplicationSyncRequest) GetRetryStrategy() V1alpha1RetryStrategy`

GetRetryStrategy returns the RetryStrategy field if non-nil, zero value otherwise.

### GetRetryStrategyOk

`func (o *ApplicationApplicationSyncRequest) GetRetryStrategyOk() (*V1alpha1RetryStrategy, bool)`

GetRetryStrategyOk returns a tuple with the RetryStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryStrategy

`func (o *ApplicationApplicationSyncRequest) SetRetryStrategy(v V1alpha1RetryStrategy)`

SetRetryStrategy sets RetryStrategy field to given value.

### HasRetryStrategy

`func (o *ApplicationApplicationSyncRequest) HasRetryStrategy() bool`

HasRetryStrategy returns a boolean if a field has been set.

### GetRevision

`func (o *ApplicationApplicationSyncRequest) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ApplicationApplicationSyncRequest) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ApplicationApplicationSyncRequest) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ApplicationApplicationSyncRequest) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetStrategy

`func (o *ApplicationApplicationSyncRequest) GetStrategy() V1alpha1SyncStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *ApplicationApplicationSyncRequest) GetStrategyOk() (*V1alpha1SyncStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *ApplicationApplicationSyncRequest) SetStrategy(v V1alpha1SyncStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *ApplicationApplicationSyncRequest) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetSyncOptions

`func (o *ApplicationApplicationSyncRequest) GetSyncOptions() ApplicationSyncOptions`

GetSyncOptions returns the SyncOptions field if non-nil, zero value otherwise.

### GetSyncOptionsOk

`func (o *ApplicationApplicationSyncRequest) GetSyncOptionsOk() (*ApplicationSyncOptions, bool)`

GetSyncOptionsOk returns a tuple with the SyncOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncOptions

`func (o *ApplicationApplicationSyncRequest) SetSyncOptions(v ApplicationSyncOptions)`

SetSyncOptions sets SyncOptions field to given value.

### HasSyncOptions

`func (o *ApplicationApplicationSyncRequest) HasSyncOptions() bool`

HasSyncOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


