# V1alpha1AppProjectSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterResourceBlacklist** | Pointer to [**[]V1GroupKind**](V1GroupKind.md) |  | [optional] 
**ClusterResourceWhitelist** | Pointer to [**[]V1GroupKind**](V1GroupKind.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Destinations** | Pointer to [**[]V1alpha1ApplicationDestination**](V1alpha1ApplicationDestination.md) |  | [optional] 
**NamespaceResourceBlacklist** | Pointer to [**[]V1GroupKind**](V1GroupKind.md) |  | [optional] 
**NamespaceResourceWhitelist** | Pointer to [**[]V1GroupKind**](V1GroupKind.md) |  | [optional] 
**OrphanedResources** | Pointer to [**V1alpha1OrphanedResourcesMonitorSettings**](V1alpha1OrphanedResourcesMonitorSettings.md) |  | [optional] 
**PermitOnlyProjectScopedClusters** | Pointer to **bool** |  | [optional] 
**Roles** | Pointer to [**[]V1alpha1ProjectRole**](V1alpha1ProjectRole.md) |  | [optional] 
**SignatureKeys** | Pointer to [**[]V1alpha1SignatureKey**](V1alpha1SignatureKey.md) |  | [optional] 
**SourceNamespaces** | Pointer to **[]string** |  | [optional] 
**SourceRepos** | Pointer to **[]string** |  | [optional] 
**SyncWindows** | Pointer to [**[]V1alpha1SyncWindow**](V1alpha1SyncWindow.md) |  | [optional] 

## Methods

### NewV1alpha1AppProjectSpec

`func NewV1alpha1AppProjectSpec() *V1alpha1AppProjectSpec`

NewV1alpha1AppProjectSpec instantiates a new V1alpha1AppProjectSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1AppProjectSpecWithDefaults

`func NewV1alpha1AppProjectSpecWithDefaults() *V1alpha1AppProjectSpec`

NewV1alpha1AppProjectSpecWithDefaults instantiates a new V1alpha1AppProjectSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterResourceBlacklist

`func (o *V1alpha1AppProjectSpec) GetClusterResourceBlacklist() []V1GroupKind`

GetClusterResourceBlacklist returns the ClusterResourceBlacklist field if non-nil, zero value otherwise.

### GetClusterResourceBlacklistOk

`func (o *V1alpha1AppProjectSpec) GetClusterResourceBlacklistOk() (*[]V1GroupKind, bool)`

GetClusterResourceBlacklistOk returns a tuple with the ClusterResourceBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterResourceBlacklist

`func (o *V1alpha1AppProjectSpec) SetClusterResourceBlacklist(v []V1GroupKind)`

SetClusterResourceBlacklist sets ClusterResourceBlacklist field to given value.

### HasClusterResourceBlacklist

`func (o *V1alpha1AppProjectSpec) HasClusterResourceBlacklist() bool`

HasClusterResourceBlacklist returns a boolean if a field has been set.

### GetClusterResourceWhitelist

`func (o *V1alpha1AppProjectSpec) GetClusterResourceWhitelist() []V1GroupKind`

GetClusterResourceWhitelist returns the ClusterResourceWhitelist field if non-nil, zero value otherwise.

### GetClusterResourceWhitelistOk

`func (o *V1alpha1AppProjectSpec) GetClusterResourceWhitelistOk() (*[]V1GroupKind, bool)`

GetClusterResourceWhitelistOk returns a tuple with the ClusterResourceWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterResourceWhitelist

`func (o *V1alpha1AppProjectSpec) SetClusterResourceWhitelist(v []V1GroupKind)`

SetClusterResourceWhitelist sets ClusterResourceWhitelist field to given value.

### HasClusterResourceWhitelist

`func (o *V1alpha1AppProjectSpec) HasClusterResourceWhitelist() bool`

HasClusterResourceWhitelist returns a boolean if a field has been set.

### GetDescription

`func (o *V1alpha1AppProjectSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1alpha1AppProjectSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1alpha1AppProjectSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1alpha1AppProjectSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestinations

`func (o *V1alpha1AppProjectSpec) GetDestinations() []V1alpha1ApplicationDestination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *V1alpha1AppProjectSpec) GetDestinationsOk() (*[]V1alpha1ApplicationDestination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *V1alpha1AppProjectSpec) SetDestinations(v []V1alpha1ApplicationDestination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *V1alpha1AppProjectSpec) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetNamespaceResourceBlacklist

