# V1alpha1ApplicationSetNestedGenerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterDecisionResource** | Pointer to [**V1alpha1DuckTypeGenerator**](V1alpha1DuckTypeGenerator.md) |  | [optional] 
**Clusters** | Pointer to [**V1alpha1ClusterGenerator**](V1alpha1ClusterGenerator.md) |  | [optional] 
**Git** | Pointer to [**V1alpha1GitGenerator**](V1alpha1GitGenerator.md) |  | [optional] 
**List** | Pointer to [**V1alpha1ListGenerator**](V1alpha1ListGenerator.md) |  | [optional] 
**Matrix** | Pointer to [**V1JSON**](V1JSON.md) |  | [optional] 
**Merge** | Pointer to [**V1JSON**](V1JSON.md) |  | [optional] 
**PullRequest** | Pointer to [**V1alpha1PullRequestGenerator**](V1alpha1PullRequestGenerator.md) |  | [optional] 
**ScmProvider** | Pointer to [**V1alpha1SCMProviderGenerator**](V1alpha1SCMProviderGenerator.md) |  | [optional] 
**Selector** | Pointer to [**V1LabelSelector**](V1LabelSelector.md) |  | [optional] 

## Methods

### NewV1alpha1ApplicationSetNestedGenerator

`func NewV1alpha1ApplicationSetNestedGenerator() *V1alpha1ApplicationSetNestedGenerator`

NewV1alpha1ApplicationSetNestedGenerator instantiates a new V1alpha1ApplicationSetNestedGenerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSetNestedGeneratorWithDefaults

`func NewV1alpha1ApplicationSetNestedGeneratorWithDefaults() *V1alpha1ApplicationSetNestedGenerator`

NewV1alpha1ApplicationSetNestedGeneratorWithDefaults instantiates a new V1alpha1ApplicationSetNestedGenerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterDecisionResource

`func (o *V1alpha1ApplicationSetNestedGenerator) GetClusterDecisionResource() V1alpha1DuckTypeGenerator`

GetClusterDecisionResource returns the ClusterDecisionResource field if non-nil, zero value otherwise.

### GetClusterDecisionResourceOk

`func (o *V1alpha1ApplicationSetNestedGenerator) GetClusterDecisionResourceOk() (*V1alpha1DuckTypeGenerator, bool)`

GetClusterDecisionResourceOk returns a tuple with the ClusterDecisionResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDecisionResource

`func (o *V1alpha1ApplicationSetNestedGenerator) SetClusterDecisionResource(v V1alpha1DuckTypeGenerator)`

SetClusterDecisionResource sets ClusterDecisionResource field to given value.

### HasClusterDecisionResource

`func (o *V1alpha1ApplicationSetNestedGenerator) HasClusterDecisionResource() bool`

HasClusterDecisionResource returns a boolean if a field has been set.

### GetClusters

`func (o *V1alpha1ApplicationSetNestedGenerator) GetClusters() V1alpha1ClusterGenerator`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V1alpha1ApplicationSetNestedGenerator) GetClustersOk() (*V1alpha1ClusterGenerator, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V1alpha1ApplicationSetNestedGenerator) SetClusters(v V1alpha1ClusterGenerator)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V1alpha1ApplicationSetNestedGenerator) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetGit

`func (o *V1alpha1ApplicationSetNestedGenerator) GetGit() V1alpha1GitGenerator`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *V1alpha1ApplicationSetNestedGenerator) GetGitOk() (*V1alpha1GitGenerator, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *V1alpha1ApplicationSetNestedGenerator) SetGit(v V1alpha1GitGenerator)`

SetGit sets Git field to given value.

### HasGit

`func (o *V1alpha1ApplicationSetNestedGenerator) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetList

`func (o *V1alpha1ApplicationSetNestedGenerator) GetList() V1alpha1ListGenerator`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *V1alpha1ApplicationSetNestedGenerator) GetListOk() (*V1alpha1ListGenerator, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *V1alpha1ApplicationSetNestedGenerator) SetList(v V1alpha1ListGenerator)`

SetList sets List field to given value.

### HasList

`func (o *V1alpha1ApplicationSetNestedGenerator) HasList() bool`

HasList returns a boolean if a field has been set.

### GetMatrix

`func (o *V1alpha1ApplicationSetNestedGenerator) GetMatrix() V1JSON`

GetMatrix returns the Matrix field if non-nil, zero value otherwise.

### GetMatrixOk

`func (o *V1alpha1ApplicationSetNestedGenerator) GetMatrixOk() (*V1JSON, bool)`

GetMatrixOk returns a tuple with the Matrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatrix

`func (o *V1alpha1ApplicationSetNestedGenerator) SetMatrix(v V1JSON)`

SetMatrix sets Matrix field to given value.

### HasMatrix

`func (o *V1alpha1ApplicationSetNestedGenerator) HasMatrix() bool`

HasMatrix returns a boolean if a field has been set.

### GetMerge

`func (o *V1alpha1ApplicationSetNestedGenerator) GetMerge() V1JSON`

GetMerge returns the Merge field if non-nil, zero value otherwise.

### GetMergeOk

`func (o *V1alpha1ApplicationSetNestedGenerator) GetMergeOk() (*V1JSON, bool)`

GetMergeOk returns a tuple with the Merge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerge

`func (o *V1alpha1ApplicationSetNestedGenerator) SetMerge(v V1JSON)`

SetMerge sets Merge field to given value.

### HasMerge

`func (o *V1alpha1ApplicationSetNestedGenerator) HasMerge() bool`

HasMerge returns a boolean if a field has been set.

### GetPullRequest

`func (o *V1alpha1ApplicationSetNestedGenerator) GetPullRequest() V1alpha1PullRequestGenerator`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *V1alpha1ApplicationSetNestedGenerator) GetPullRequestOk() (*V1alpha1PullRequestGenerator, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *V1alpha1ApplicationSetNestedGenerator) SetPullRequest(v V1alpha1PullRequestGenerator)`

SetPullRequest sets PullRequest field to given value.

### HasPullRequest

`func (o *V1alpha1ApplicationSetNestedGenerator) HasPullRequest() bool`

HasPullRequest returns a boolean if a field has been set.

### GetScmProvider

`func (o *V1alpha1ApplicationSetNestedGenerator) GetScmProvider() V1alpha1SCMProviderGenerator`

GetScmProvider returns the ScmProvider field if non-nil, zero value otherwise.

### GetScmProviderOk

`func (o *V1alpha1ApplicationSetNestedGenerator) GetScmProviderOk() (*V1alpha1SCMProviderGenerator, bool)`

GetScmProviderOk returns a tuple with the ScmProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmProvider

`func (o *V1alpha1ApplicationSetNestedGenerator) SetScmProvider(v V1alpha1SCMProviderGenerator)`

SetScmProvider sets ScmProvider field to given value.

### HasScmProvider

`func (o *V1alpha1ApplicationSetNestedGenerator) HasScmProvider() bool`

HasScmProvider returns a boolean if a field has been set.

### GetSelector

`func (o *V1alpha1ApplicationSetNestedGenerator) GetSelector() V1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1alpha1ApplicationSetNestedGenerator) GetSelectorOk() (*V1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1alpha1ApplicationSetNestedGenerator) SetSelector(v V1LabelSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *V1alpha1ApplicationSetNestedGenerator) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


