# V1alpha1PullRequestGeneratorGitea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Api** | Pointer to **string** |  | [optional] 
**Insecure** | Pointer to **bool** | Allow insecure tls, for self-signed certificates; default: false. | [optional] 
**Owner** | Pointer to **string** | Gitea org or user to scan. Required. | [optional] 
**Repo** | Pointer to **string** | Gitea repo name to scan. Required. | [optional] 
**TokenRef** | Pointer to [**V1alpha1SecretRef**](V1alpha1SecretRef.md) |  | [optional] 

## Methods

### NewV1alpha1PullRequestGeneratorGitea

`func NewV1alpha1PullRequestGeneratorGitea() *V1alpha1PullRequestGeneratorGitea`

NewV1alpha1PullRequestGeneratorGitea instantiates a new V1alpha1PullRequestGeneratorGitea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1PullRequestGeneratorGiteaWithDefaults

`func NewV1alpha1PullRequestGeneratorGiteaWithDefaults() *V1alpha1PullRequestGeneratorGitea`

NewV1alpha1PullRequestGeneratorGiteaWithDefaults instantiates a new V1alpha1PullRequestGeneratorGitea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApi

`func (o *V1alpha1PullRequestGeneratorGitea) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *V1alpha1PullRequestGeneratorGitea) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *V1alpha1PullRequestGeneratorGitea) SetApi(v string)`

SetApi sets Api field to given value.

### HasApi

`func (o *V1alpha1PullRequestGeneratorGitea) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetInsecure

`func (o *V1alpha1PullRequestGeneratorGitea) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *V1alpha1PullRequestGeneratorGitea) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *V1alpha1PullRequestGeneratorGitea) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *V1alpha1PullRequestGeneratorGitea) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetOwner

`func (o *V1alpha1PullRequestGeneratorGitea) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V1alpha1PullRequestGeneratorGitea) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V1alpha1PullRequestGeneratorGitea) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V1alpha1PullRequestGeneratorGitea) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRepo

`func (o *V1alpha1PullRequestGeneratorGitea) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *V1alpha1PullRequestGeneratorGitea) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *V1alpha1PullRequestGeneratorGitea) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *V1alpha1PullRequestGeneratorGitea) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetTokenRef

`func (o *V1alpha1PullRequestGeneratorGitea) GetTokenRef() V1alpha1SecretRef`

GetTokenRef returns the TokenRef field if non-nil, zero value otherwise.

### GetTokenRefOk

`func (o *V1alpha1PullRequestGeneratorGitea) GetTokenRefOk() (*V1alpha1SecretRef, bool)`

GetTokenRefOk returns a tuple with the TokenRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRef

`func (o *V1alpha1PullRequestGeneratorGitea) SetTokenRef(v V1alpha1SecretRef)`

SetTokenRef sets TokenRef field to given value.

### HasTokenRef

`func (o *V1alpha1PullRequestGeneratorGitea) HasTokenRef() bool`

HasTokenRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


