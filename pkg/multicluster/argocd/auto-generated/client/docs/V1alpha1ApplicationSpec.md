# V1alpha1ApplicationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to [**V1alpha1ApplicationDestination**](V1alpha1ApplicationDestination.md) |  | [optional] 
**IgnoreDifferences** | Pointer to [**[]V1alpha1ResourceIgnoreDifferences**](V1alpha1ResourceIgnoreDifferences.md) |  | [optional] 
**Info** | Pointer to [**[]V1alpha1Info**](V1alpha1Info.md) |  | [optional] 
**Project** | Pointer to **string** | Project is a reference to the project this application belongs to. The empty string means that application belongs to the &#39;default&#39; project. | [optional] 
**RevisionHistoryLimit** | Pointer to **string** | RevisionHistoryLimit limits the number of items kept in the application&#39;s revision history, which is used for informational purposes as well as for rollbacks to previous versions. This should only be changed in exceptional circumstances. Setting to zero will store no history. This will reduce storage used. Increasing will increase the space used to store the history, so we do not recommend increasing it. Default is 10. | [optional] 
**Source** | Pointer to [**V1alpha1ApplicationSource**](V1alpha1ApplicationSource.md) |  | [optional] 
**Sources** | Pointer to [**[]V1alpha1ApplicationSource**](V1alpha1ApplicationSource.md) |  | [optional] 
**SyncPolicy** | Pointer to [**V1alpha1SyncPolicy**](V1alpha1SyncPolicy.md) |  | [optional] 

## Methods

### NewV1alpha1ApplicationSpec

`func NewV1alpha1ApplicationSpec() *V1alpha1ApplicationSpec`

NewV1alpha1ApplicationSpec instantiates a new V1alpha1ApplicationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSpecWithDefaults

`func NewV1alpha1ApplicationSpecWithDefaults() *V1alpha1ApplicationSpec`

NewV1alpha1ApplicationSpecWithDefaults instantiates a new V1alpha1ApplicationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *V1alpha1ApplicationSpec) GetDestination() V1alpha1ApplicationDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *V1alpha1ApplicationSpec) GetDestinationOk() (*V1alpha1ApplicationDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *V1alpha1ApplicationSpec) SetDestination(v V1alpha1ApplicationDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *V1alpha1ApplicationSpec) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetIgnoreDifferences

`func (o *V1alpha1ApplicationSpec) GetIgnoreDifferences() []V1alpha1ResourceIgnoreDifferences`

GetIgnoreDifferences returns the IgnoreDifferences field if non-nil, zero value otherwise.

### GetIgnoreDifferencesOk

`func (o *V1alpha1ApplicationSpec) GetIgnoreDifferencesOk() (*[]V1alpha1ResourceIgnoreDifferences, bool)`

GetIgnoreDifferencesOk returns a tuple with the IgnoreDifferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDifferences

`func (o *V1alpha1ApplicationSpec) SetIgnoreDifferences(v []V1alpha1ResourceIgnoreDifferences)`

SetIgnoreDifferences sets IgnoreDifferences field to given value.

### HasIgnoreDifferences

`func (o *V1alpha1ApplicationSpec) HasIgnoreDifferences() bool`

HasIgnoreDifferences returns a boolean if a field has been set.

### GetInfo

`func (o *V1alpha1ApplicationSpec) GetInfo() []V1alpha1Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *V1alpha1ApplicationSpec) GetInfoOk() (*[]V1alpha1Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *V1alpha1ApplicationSpec) SetInfo(v []V1alpha1Info)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *V1alpha1ApplicationSpec) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetProject

`func (o *V1alpha1ApplicationSpec) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *V1alpha1ApplicationSpec) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *V1alpha1ApplicationSpec) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *V1alpha1ApplicationSpec) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRevisionHistoryLimit

`func (o *V1alpha1ApplicationSpec) GetRevisionHistoryLimit() string`

GetRevisionHistoryLimit returns the RevisionHistoryLimit field if non-nil, zero value otherwise.

### GetRevisionHistoryLimitOk

`func (o *V1alpha1ApplicationSpec) GetRevisionHistoryLimitOk() (*string, bool)`

GetRevisionHistoryLimitOk returns a tuple with the RevisionHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionHistoryLimit

`func (o *V1alpha1ApplicationSpec) SetRevisionHistoryLimit(v string)`

SetRevisionHistoryLimit sets RevisionHistoryLimit field to given value.

### HasRevisionHistoryLimit

`func (o *V1alpha1ApplicationSpec) HasRevisionHistoryLimit() bool`

HasRevisionHistoryLimit returns a boolean if a field has been set.

### GetSource

`func (o *V1alpha1ApplicationSpec) GetSource() V1alpha1ApplicationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1alpha1ApplicationSpec) GetSourceOk() (*V1alpha1ApplicationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1alpha1ApplicationSpec) SetSource(v V1alpha1ApplicationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1alpha1ApplicationSpec) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSources

`func (o *V1alpha1ApplicationSpec) GetSources() []V1alpha1ApplicationSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *V1alpha1ApplicationSpec) GetSourcesOk() (*[]V1alpha1ApplicationSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *V1alpha1ApplicationSpec) SetSources(v []V1alpha1ApplicationSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *V1alpha1ApplicationSpec) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSyncPolicy

`func (o *V1alpha1ApplicationSpec) GetSyncPolicy() V1alpha1SyncPolicy`

GetSyncPolicy returns the SyncPolicy field if non-nil, zero value otherwise.

### GetSyncPolicyOk

`func (o *V1alpha1ApplicationSpec) GetSyncPolicyOk() (*V1alpha1SyncPolicy, bool)`

GetSyncPolicyOk returns a tuple with the SyncPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPolicy

`func (o *V1alpha1ApplicationSpec) SetSyncPolicy(v V1alpha1SyncPolicy)`

SetSyncPolicy sets SyncPolicy field to given value.

### HasSyncPolicy

`func (o *V1alpha1ApplicationSpec) HasSyncPolicy() bool`

HasSyncPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


