# V1alpha1PullRequestGenerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BitbucketServer** | Pointer to [**V1alpha1PullRequestGeneratorBitbucketServer**](V1alpha1PullRequestGeneratorBitbucketServer.md) |  | [optional] 
**Filters** | Pointer to [**[]V1alpha1PullRequestGeneratorFilter**](V1alpha1PullRequestGeneratorFilter.md) | Filters for which pull requests should be considered. | [optional] 
**Gitea** | Pointer to [**V1alpha1PullRequestGeneratorGitea**](V1alpha1PullRequestGeneratorGitea.md) |  | [optional] 
**Github** | Pointer to [**V1alpha1PullRequestGeneratorGithub**](V1alpha1PullRequestGeneratorGithub.md) |  | [optional] 
**Gitlab** | Pointer to [**V1alpha1PullRequestGeneratorGitLab**](V1alpha1PullRequestGeneratorGitLab.md) |  | [optional] 
**RequeueAfterSeconds** | Pointer to **string** | Standard parameters. | [optional] 
**Template** | Pointer to [**V1alpha1ApplicationSetTemplate**](V1alpha1ApplicationSetTemplate.md) |  | [optional] 

## Methods

### NewV1alpha1PullRequestGenerator

`func NewV1alpha1PullRequestGenerator() *V1alpha1PullRequestGenerator`

NewV1alpha1PullRequestGenerator instantiates a new V1alpha1PullRequestGenerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1PullRequestGeneratorWithDefaults

`func NewV1alpha1PullRequestGeneratorWithDefaults() *V1alpha1PullRequestGenerator`

NewV1alpha1PullRequestGeneratorWithDefaults instantiates a new V1alpha1PullRequestGenerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBitbucketServer

`func (o *V1alpha1PullRequestGenerator) GetBitbucketServer() V1alpha1PullRequestGeneratorBitbucketServer`

GetBitbucketServer returns the BitbucketServer field if non-nil, zero value otherwise.

### GetBitbucketServerOk

`func (o *V1alpha1PullRequestGenerator) GetBitbucketServerOk() (*V1alpha1PullRequestGeneratorBitbucketServer, bool)`

GetBitbucketServerOk returns a tuple with the BitbucketServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitbucketServer

`func (o *V1alpha1PullRequestGenerator) SetBitbucketServer(v V1alpha1PullRequestGeneratorBitbucketServer)`

SetBitbucketServer sets BitbucketServer field to given value.

### HasBitbucketServer

`func (o *V1alpha1PullRequestGenerator) HasBitbucketServer() bool`

HasBitbucketServer returns a boolean if a field has been set.

### GetFilters

`func (o *V1alpha1PullRequestGenerator) GetFilters() []V1alpha1PullRequestGeneratorFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *V1alpha1PullRequestGenerator) GetFiltersOk() (*[]V1alpha1PullRequestGeneratorFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *V1alpha1PullRequestGenerator) SetFilters(v []V1alpha1PullRequestGeneratorFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *V1alpha1PullRequestGenerator) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGitea

`func (o *V1alpha1PullRequestGenerator) GetGitea() V1alpha1PullRequestGeneratorGitea`

GetGitea returns the Gitea field if non-nil, zero value otherwise.

### GetGiteaOk

`func (o *V1alpha1PullRequestGenerator) GetGiteaOk() (*V1alpha1PullRequestGeneratorGitea, bool)`

GetGiteaOk returns a tuple with the Gitea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitea

`func (o *V1alpha1PullRequestGenerator) SetGitea(v V1alpha1PullRequestGeneratorGitea)`

SetGitea sets Gitea field to given value.

### HasGitea

`func (o *V1alpha1PullRequestGenerator) HasGitea() bool`

HasGitea returns a boolean if a field has been set.

### GetGithub

`func (o *V1alpha1PullRequestGenerator) GetGithub() V1alpha1PullRequestGeneratorGithub`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *V1alpha1PullRequestGenerator) GetGithubOk() (*V1alpha1PullRequestGeneratorGithub, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *V1alpha1PullRequestGenerator) SetGithub(v V1alpha1PullRequestGeneratorGithub)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *V1alpha1PullRequestGenerator) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetGitlab

`func (o *V1alpha1PullRequestGenerator) GetGitlab() V1alpha1PullRequestGeneratorGitLab`

GetGitlab returns the Gitlab field if non-nil, zero value otherwise.

### GetGitlabOk

`func (o *V1alpha1PullRequestGenerator) GetGitlabOk() (*V1alpha1PullRequestGeneratorGitLab, bool)`

GetGitlabOk returns a tuple with the Gitlab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlab

`func (o *V1alpha1PullRequestGenerator) SetGitlab(v V1alpha1PullRequestGeneratorGitLab)`

SetGitlab sets Gitlab field to given value.

### HasGitlab

`func (o *V1alpha1PullRequestGenerator) HasGitlab() bool`

HasGitlab returns a boolean if a field has been set.

### GetRequeueAfterSeconds

`func (o *V1alpha1PullRequestGenerator) GetRequeueAfterSeconds() string`

GetRequeueAfterSeconds returns the RequeueAfterSeconds field if non-nil, zero value otherwise.

### GetRequeueAfterSecondsOk

`func (o *V1alpha1PullRequestGenerator) GetRequeueAfterSecondsOk() (*string, bool)`

GetRequeueAfterSecondsOk returns a tuple with the RequeueAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequeueAfterSeconds

`func (o *V1alpha1PullRequestGenerator) SetRequeueAfterSeconds(v string)`

SetRequeueAfterSeconds sets RequeueAfterSeconds field to given value.

### HasRequeueAfterSeconds

`func (o *V1alpha1PullRequestGenerator) HasRequeueAfterSeconds() bool`

HasRequeueAfterSeconds returns a boolean if a field has been set.

### GetTemplate

`func (o *V1alpha1PullRequestGenerator) GetTemplate() V1alpha1ApplicationSetTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1alpha1PullRequestGenerator) GetTemplateOk() (*V1alpha1ApplicationSetTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1alpha1PullRequestGenerator) SetTemplate(v V1alpha1ApplicationSetTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *V1alpha1PullRequestGenerator) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


