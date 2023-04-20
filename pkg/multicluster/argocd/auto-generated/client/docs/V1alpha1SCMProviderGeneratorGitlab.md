# V1alpha1SCMProviderGeneratorGitlab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllBranches** | Pointer to **bool** | Scan all branches instead of just the default branch. | [optional] 
**Api** | Pointer to **string** | The Gitlab API URL to talk to. | [optional] 
**Group** | Pointer to **string** | Gitlab group to scan. Required.  You can use either the project id (recommended) or the full namespaced path. | [optional] 
**IncludeSubgroups** | Pointer to **bool** |  | [optional] 
**TokenRef** | Pointer to [**V1alpha1SecretRef**](V1alpha1SecretRef.md) |  | [optional] 

## Methods

### NewV1alpha1SCMProviderGeneratorGitlab

`func NewV1alpha1SCMProviderGeneratorGitlab() *V1alpha1SCMProviderGeneratorGitlab`

NewV1alpha1SCMProviderGeneratorGitlab instantiates a new V1alpha1SCMProviderGeneratorGitlab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1SCMProviderGeneratorGitlabWithDefaults

`func NewV1alpha1SCMProviderGeneratorGitlabWithDefaults() *V1alpha1SCMProviderGeneratorGitlab`

NewV1alpha1SCMProviderGeneratorGitlabWithDefaults instantiates a new V1alpha1SCMProviderGeneratorGitlab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllBranches

`func (o *V1alpha1SCMProviderGeneratorGitlab) GetAllBranches() bool`

GetAllBranches returns the AllBranches field if non-nil, zero value otherwise.

### GetAllBranchesOk

`func (o *V1alpha1SCMProviderGeneratorGitlab) GetAllBranchesOk() (*bool, bool)`

GetAllBranchesOk returns a tuple with the AllBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllBranches

`func (o *V1alpha1SCMProviderGeneratorGitlab) SetAllBranches(v bool)`

SetAllBranches sets AllBranches field to given value.

### HasAllBranches

`func (o *V1alpha1SCMProviderGeneratorGitlab) HasAllBranches() bool`

HasAllBranches returns a boolean if a field has been set.

### GetApi

`func (o *V1alpha1SCMProviderGeneratorGitlab) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *V1alpha1SCMProviderGeneratorGitlab) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *V1alpha1SCMProviderGeneratorGitlab) SetApi(v string)`

SetApi sets Api field to given value.

### HasApi

`func (o *V1alpha1SCMProviderGeneratorGitlab) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetGroup

`func (o *V1alpha1SCMProviderGeneratorGitlab) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *V1alpha1SCMProviderGeneratorGitlab) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *V1alpha1SCMProviderGeneratorGitlab) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *V1alpha1SCMProviderGeneratorGitlab) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetIncludeSubgroups

`func (o *V1alpha1SCMProviderGeneratorGitlab) GetIncludeSubgroups() bool`

GetIncludeSubgroups returns the IncludeSubgroups field if non-nil, zero value otherwise.

### GetIncludeSubgroupsOk

`func (o *V1alpha1SCMProviderGeneratorGitlab) GetIncludeSubgroupsOk() (*bool, bool)`

GetIncludeSubgroupsOk returns a tuple with the IncludeSubgroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubgroups

`func (o *V1alpha1SCMProviderGeneratorGitlab) SetIncludeSubgroups(v bool)`

SetIncludeSubgroups sets IncludeSubgroups field to given value.

### HasIncludeSubgroups

`func (o *V1alpha1SCMProviderGeneratorGitlab) HasIncludeSubgroups() bool`

HasIncludeSubgroups returns a boolean if a field has been set.

### GetTokenRef

`func (o *V1alpha1SCMProviderGeneratorGitlab) GetTokenRef() V1alpha1SecretRef`

GetTokenRef returns the TokenRef field if non-nil, zero value otherwise.

### GetTokenRefOk

`func (o *V1alpha1SCMProviderGeneratorGitlab) GetTokenRefOk() (*V1alpha1SecretRef, bool)`

GetTokenRefOk returns a tuple with the TokenRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRef

`func (o *V1alpha1SCMProviderGeneratorGitlab) SetTokenRef(v V1alpha1SecretRef)`

SetTokenRef sets TokenRef field to given value.

### HasTokenRef

`func (o *V1alpha1SCMProviderGeneratorGitlab) HasTokenRef() bool`

HasTokenRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


