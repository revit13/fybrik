# V1alpha1SCMProviderGeneratorFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchMatch** | Pointer to **string** | A regex which must match the branch name. | [optional] 
**LabelMatch** | Pointer to **string** | A regex which must match at least one label. | [optional] 
**PathsDoNotExist** | Pointer to **[]string** | An array of paths, all of which must not exist. | [optional] 
**PathsExist** | Pointer to **[]string** | An array of paths, all of which must exist. | [optional] 
**RepositoryMatch** | Pointer to **string** | A regex for repo names. | [optional] 

## Methods

### NewV1alpha1SCMProviderGeneratorFilter

`func NewV1alpha1SCMProviderGeneratorFilter() *V1alpha1SCMProviderGeneratorFilter`

NewV1alpha1SCMProviderGeneratorFilter instantiates a new V1alpha1SCMProviderGeneratorFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1SCMProviderGeneratorFilterWithDefaults

`func NewV1alpha1SCMProviderGeneratorFilterWithDefaults() *V1alpha1SCMProviderGeneratorFilter`

NewV1alpha1SCMProviderGeneratorFilterWithDefaults instantiates a new V1alpha1SCMProviderGeneratorFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchMatch

`func (o *V1alpha1SCMProviderGeneratorFilter) GetBranchMatch() string`

GetBranchMatch returns the BranchMatch field if non-nil, zero value otherwise.

### GetBranchMatchOk

`func (o *V1alpha1SCMProviderGeneratorFilter) GetBranchMatchOk() (*string, bool)`

GetBranchMatchOk returns a tuple with the BranchMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchMatch

`func (o *V1alpha1SCMProviderGeneratorFilter) SetBranchMatch(v string)`

SetBranchMatch sets BranchMatch field to given value.

### HasBranchMatch

`func (o *V1alpha1SCMProviderGeneratorFilter) HasBranchMatch() bool`

HasBranchMatch returns a boolean if a field has been set.

### GetLabelMatch

`func (o *V1alpha1SCMProviderGeneratorFilter) GetLabelMatch() string`

GetLabelMatch returns the LabelMatch field if non-nil, zero value otherwise.

### GetLabelMatchOk

`func (o *V1alpha1SCMProviderGeneratorFilter) GetLabelMatchOk() (*string, bool)`

GetLabelMatchOk returns a tuple with the LabelMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelMatch

`func (o *V1alpha1SCMProviderGeneratorFilter) SetLabelMatch(v string)`

SetLabelMatch sets LabelMatch field to given value.

### HasLabelMatch

`func (o *V1alpha1SCMProviderGeneratorFilter) HasLabelMatch() bool`

HasLabelMatch returns a boolean if a field has been set.

### GetPathsDoNotExist

`func (o *V1alpha1SCMProviderGeneratorFilter) GetPathsDoNotExist() []string`

GetPathsDoNotExist returns the PathsDoNotExist field if non-nil, zero value otherwise.

### GetPathsDoNotExistOk

`func (o *V1alpha1SCMProviderGeneratorFilter) GetPathsDoNotExistOk() (*[]string, bool)`

GetPathsDoNotExistOk returns a tuple with the PathsDoNotExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathsDoNotExist

`func (o *V1alpha1SCMProviderGeneratorFilter) SetPathsDoNotExist(v []string)`

SetPathsDoNotExist sets PathsDoNotExist field to given value.

### HasPathsDoNotExist

`func (o *V1alpha1SCMProviderGeneratorFilter) HasPathsDoNotExist() bool`

HasPathsDoNotExist returns a boolean if a field has been set.

### GetPathsExist

`func (o *V1alpha1SCMProviderGeneratorFilter) GetPathsExist() []string`

GetPathsExist returns the PathsExist field if non-nil, zero value otherwise.

### GetPathsExistOk

`func (o *V1alpha1SCMProviderGeneratorFilter) GetPathsExistOk() (*[]string, bool)`

GetPathsExistOk returns a tuple with the PathsExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathsExist

`func (o *V1alpha1SCMProviderGeneratorFilter) SetPathsExist(v []string)`

SetPathsExist sets PathsExist field to given value.

### HasPathsExist

`func (o *V1alpha1SCMProviderGeneratorFilter) HasPathsExist() bool`

HasPathsExist returns a boolean if a field has been set.

### GetRepositoryMatch

`func (o *V1alpha1SCMProviderGeneratorFilter) GetRepositoryMatch() string`

GetRepositoryMatch returns the RepositoryMatch field if non-nil, zero value otherwise.

### GetRepositoryMatchOk

`func (o *V1alpha1SCMProviderGeneratorFilter) GetRepositoryMatchOk() (*string, bool)`

GetRepositoryMatchOk returns a tuple with the RepositoryMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryMatch

`func (o *V1alpha1SCMProviderGeneratorFilter) SetRepositoryMatch(v string)`

SetRepositoryMatch sets RepositoryMatch field to given value.

### HasRepositoryMatch

`func (o *V1alpha1SCMProviderGeneratorFilter) HasRepositoryMatch() bool`

HasRepositoryMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


