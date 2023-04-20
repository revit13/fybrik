# V1alpha1SCMProviderGeneratorAzureDevOps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenRef** | Pointer to [**V1alpha1SecretRef**](V1alpha1SecretRef.md) |  | [optional] 
**AllBranches** | Pointer to **bool** | Scan all branches instead of just the default branch. | [optional] 
**Api** | Pointer to **string** | The URL to Azure DevOps. If blank, use https://dev.azure.com. | [optional] 
**Organization** | Pointer to **string** | Azure Devops organization. Required. E.g. \&quot;my-organization\&quot;. | [optional] 
**TeamProject** | Pointer to **string** | Azure Devops team project. Required. E.g. \&quot;my-team\&quot;. | [optional] 

## Methods

### NewV1alpha1SCMProviderGeneratorAzureDevOps

`func NewV1alpha1SCMProviderGeneratorAzureDevOps() *V1alpha1SCMProviderGeneratorAzureDevOps`

NewV1alpha1SCMProviderGeneratorAzureDevOps instantiates a new V1alpha1SCMProviderGeneratorAzureDevOps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1SCMProviderGeneratorAzureDevOpsWithDefaults

`func NewV1alpha1SCMProviderGeneratorAzureDevOpsWithDefaults() *V1alpha1SCMProviderGeneratorAzureDevOps`

NewV1alpha1SCMProviderGeneratorAzureDevOpsWithDefaults instantiates a new V1alpha1SCMProviderGeneratorAzureDevOps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenRef

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) GetAccessTokenRef() V1alpha1SecretRef`

GetAccessTokenRef returns the AccessTokenRef field if non-nil, zero value otherwise.

### GetAccessTokenRefOk

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) GetAccessTokenRefOk() (*V1alpha1SecretRef, bool)`

GetAccessTokenRefOk returns a tuple with the AccessTokenRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenRef

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) SetAccessTokenRef(v V1alpha1SecretRef)`

SetAccessTokenRef sets AccessTokenRef field to given value.

### HasAccessTokenRef

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) HasAccessTokenRef() bool`

HasAccessTokenRef returns a boolean if a field has been set.

### GetAllBranches

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) GetAllBranches() bool`

GetAllBranches returns the AllBranches field if non-nil, zero value otherwise.

### GetAllBranchesOk

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) GetAllBranchesOk() (*bool, bool)`

GetAllBranchesOk returns a tuple with the AllBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllBranches

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) SetAllBranches(v bool)`

SetAllBranches sets AllBranches field to given value.

### HasAllBranches

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) HasAllBranches() bool`

HasAllBranches returns a boolean if a field has been set.

### GetApi

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) SetApi(v string)`

SetApi sets Api field to given value.

### HasApi

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetOrganization

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetTeamProject

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) GetTeamProject() string`

GetTeamProject returns the TeamProject field if non-nil, zero value otherwise.

### GetTeamProjectOk

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) GetTeamProjectOk() (*string, bool)`

GetTeamProjectOk returns a tuple with the TeamProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamProject

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) SetTeamProject(v string)`

SetTeamProject sets TeamProject field to given value.

### HasTeamProject

`func (o *V1alpha1SCMProviderGeneratorAzureDevOps) HasTeamProject() bool`

HasTeamProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


