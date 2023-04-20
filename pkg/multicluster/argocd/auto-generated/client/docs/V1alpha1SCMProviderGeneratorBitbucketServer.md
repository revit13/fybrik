# V1alpha1SCMProviderGeneratorBitbucketServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllBranches** | Pointer to **bool** | Scan all branches instead of just the default branch. | [optional] 
**Api** | Pointer to **string** | The Bitbucket Server REST API URL to talk to. Required. | [optional] 
**BasicAuth** | Pointer to [**V1alpha1BasicAuthBitbucketServer**](V1alpha1BasicAuthBitbucketServer.md) |  | [optional] 
**Project** | Pointer to **string** | Project to scan. Required. | [optional] 

## Methods

### NewV1alpha1SCMProviderGeneratorBitbucketServer

`func NewV1alpha1SCMProviderGeneratorBitbucketServer() *V1alpha1SCMProviderGeneratorBitbucketServer`

NewV1alpha1SCMProviderGeneratorBitbucketServer instantiates a new V1alpha1SCMProviderGeneratorBitbucketServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1SCMProviderGeneratorBitbucketServerWithDefaults

`func NewV1alpha1SCMProviderGeneratorBitbucketServerWithDefaults() *V1alpha1SCMProviderGeneratorBitbucketServer`

NewV1alpha1SCMProviderGeneratorBitbucketServerWithDefaults instantiates a new V1alpha1SCMProviderGeneratorBitbucketServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllBranches

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetAllBranches() bool`

GetAllBranches returns the AllBranches field if non-nil, zero value otherwise.

### GetAllBranchesOk

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetAllBranchesOk() (*bool, bool)`

GetAllBranchesOk returns a tuple with the AllBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllBranches

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) SetAllBranches(v bool)`

SetAllBranches sets AllBranches field to given value.

### HasAllBranches

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) HasAllBranches() bool`

HasAllBranches returns a boolean if a field has been set.

### GetApi

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) SetApi(v string)`

SetApi sets Api field to given value.

### HasApi

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetBasicAuth

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetBasicAuth() V1alpha1BasicAuthBitbucketServer`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetBasicAuthOk() (*V1alpha1BasicAuthBitbucketServer, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) SetBasicAuth(v V1alpha1BasicAuthBitbucketServer)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetProject

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *V1alpha1SCMProviderGeneratorBitbucketServer) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


