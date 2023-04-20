# V1alpha1SCMProviderGeneratorGithub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllBranches** | Pointer to **bool** | Scan all branches instead of just the default branch. | [optional] 
**Api** | Pointer to **string** | The GitHub API URL to talk to. If blank, use https://api.github.com/. | [optional] 
**AppSecretName** | Pointer to **string** | AppSecretName is a reference to a GitHub App repo-creds secret. | [optional] 
**Organization** | Pointer to **string** | GitHub org to scan. Required. | [optional] 
**TokenRef** | Pointer to [**V1alpha1SecretRef**](V1alpha1SecretRef.md) |  | [optional] 

## Methods

### NewV1alpha1SCMProviderGeneratorGithub

`func NewV1alpha1SCMProviderGeneratorGithub() *V1alpha1SCMProviderGeneratorGithub`

NewV1alpha1SCMProviderGeneratorGithub instantiates a new V1alpha1SCMProviderGeneratorGithub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1SCMProviderGeneratorGithubWithDefaults

`func NewV1alpha1SCMProviderGeneratorGithubWithDefaults() *V1alpha1SCMProviderGeneratorGithub`

NewV1alpha1SCMProviderGeneratorGithubWithDefaults instantiates a new V1alpha1SCMProviderGeneratorGithub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllBranches

`func (o *V1alpha1SCMProviderGeneratorGithub) GetAllBranches() bool`

GetAllBranches returns the AllBranches field if non-nil, zero value otherwise.

### GetAllBranchesOk

`func (o *V1alpha1SCMProviderGeneratorGithub) GetAllBranchesOk() (*bool, bool)`

GetAllBranchesOk returns a tuple with the AllBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllBranches

`func (o *V1alpha1SCMProviderGeneratorGithub) SetAllBranches(v bool)`

SetAllBranches sets AllBranches field to given value.

### HasAllBranches

`func (o *V1alpha1SCMProviderGeneratorGithub) HasAllBranches() bool`

HasAllBranches returns a boolean if a field has been set.

### GetApi

`func (o *V1alpha1SCMProviderGeneratorGithub) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *V1alpha1SCMProviderGeneratorGithub) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *V1alpha1SCMProviderGeneratorGithub) SetApi(v string)`

SetApi sets Api field to given value.

### HasApi

`func (o *V1alpha1SCMProviderGeneratorGithub) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetAppSecretName

`func (o *V1alpha1SCMProviderGeneratorGithub) GetAppSecretName() string`

GetAppSecretName returns the AppSecretName field if non-nil, zero value otherwise.

### GetAppSecretNameOk

`func (o *V1alpha1SCMProviderGeneratorGithub) GetAppSecretNameOk() (*string, bool)`

GetAppSecretNameOk returns a tuple with the AppSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSecretName

`func (o *V1alpha1SCMProviderGeneratorGithub) SetAppSecretName(v string)`

SetAppSecretName sets AppSecretName field to given value.

### HasAppSecretName

`func (o *V1alpha1SCMProviderGeneratorGithub) HasAppSecretName() bool`

HasAppSecretName returns a boolean if a field has been set.

### GetOrganization

`func (o *V1alpha1SCMProviderGeneratorGithub) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *V1alpha1SCMProviderGeneratorGithub) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *V1alpha1SCMProviderGeneratorGithub) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *V1alpha1SCMProviderGeneratorGithub) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetTokenRef

`func (o *V1alpha1SCMProviderGeneratorGithub) GetTokenRef() V1alpha1SecretRef`

GetTokenRef returns the TokenRef field if non-nil, zero value otherwise.

### GetTokenRefOk

`func (o *V1alpha1SCMProviderGeneratorGithub) GetTokenRefOk() (*V1alpha1SecretRef, bool)`

GetTokenRefOk returns a tuple with the TokenRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRef

`func (o *V1alpha1SCMProviderGeneratorGithub) SetTokenRef(v V1alpha1SecretRef)`

SetTokenRef sets TokenRef field to given value.

### HasTokenRef

`func (o *V1alpha1SCMProviderGeneratorGithub) HasTokenRef() bool`

HasTokenRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


