# V1alpha1SCMProviderGenerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureDevOps** | Pointer to [**V1alpha1SCMProviderGeneratorAzureDevOps**](V1alpha1SCMProviderGeneratorAzureDevOps.md) |  | [optional] 
**Bitbucket** | Pointer to [**V1alpha1SCMProviderGeneratorBitbucket**](V1alpha1SCMProviderGeneratorBitbucket.md) |  | [optional] 
**BitbucketServer** | Pointer to [**V1alpha1SCMProviderGeneratorBitbucketServer**](V1alpha1SCMProviderGeneratorBitbucketServer.md) |  | [optional] 
**CloneProtocol** | Pointer to **string** | Which protocol to use for the SCM URL. Default is provider-specific but ssh if possible. Not all providers necessarily support all protocols. | [optional] 
**Filters** | Pointer to [**[]V1alpha1SCMProviderGeneratorFilter**](V1alpha1SCMProviderGeneratorFilter.md) | Filters for which repos should be considered. | [optional] 
**Gitea** | Pointer to [**V1alpha1SCMProviderGeneratorGitea**](V1alpha1SCMProviderGeneratorGitea.md) |  | [optional] 
**Github** | Pointer to [**V1alpha1SCMProviderGeneratorGithub**](V1alpha1SCMProviderGeneratorGithub.md) |  | [optional] 
**Gitlab** | Pointer to [**V1alpha1SCMProviderGeneratorGitlab**](V1alpha1SCMProviderGeneratorGitlab.md) |  | [optional] 
**RequeueAfterSeconds** | Pointer to **string** | Standard parameters. | [optional] 
**Template** | Pointer to [**V1alpha1ApplicationSetTemplate**](V1alpha1ApplicationSetTemplate.md) |  | [optional] 

## Methods

### NewV1alpha1SCMProviderGenerator

`func NewV1alpha1SCMProviderGenerator() *V1alpha1SCMProviderGenerator`

NewV1alpha1SCMProviderGenerator instantiates a new V1alpha1SCMProviderGenerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1SCMProviderGeneratorWithDefaults

`func NewV1alpha1SCMProviderGeneratorWithDefaults() *V1alpha1SCMProviderGenerator`

NewV1alpha1SCMProviderGeneratorWithDefaults instantiates a new V1alpha1SCMProviderGenerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureDevOps

`func (o *V1alpha1SCMProviderGenerator) GetAzureDevOps() V1alpha1SCMProviderGeneratorAzureDevOps`

GetAzureDevOps returns the AzureDevOps field if non-nil, zero value otherwise.

### GetAzureDevOpsOk

`func (o *V1alpha1SCMProviderGenerator) GetAzureDevOpsOk() (*V1alpha1SCMProviderGeneratorAzureDevOps, bool)`

GetAzureDevOpsOk returns a tuple with the AzureDevOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureDevOps

`func (o *V1alpha1SCMProviderGenerator) SetAzureDevOps(v V1alpha1SCMProviderGeneratorAzureDevOps)`

SetAzureDevOps sets AzureDevOps field to given value.

### HasAzureDevOps

`func (o *V1alpha1SCMProviderGenerator) HasAzureDevOps() bool`

HasAzureDevOps returns a boolean if a field has been set.

### GetBitbucket

`func (o *V1alpha1SCMProviderGenerator) GetBitbucket() V1alpha1SCMProviderGeneratorBitbucket`

GetBitbucket returns the Bitbucket field if non-nil, zero value otherwise.

### GetBitbucketOk

`func (o *V1alpha1SCMProviderGenerator) GetBitbucketOk() (*V1alpha1SCMProviderGeneratorBitbucket, bool)`

GetBitbucketOk returns a tuple with the Bitbucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitbucket

`func (o *V1alpha1SCMProviderGenerator) SetBitbucket(v V1alpha1SCMProviderGeneratorBitbucket)`

SetBitbucket sets Bitbucket field to given value.

### HasBitbucket

`func (o *V1alpha1SCMProviderGenerator) HasBitbucket() bool`

HasBitbucket returns a boolean if a field has been set.

### GetBitbucketServer

`func (o *V1alpha1SCMProviderGenerator) GetBitbucketServer() V1alpha1SCMProviderGeneratorBitbucketServer`

GetBitbucketServer returns the BitbucketServer field if non-nil, zero value otherwise.

### GetBitbucketServerOk

`func (o *V1alpha1SCMProviderGenerator) GetBitbucketServerOk() (*V1alpha1SCMProviderGeneratorBitbucketServer, bool)`

GetBitbucketServerOk returns a tuple with the BitbucketServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitbucketServer

`func (o *V1alpha1SCMProviderGenerator) SetBitbucketServer(v V1alpha1SCMProviderGeneratorBitbucketServer)`

