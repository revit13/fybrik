# V1alpha1PullRequestGeneratorGithub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Api** | Pointer to **string** | The GitHub API URL to talk to. If blank, use https://api.github.com/. | [optional] 
**AppSecretName** | Pointer to **string** | AppSecretName is a reference to a GitHub App repo-creds secret with permission to access pull requests. | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Owner** | Pointer to **string** | GitHub org or user to scan. Required. | [optional] 
**Repo** | Pointer to **string** | GitHub repo name to scan. Required. | [optional] 
**TokenRef** | Pointer to [**V1alpha1SecretRef**](V1alpha1SecretRef.md) |  | [optional] 

## Methods

### NewV1alpha1PullRequestGeneratorGithub

`func NewV1alpha1PullRequestGeneratorGithub() *V1alpha1PullRequestGeneratorGithub`

NewV1alpha1PullRequestGeneratorGithub instantiates a new V1alpha1PullRequestGeneratorGithub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1PullRequestGeneratorGithubWithDefaults

`func NewV1alpha1PullRequestGeneratorGithubWithDefaults() *V1alpha1PullRequestGeneratorGithub`

NewV1alpha1PullRequestGeneratorGithubWithDefaults instantiates a new V1alpha1PullRequestGeneratorGithub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApi

`func (o *V1alpha1PullRequestGeneratorGithub) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *V1alpha1PullRequestGeneratorGithub) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *V1alpha1PullRequestGeneratorGithub) SetApi(v string)`

SetApi sets Api field to given value.

### HasApi

`func (o *V1alpha1PullRequestGeneratorGithub) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetAppSecretName

`func (o *V1alpha1PullRequestGeneratorGithub) GetAppSecretName() string`

GetAppSecretName returns the AppSecretName field if non-nil, zero value otherwise.

### GetAppSecretNameOk

`func (o *V1alpha1PullRequestGeneratorGithub) GetAppSecretNameOk() (*string, bool)`

GetAppSecretNameOk returns a tuple with the AppSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSecretName

`func (o *V1alpha1PullRequestGeneratorGithub) SetAppSecretName(v string)`

SetAppSecretName sets AppSecretName field to given value.

### HasAppSecretName

`func (o *V1alpha1PullRequestGeneratorGithub) HasAppSecretName() bool`

HasAppSecretName returns a boolean if a field has been set.

### GetLabels

`func (o *V1alpha1PullRequestGeneratorGithub) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1alpha1PullRequestGeneratorGithub) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1alpha1PullRequestGeneratorGithub) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1alpha1PullRequestGeneratorGithub) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOwner

`func (o *V1alpha1PullRequestGeneratorGithub) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V1alpha1PullRequestGeneratorGithub) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V1alpha1PullRequestGeneratorGithub) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V1alpha1PullRequestGeneratorGithub) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRepo

`func (o *V1alpha1PullRequestGeneratorGithub) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *V1alpha1PullRequestGeneratorGithub) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *V1alpha1PullRequestGeneratorGithub) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *V1alpha1PullRequestGeneratorGithub) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetTokenRef

`func (o *V1alpha1PullRequestGeneratorGithub) GetTokenRef() V1alpha1SecretRef`

GetTokenRef returns the TokenRef field if non-nil, zero value otherwise.

### GetTokenRefOk

`func (o *V1alpha1PullRequestGeneratorGithub) GetTokenRefOk() (*V1alpha1SecretRef, bool)`

GetTokenRefOk returns a tuple with the TokenRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRef

`func (o *V1alpha1PullRequestGeneratorGithub) SetTokenRef(v V1alpha1SecretRef)`

SetTokenRef sets TokenRef field to given value.

### HasTokenRef

`func (o *V1alpha1PullRequestGeneratorGithub) HasTokenRef() bool`

HasTokenRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


