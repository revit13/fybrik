# RepositoryRepoAppDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directory** | Pointer to **map[string]interface{}** |  | [optional] 
**Helm** | Pointer to [**RepositoryHelmAppSpec**](RepositoryHelmAppSpec.md) |  | [optional] 
**Kustomize** | Pointer to [**RepositoryKustomizeAppSpec**](RepositoryKustomizeAppSpec.md) |  | [optional] 
**Plugin** | Pointer to [**RepositoryPluginAppSpec**](RepositoryPluginAppSpec.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewRepositoryRepoAppDetailsResponse

`func NewRepositoryRepoAppDetailsResponse() *RepositoryRepoAppDetailsResponse`

NewRepositoryRepoAppDetailsResponse instantiates a new RepositoryRepoAppDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryRepoAppDetailsResponseWithDefaults

`func NewRepositoryRepoAppDetailsResponseWithDefaults() *RepositoryRepoAppDetailsResponse`

NewRepositoryRepoAppDetailsResponseWithDefaults instantiates a new RepositoryRepoAppDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectory

`func (o *RepositoryRepoAppDetailsResponse) GetDirectory() map[string]interface{}`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *RepositoryRepoAppDetailsResponse) GetDirectoryOk() (*map[string]interface{}, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *RepositoryRepoAppDetailsResponse) SetDirectory(v map[string]interface{})`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *RepositoryRepoAppDetailsResponse) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetHelm

`func (o *RepositoryRepoAppDetailsResponse) GetHelm() RepositoryHelmAppSpec`

GetHelm returns the Helm field if non-nil, zero value otherwise.

### GetHelmOk

`func (o *RepositoryRepoAppDetailsResponse) GetHelmOk() (*RepositoryHelmAppSpec, bool)`

GetHelmOk returns a tuple with the Helm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelm

`func (o *RepositoryRepoAppDetailsResponse) SetHelm(v RepositoryHelmAppSpec)`

SetHelm sets Helm field to given value.

### HasHelm

`func (o *RepositoryRepoAppDetailsResponse) HasHelm() bool`

HasHelm returns a boolean if a field has been set.

### GetKustomize

`func (o *RepositoryRepoAppDetailsResponse) GetKustomize() RepositoryKustomizeAppSpec`

GetKustomize returns the Kustomize field if non-nil, zero value otherwise.

### GetKustomizeOk

`func (o *RepositoryRepoAppDetailsResponse) GetKustomizeOk() (*RepositoryKustomizeAppSpec, bool)`

GetKustomizeOk returns a tuple with the Kustomize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKustomize

`func (o *RepositoryRepoAppDetailsResponse) SetKustomize(v RepositoryKustomizeAppSpec)`

SetKustomize sets Kustomize field to given value.

### HasKustomize

`func (o *RepositoryRepoAppDetailsResponse) HasKustomize() bool`

HasKustomize returns a boolean if a field has been set.

### GetPlugin

`func (o *RepositoryRepoAppDetailsResponse) GetPlugin() RepositoryPluginAppSpec`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *RepositoryRepoAppDetailsResponse) GetPluginOk() (*RepositoryPluginAppSpec, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *RepositoryRepoAppDetailsResponse) SetPlugin(v RepositoryPluginAppSpec)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *RepositoryRepoAppDetailsResponse) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetType

`func (o *RepositoryRepoAppDetailsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RepositoryRepoAppDetailsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RepositoryRepoAppDetailsResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RepositoryRepoAppDetailsResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


