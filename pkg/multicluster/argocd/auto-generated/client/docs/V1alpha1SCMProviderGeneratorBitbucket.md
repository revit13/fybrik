# V1alpha1SCMProviderGeneratorBitbucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllBranches** | Pointer to **bool** | Scan all branches instead of just the main branch. | [optional] 
**AppPasswordRef** | Pointer to [**V1alpha1SecretRef**](V1alpha1SecretRef.md) |  | [optional] 
**Owner** | Pointer to **string** | Bitbucket workspace to scan. Required. | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1SCMProviderGeneratorBitbucket

`func NewV1alpha1SCMProviderGeneratorBitbucket() *V1alpha1SCMProviderGeneratorBitbucket`

NewV1alpha1SCMProviderGeneratorBitbucket instantiates a new V1alpha1SCMProviderGeneratorBitbucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1SCMProviderGeneratorBitbucketWithDefaults

`func NewV1alpha1SCMProviderGeneratorBitbucketWithDefaults() *V1alpha1SCMProviderGeneratorBitbucket`

NewV1alpha1SCMProviderGeneratorBitbucketWithDefaults instantiates a new V1alpha1SCMProviderGeneratorBitbucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllBranches

`func (o *V1alpha1SCMProviderGeneratorBitbucket) GetAllBranches() bool`

GetAllBranches returns the AllBranches field if non-nil, zero value otherwise.

### GetAllBranchesOk

`func (o *V1alpha1SCMProviderGeneratorBitbucket) GetAllBranchesOk() (*bool, bool)`

GetAllBranchesOk returns a tuple with the AllBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllBranches

`func (o *V1alpha1SCMProviderGeneratorBitbucket) SetAllBranches(v bool)`

SetAllBranches sets AllBranches field to given value.

### HasAllBranches

`func (o *V1alpha1SCMProviderGeneratorBitbucket) HasAllBranches() bool`

HasAllBranches returns a boolean if a field has been set.

### GetAppPasswordRef

`func (o *V1alpha1SCMProviderGeneratorBitbucket) GetAppPasswordRef() V1alpha1SecretRef`

GetAppPasswordRef returns the AppPasswordRef field if non-nil, zero value otherwise.

### GetAppPasswordRefOk

`func (o *V1alpha1SCMProviderGeneratorBitbucket) GetAppPasswordRefOk() (*V1alpha1SecretRef, bool)`

GetAppPasswordRefOk returns a tuple with the AppPasswordRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPasswordRef

`func (o *V1alpha1SCMProviderGeneratorBitbucket) SetAppPasswordRef(v V1alpha1SecretRef)`

SetAppPasswordRef sets AppPasswordRef field to given value.

### HasAppPasswordRef

`func (o *V1alpha1SCMProviderGeneratorBitbucket) HasAppPasswordRef() bool`

HasAppPasswordRef returns a boolean if a field has been set.

### GetOwner

`func (o *V1alpha1SCMProviderGeneratorBitbucket) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V1alpha1SCMProviderGeneratorBitbucket) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V1alpha1SCMProviderGeneratorBitbucket) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V1alpha1SCMProviderGeneratorBitbucket) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetUser

`func (o *V1alpha1SCMProviderGeneratorBitbucket) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V1alpha1SCMProviderGeneratorBitbucket) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V1alpha1SCMProviderGeneratorBitbucket) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *V1alpha1SCMProviderGeneratorBitbucket) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


