# V1alpha1ApplicationSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chart** | Pointer to **string** | Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo. | [optional] 
**Directory** | Pointer to [**V1alpha1ApplicationSourceDirectory**](V1alpha1ApplicationSourceDirectory.md) |  | [optional] 
**Helm** | Pointer to [**V1alpha1ApplicationSourceHelm**](V1alpha1ApplicationSourceHelm.md) |  | [optional] 
**Kustomize** | Pointer to [**V1alpha1ApplicationSourceKustomize**](V1alpha1ApplicationSourceKustomize.md) |  | [optional] 
**Path** | Pointer to **string** | Path is a directory path within the Git repository, and is only valid for applications sourced from Git. | [optional] 
**Plugin** | Pointer to [**V1alpha1ApplicationSourcePlugin**](V1alpha1ApplicationSourcePlugin.md) |  | [optional] 
**Ref** | Pointer to **string** | Ref is reference to another source within sources field. This field will not be used if used with a &#x60;source&#x60; tag. | [optional] 
**RepoURL** | Pointer to **string** |  | [optional] 
**TargetRevision** | Pointer to **string** | TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart&#39;s version. | [optional] 

## Methods

### NewV1alpha1ApplicationSource

`func NewV1alpha1ApplicationSource() *V1alpha1ApplicationSource`

NewV1alpha1ApplicationSource instantiates a new V1alpha1ApplicationSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSourceWithDefaults

`func NewV1alpha1ApplicationSourceWithDefaults() *V1alpha1ApplicationSource`

NewV1alpha1ApplicationSourceWithDefaults instantiates a new V1alpha1ApplicationSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChart

`func (o *V1alpha1ApplicationSource) GetChart() string`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *V1alpha1ApplicationSource) GetChartOk() (*string, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *V1alpha1ApplicationSource) SetChart(v string)`

SetChart sets Chart field to given value.

### HasChart

`func (o *V1alpha1ApplicationSource) HasChart() bool`

HasChart returns a boolean if a field has been set.

### GetDirectory

`func (o *V1alpha1ApplicationSource) GetDirectory() V1alpha1ApplicationSourceDirectory`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *V1alpha1ApplicationSource) GetDirectoryOk() (*V1alpha1ApplicationSourceDirectory, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *V1alpha1ApplicationSource) SetDirectory(v V1alpha1ApplicationSourceDirectory)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *V1alpha1ApplicationSource) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetHelm

`func (o *V1alpha1ApplicationSource) GetHelm() V1alpha1ApplicationSourceHelm`

GetHelm returns the Helm field if non-nil, zero value otherwise.

### GetHelmOk

`func (o *V1alpha1ApplicationSource) GetHelmOk() (*V1alpha1ApplicationSourceHelm, bool)`

GetHelmOk returns a tuple with the Helm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelm

`func (o *V1alpha1ApplicationSource) SetHelm(v V1alpha1ApplicationSourceHelm)`

SetHelm sets Helm field to given value.

### HasHelm

`func (o *V1alpha1ApplicationSource) HasHelm() bool`

HasHelm returns a boolean if a field has been set.

### GetKustomize

`func (o *V1alpha1ApplicationSource) GetKustomize() V1alpha1ApplicationSourceKustomize`

GetKustomize returns the Kustomize field if non-nil, zero value otherwise.

### GetKustomizeOk

`func (o *V1alpha1ApplicationSource) GetKustomizeOk() (*V1alpha1ApplicationSourceKustomize, bool)`

GetKustomizeOk returns a tuple with the Kustomize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKustomize

`func (o *V1alpha1ApplicationSource) SetKustomize(v V1alpha1ApplicationSourceKustomize)`

SetKustomize sets Kustomize field to given value.

### HasKustomize

`func (o *V1alpha1ApplicationSource) HasKustomize() bool`

HasKustomize returns a boolean if a field has been set.

### GetPath

`func (o *V1alpha1ApplicationSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *V1alpha1ApplicationSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *V1alpha1ApplicationSource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *V1alpha1ApplicationSource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPlugin

`func (o *V1alpha1ApplicationSource) GetPlugin() V1alpha1ApplicationSourcePlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *V1alpha1ApplicationSource) GetPluginOk() (*V1alpha1ApplicationSourcePlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *V1alpha1ApplicationSource) SetPlugin(v V1alpha1ApplicationSourcePlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *V1alpha1ApplicationSource) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetRef

`func (o *V1alpha1ApplicationSource) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *V1alpha1ApplicationSource) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *V1alpha1ApplicationSource) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *V1alpha1ApplicationSource) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRepoURL

`func (o *V1alpha1ApplicationSource) GetRepoURL() string`

GetRepoURL returns the RepoURL field if non-nil, zero value otherwise.

### GetRepoURLOk

`func (o *V1alpha1ApplicationSource) GetRepoURLOk() (*string, bool)`

GetRepoURLOk returns a tuple with the RepoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoURL

`func (o *V1alpha1ApplicationSource) SetRepoURL(v string)`

SetRepoURL sets RepoURL field to given value.

### HasRepoURL

`func (o *V1alpha1ApplicationSource) HasRepoURL() bool`

HasRepoURL returns a boolean if a field has been set.

### GetTargetRevision

`func (o *V1alpha1ApplicationSource) GetTargetRevision() string`

GetTargetRevision returns the TargetRevision field if non-nil, zero value otherwise.

### GetTargetRevisionOk

`func (o *V1alpha1ApplicationSource) GetTargetRevisionOk() (*string, bool)`

GetTargetRevisionOk returns a tuple with the TargetRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRevision

`func (o *V1alpha1ApplicationSource) SetTargetRevision(v string)`

SetTargetRevision sets TargetRevision field to given value.

### HasTargetRevision

`func (o *V1alpha1ApplicationSource) HasTargetRevision() bool`

HasTargetRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