`func (o *V1alpha1AppProjectSpec) GetNamespaceResourceBlacklist() []V1GroupKind`

GetNamespaceResourceBlacklist returns the NamespaceResourceBlacklist field if non-nil, zero value otherwise.

### GetNamespaceResourceBlacklistOk

`func (o *V1alpha1AppProjectSpec) GetNamespaceResourceBlacklistOk() (*[]V1GroupKind, bool)`

GetNamespaceResourceBlacklistOk returns a tuple with the NamespaceResourceBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceResourceBlacklist

`func (o *V1alpha1AppProjectSpec) SetNamespaceResourceBlacklist(v []V1GroupKind)`

SetNamespaceResourceBlacklist sets NamespaceResourceBlacklist field to given value.

### HasNamespaceResourceBlacklist

`func (o *V1alpha1AppProjectSpec) HasNamespaceResourceBlacklist() bool`

HasNamespaceResourceBlacklist returns a boolean if a field has been set.

### GetNamespaceResourceWhitelist

`func (o *V1alpha1AppProjectSpec) GetNamespaceResourceWhitelist() []V1GroupKind`

GetNamespaceResourceWhitelist returns the NamespaceResourceWhitelist field if non-nil, zero value otherwise.

### GetNamespaceResourceWhitelistOk

`func (o *V1alpha1AppProjectSpec) GetNamespaceResourceWhitelistOk() (*[]V1GroupKind, bool)`

GetNamespaceResourceWhitelistOk returns a tuple with the NamespaceResourceWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceResourceWhitelist

`func (o *V1alpha1AppProjectSpec) SetNamespaceResourceWhitelist(v []V1GroupKind)`

SetNamespaceResourceWhitelist sets NamespaceResourceWhitelist field to given value.

### HasNamespaceResourceWhitelist

`func (o *V1alpha1AppProjectSpec) HasNamespaceResourceWhitelist() bool`

HasNamespaceResourceWhitelist returns a boolean if a field has been set.

### GetOrphanedResources

`func (o *V1alpha1AppProjectSpec) GetOrphanedResources() V1alpha1OrphanedResourcesMonitorSettings`

GetOrphanedResources returns the OrphanedResources field if non-nil, zero value otherwise.

### GetOrphanedResourcesOk

`func (o *V1alpha1AppProjectSpec) GetOrphanedResourcesOk() (*V1alpha1OrphanedResourcesMonitorSettings, bool)`

GetOrphanedResourcesOk returns a tuple with the OrphanedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphanedResources

`func (o *V1alpha1AppProjectSpec) SetOrphanedResources(v V1alpha1OrphanedResourcesMonitorSettings)`

SetOrphanedResources sets OrphanedResources field to given value.

### HasOrphanedResources

`func (o *V1alpha1AppProjectSpec) HasOrphanedResources() bool`

HasOrphanedResources returns a boolean if a field has been set.

### GetPermitOnlyProjectScopedClusters

`func (o *V1alpha1AppProjectSpec) GetPermitOnlyProjectScopedClusters() bool`

GetPermitOnlyProjectScopedClusters returns the PermitOnlyProjectScopedClusters field if non-nil, zero value otherwise.

### GetPermitOnlyProjectScopedClustersOk

`func (o *V1alpha1AppProjectSpec) GetPermitOnlyProjectScopedClustersOk() (*bool, bool)`

