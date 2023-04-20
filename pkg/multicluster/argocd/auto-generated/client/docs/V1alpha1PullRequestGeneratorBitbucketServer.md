# V1alpha1PullRequestGeneratorBitbucketServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Api** | Pointer to **string** | The Bitbucket REST API URL to talk to e.g. https://bitbucket.org/rest Required. | [optional] 
**BasicAuth** | Pointer to [**V1alpha1BasicAuthBitbucketServer**](V1alpha1BasicAuthBitbucketServer.md) |  | [optional] 
**Project** | Pointer to **string** | Project to scan. Required. | [optional] 
**Repo** | Pointer to **string** | Repo name to scan. Required. | [optional] 

## Methods

### NewV1alpha1PullRequestGeneratorBitbucketServer

`func NewV1alpha1PullRequestGeneratorBitbucketServer() *V1alpha1PullRequestGeneratorBitbucketServer`

NewV1alpha1PullRequestGeneratorBitbucketServer instantiates a new V1alpha1PullRequestGeneratorBitbucketServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1PullRequestGeneratorBitbucketServerWithDefaults

`func NewV1alpha1PullRequestGeneratorBitbucketServerWithDefaults() *V1alpha1PullRequestGeneratorBitbucketServer`

NewV1alpha1PullRequestGeneratorBitbucketServerWithDefaults instantiates a new V1alpha1PullRequestGeneratorBitbucketServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApi

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) SetApi(v string)`

SetApi sets Api field to given value.

### HasApi

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetBasicAuth

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetBasicAuth() V1alpha1BasicAuthBitbucketServer`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetBasicAuthOk() (*V1alpha1BasicAuthBitbucketServer, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) SetBasicAuth(v V1alpha1BasicAuthBitbucketServer)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetProject

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRepo

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *V1alpha1PullRequestGeneratorBitbucketServer) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


