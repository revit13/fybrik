# V1alpha1SCMProviderGeneratorGitea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllBranches** | Pointer to **bool** | Scan all branches instead of just the default branch. | [optional] 
**Api** | Pointer to **string** | The Gitea URL to talk to. For example https://gitea.mydomain.com/. | [optional] 
**Insecure** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to **string** | Gitea organization or user to scan. Required. | [optional] 
**TokenRef** | Pointer to [**V1alpha1SecretRef**](V1alpha1SecretRef.md) |  | [optional] 

## Methods

### NewV1alpha1SCMProviderGeneratorGitea

`func NewV1alpha1SCMProviderGeneratorGitea() *V1alpha1SCMProviderGeneratorGitea`

NewV1alpha1SCMProviderGeneratorGitea instantiates a new V1alpha1SCMProviderGeneratorGitea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1SCMProviderGeneratorGiteaWithDefaults

`func NewV1alpha1SCMProviderGeneratorGiteaWithDefaults() *V1alpha1SCMProviderGeneratorGitea`

NewV1alpha1SCMProviderGeneratorGiteaWithDefaults instantiates a new V1alpha1SCMProviderGeneratorGitea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllBranches

`func (o *V1alpha1SCMProviderGeneratorGitea) GetAllBranches() bool`

GetAllBranches returns the AllBranches field if non-nil, zero value otherwise.

### GetAllBranchesOk

`func (o *V1alpha1SCMProviderGeneratorGitea) GetAllBranchesOk() (*bool, bool)`

GetAllBranchesOk returns a tuple with the AllBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllBranches

`func (o *V1alpha1SCMProviderGeneratorGitea) SetAllBranches(v bool)`

SetAllBranches sets AllBranches field to given value.

### HasAllBranches

`func (o *V1alpha1SCMProviderGeneratorGitea) HasAllBranches() bool`

HasAllBranches returns a boolean if a field has been set.

### GetApi

`func (o *V1alpha1SCMProviderGeneratorGitea) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *V1alpha1SCMProviderGeneratorGitea) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *V1alpha1SCMProviderGeneratorGitea) SetApi(v string)`

SetApi sets Api field to given value.

### HasApi

`func (o *V1alpha1SCMProviderGeneratorGitea) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetInsecure

`func (o *V1alpha1SCMProviderGeneratorGitea) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *V1alpha1SCMProviderGeneratorGitea) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *V1alpha1SCMProviderGeneratorGitea) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *V1alpha1SCMProviderGeneratorGitea) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetOwner

`func (o *V1alpha1SCMProviderGeneratorGitea) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V1alpha1SCMProviderGeneratorGitea) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V1alpha1SCMProviderGeneratorGitea) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V1alpha1SCMProviderGeneratorGitea) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTokenRef

`func (o *V1alpha1SCMProviderGeneratorGitea) GetTokenRef() V1alpha1SecretRef`

GetTokenRef returns the TokenRef field if non-nil, zero value otherwise.

### GetTokenRefOk

`func (o *V1alpha1SCMProviderGeneratorGitea) GetTokenRefOk() (*V1alpha1SecretRef, bool)`

GetTokenRefOk returns a tuple with the TokenRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRef

`func (o *V1alpha1SCMProviderGeneratorGitea) SetTokenRef(v V1alpha1SecretRef)`

SetTokenRef sets TokenRef field to given value.

### HasTokenRef

`func (o *V1alpha1SCMProviderGeneratorGitea) HasTokenRef() bool`

HasTokenRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