GetPermitOnlyProjectScopedClustersOk returns a tuple with the PermitOnlyProjectScopedClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitOnlyProjectScopedClusters

`func (o *V1alpha1AppProjectSpec) SetPermitOnlyProjectScopedClusters(v bool)`

SetPermitOnlyProjectScopedClusters sets PermitOnlyProjectScopedClusters field to given value.

### HasPermitOnlyProjectScopedClusters

`func (o *V1alpha1AppProjectSpec) HasPermitOnlyProjectScopedClusters() bool`

HasPermitOnlyProjectScopedClusters returns a boolean if a field has been set.

### GetRoles

`func (o *V1alpha1AppProjectSpec) GetRoles() []V1alpha1ProjectRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *V1alpha1AppProjectSpec) GetRolesOk() (*[]V1alpha1ProjectRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *V1alpha1AppProjectSpec) SetRoles(v []V1alpha1ProjectRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *V1alpha1AppProjectSpec) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSignatureKeys

`func (o *V1alpha1AppProjectSpec) GetSignatureKeys() []V1alpha1SignatureKey`

GetSignatureKeys returns the SignatureKeys field if non-nil, zero value otherwise.

### GetSignatureKeysOk

`func (o *V1alpha1AppProjectSpec) GetSignatureKeysOk() (*[]V1alpha1SignatureKey, bool)`

GetSignatureKeysOk returns a tuple with the SignatureKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureKeys

`func (o *V1alpha1AppProjectSpec) SetSignatureKeys(v []V1alpha1SignatureKey)`

SetSignatureKeys sets SignatureKeys field to given value.

### HasSignatureKeys

`func (o *V1alpha1AppProjectSpec) HasSignatureKeys() bool`

HasSignatureKeys returns a boolean if a field has been set.

### GetSourceNamespaces

`func (o *V1alpha1AppProjectSpec) GetSourceNamespaces() []string`

GetSourceNamespaces returns the SourceNamespaces field if non-nil, zero value otherwise.

### GetSourceNamespacesOk

`func (o *V1alpha1AppProjectSpec) GetSourceNamespacesOk() (*[]string, bool)`

GetSourceNamespacesOk returns a tuple with the SourceNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNamespaces

`func (o *V1alpha1AppProjectSpec) SetSourceNamespaces(v []string)`

SetSourceNamespaces sets SourceNamespaces field to given value.

### HasSourceNamespaces

`func (o *V1alpha1AppProjectSpec) HasSourceNamespaces() bool`

HasSourceNamespaces returns a boolean if a field has been set.

### GetSourceRepos

`func (o *V1alpha1AppProjectSpec) GetSourceRepos() []string`

GetSourceRepos returns the SourceRepos field if non-nil, zero value otherwise.

### GetSourceReposOk

`func (o *V1alpha1AppProjectSpec) GetSourceReposOk() (*[]string, bool)`

GetSourceReposOk returns a tuple with the SourceRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepos

`func (o *V1alpha1AppProjectSpec) SetSourceRepos(v []string)`

SetSourceRepos sets SourceRepos field to given value.

### HasSourceRepos

`func (o *V1alpha1AppProjectSpec) HasSourceRepos() bool`

HasSourceRepos returns a boolean if a field has been set.

### GetSyncWindows

`func (o *V1alpha1AppProjectSpec) GetSyncWindows() []V1alpha1SyncWindow`

GetSyncWindows returns the SyncWindows field if non-nil, zero value otherwise.

### GetSyncWindowsOk

`func (o *V1alpha1AppProjectSpec) GetSyncWindowsOk() (*[]V1alpha1SyncWindow, bool)`

GetSyncWindowsOk returns a tuple with the SyncWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncWindows

`func (o *V1alpha1AppProjectSpec) SetSyncWindows(v []V1alpha1SyncWindow)`

SetSyncWindows sets SyncWindows field to given value.

### HasSyncWindows

`func (o *V1alpha1AppProjectSpec) HasSyncWindows() bool`

HasSyncWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