SetBitbucketServer sets BitbucketServer field to given value.

### HasBitbucketServer

`func (o *V1alpha1SCMProviderGenerator) HasBitbucketServer() bool`

HasBitbucketServer returns a boolean if a field has been set.

### GetCloneProtocol

`func (o *V1alpha1SCMProviderGenerator) GetCloneProtocol() string`

GetCloneProtocol returns the CloneProtocol field if non-nil, zero value otherwise.

### GetCloneProtocolOk

`func (o *V1alpha1SCMProviderGenerator) GetCloneProtocolOk() (*string, bool)`

GetCloneProtocolOk returns a tuple with the CloneProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneProtocol

`func (o *V1alpha1SCMProviderGenerator) SetCloneProtocol(v string)`

SetCloneProtocol sets CloneProtocol field to given value.

### HasCloneProtocol

`func (o *V1alpha1SCMProviderGenerator) HasCloneProtocol() bool`

HasCloneProtocol returns a boolean if a field has been set.

### GetFilters

`func (o *V1alpha1SCMProviderGenerator) GetFilters() []V1alpha1SCMProviderGeneratorFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *V1alpha1SCMProviderGenerator) GetFiltersOk() (*[]V1alpha1SCMProviderGeneratorFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *V1alpha1SCMProviderGenerator) SetFilters(v []V1alpha1SCMProviderGeneratorFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *V1alpha1SCMProviderGenerator) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGitea

`func (o *V1alpha1SCMProviderGenerator) GetGitea() V1alpha1SCMProviderGeneratorGitea`

GetGitea returns the Gitea field if non-nil, zero value otherwise.

### GetGiteaOk

`func (o *V1alpha1SCMProviderGenerator) GetGiteaOk() (*V1alpha1SCMProviderGeneratorGitea, bool)`

GetGiteaOk returns a tuple with the Gitea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitea

`func (o *V1alpha1SCMProviderGenerator) SetGitea(v V1alpha1SCMProviderGeneratorGitea)`

SetGitea sets Gitea field to given value.

### HasGitea

`func (o *V1alpha1SCMProviderGenerator) HasGitea() bool`

HasGitea returns a boolean if a field has been set.

### GetGithub

`func (o *V1alpha1SCMProviderGenerator) GetGithub() V1alpha1SCMProviderGeneratorGithub`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *V1alpha1SCMProviderGenerator) GetGithubOk() (*V1alpha1SCMProviderGeneratorGithub, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *V1alpha1SCMProviderGenerator) SetGithub(v V1alpha1SCMProviderGeneratorGithub)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *V1alpha1SCMProviderGenerator) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetGitlab

`func (o *V1alpha1SCMProviderGenerator) GetGitlab() V1alpha1SCMProviderGeneratorGitlab`

GetGitlab returns the Gitlab field if non-nil, zero value otherwise.

### GetGitlabOk

`func (o *V1alpha1SCMProviderGenerator) GetGitlabOk() (*V1alpha1SCMProviderGeneratorGitlab, bool)`

GetGitlabOk returns a tuple with the Gitlab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlab

`func (o *V1alpha1SCMProviderGenerator) SetGitlab(v V1alpha1SCMProviderGeneratorGitlab)`

SetGitlab sets Gitlab field to given value.

### HasGitlab

`func (o *V1alpha1SCMProviderGenerator) HasGitlab() bool`

HasGitlab returns a boolean if a field has been set.

### GetRequeueAfterSeconds

`func (o *V1alpha1SCMProviderGenerator) GetRequeueAfterSeconds() string`

GetRequeueAfterSeconds returns the RequeueAfterSeconds field if non-nil, zero value otherwise.

### GetRequeueAfterSecondsOk

`func (o *V1alpha1SCMProviderGenerator) GetRequeueAfterSecondsOk() (*string, bool)`

GetRequeueAfterSecondsOk returns a tuple with the RequeueAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequeueAfterSeconds

`func (o *V1alpha1SCMProviderGenerator) SetRequeueAfterSeconds(v string)`

SetRequeueAfterSeconds sets RequeueAfterSeconds field to given value.

### HasRequeueAfterSeconds

`func (o *V1alpha1SCMProviderGenerator) HasRequeueAfterSeconds() bool`

HasRequeueAfterSeconds returns a boolean if a field has been set.

### GetTemplate

`func (o *V1alpha1SCMProviderGenerator) GetTemplate() V1alpha1ApplicationSetTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1alpha1SCMProviderGenerator) GetTemplateOk() (*V1alpha1ApplicationSetTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1alpha1SCMProviderGenerator) SetTemplate(v V1alpha1ApplicationSetTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *V1alpha1SCMProviderGenerator) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


