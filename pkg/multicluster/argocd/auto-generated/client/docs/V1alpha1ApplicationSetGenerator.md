# V1alpha1ApplicationSetGenerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterDecisionResource** | Pointer to [**V1alpha1DuckTypeGenerator**](V1alpha1DuckTypeGenerator.md) |  | [optional] 
**Clusters** | Pointer to [**V1alpha1ClusterGenerator**](V1alpha1ClusterGenerator.md) |  | [optional] 
**Git** | Pointer to [**V1alpha1GitGenerator**](V1alpha1GitGenerator.md) |  | [optional] 
**List** | Pointer to [**V1alpha1ListGenerator**](V1alpha1ListGenerator.md) |  | [optional] 
**Matrix** | Pointer to [**V1alpha1MatrixGenerator**](V1alpha1MatrixGenerator.md) |  | [optional] 
**Merge** | Pointer to [**V1alpha1MergeGenerator**](V1alpha1MergeGenerator.md) |  | [optional] 
**PullRequest** | Pointer to [**V1alpha1PullRequestGenerator**](V1alpha1PullRequestGenerator.md) |  | [optional] 
**ScmProvider** | Pointer to [**V1alpha1SCMProviderGenerator**](V1alpha1SCMProviderGenerator.md) |  | [optional] 
**Selector** | Pointer to [**V1LabelSelector**](V1LabelSelector.md) |  | [optional] 

## Methods

### NewV1alpha1ApplicationSetGenerator

`func NewV1alpha1ApplicationSetGenerator() *V1alpha1ApplicationSetGenerator`

NewV1alpha1ApplicationSetGenerator instantiates a new V1alpha1ApplicationSetGenerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSetGeneratorWithDefaults

`func NewV1alpha1ApplicationSetGeneratorWithDefaults() *V1alpha1ApplicationSetGenerator`

NewV1alpha1ApplicationSetGeneratorWithDefaults instantiates a new V1alpha1ApplicationSetGenerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterDecisionResource

`func (o *V1alpha1ApplicationSetGenerator) GetClusterDecisionResource() V1alpha1DuckTypeGenerator`

GetClusterDecisionResource returns the ClusterDecisionResource field if non-nil, zero value otherwise.

### GetClusterDecisionResourceOk

`func (o *V1alpha1ApplicationSetGenerator) GetClusterDecisionResourceOk() (*V1alpha1DuckTypeGenerator, bool)`

GetClusterDecisionResourceOk returns a tuple with the ClusterDecisionResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDecisionResource

`func (o *V1alpha1ApplicationSetGenerator) SetClusterDecisionResource(v V1alpha1DuckTypeGenerator)`

SetClusterDecisionResource sets ClusterDecisionResource field to given value.

### HasClusterDecisionResource

`func (o *V1alpha1ApplicationSetGenerator) HasClusterDecisionResource() bool`

HasClusterDecisionResource returns a boolean if a field has been set.

### GetClusters

`func (o *V1alpha1ApplicationSetGenerator) GetClusters() V1alpha1ClusterGenerator`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V1alpha1ApplicationSetGenerator) GetClustersOk() (*V1alpha1ClusterGenerator, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V1alpha1ApplicationSetGenerator) SetClusters(v V1alpha1ClusterGenerator)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V1alpha1ApplicationSetGenerator) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetGit

`func (o *V1alpha1ApplicationSetGenerator) GetGit() V1alpha1GitGenerator`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *V1alpha1ApplicationSetGenerator) GetGitOk() (*V1alpha1GitGenerator, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *V1alpha1ApplicationSetGenerator) SetGit(v V1alpha1GitGenerator)`

SetGit sets Git field to given value.

### HasGit

`func (o *V1alpha1ApplicationSetGenerator) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetList

`func (o *V1alpha1ApplicationSetGenerator) GetList() V1alpha1ListGenerator`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *V1alpha1ApplicationSetGenerator) GetListOk() (*V1alpha1ListGenerator, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *V1alpha1ApplicationSetGenerator) SetList(v V1alpha1ListGenerator)`

SetList sets List field to given value.

### HasList

`func (o *V1alpha1ApplicationSetGenerator) HasList() bool`

HasList returns a boolean if a field has been set.

### GetMatrix

`func (o *V1alpha1ApplicationSetGenerator) GetMatrix() V1alpha1MatrixGenerator`

GetMatrix returns the Matrix field if non-nil, zero value otherwise.

### GetMatrixOk

`func (o *V1alpha1ApplicationSetGenerator) GetMatrixOk() (*V1alpha1MatrixGenerator, bool)`

GetMatrixOk returns a tuple with the Matrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatrix

`func (o *V1alpha1ApplicationSetGenerator) SetMatrix(v V1alpha1MatrixGenerator)`

SetMatrix sets Matrix field to given value.

### HasMatrix

`func (o *V1alpha1ApplicationSetGenerator) HasMatrix() bool`

HasMatrix returns a boolean if a field has been set.

### GetMerge

`func (o *V1alpha1ApplicationSetGenerator) GetMerge() V1alpha1MergeGenerator`

GetMerge returns the Merge field if non-nil, zero value otherwise.

### GetMergeOk

`func (o *V1alpha1ApplicationSetGenerator) GetMergeOk() (*V1alpha1MergeGenerator, bool)`

GetMergeOk returns a tuple with the Merge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerge

`func (o *V1alpha1ApplicationSetGenerator) SetMerge(v V1alpha1MergeGenerator)`

SetMerge sets Merge field to given value.

### HasMerge

`func (o *V1alpha1ApplicationSetGenerator) HasMerge() bool`

HasMerge returns a boolean if a field has been set.

### GetPullRequest

`func (o *V1alpha1ApplicationSetGenerator) GetPullRequest() V1alpha1PullRequestGenerator`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *V1alpha1ApplicationSetGenerator) GetPullRequestOk() (*V1alpha1PullRequestGenerator, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *V1alpha1ApplicationSetGenerator) SetPullRequest(v V1alpha1PullRequestGenerator)`

SetPullRequest sets PullRequest field to given value.

### HasPullRequest

`func (o *V1alpha1ApplicationSetGenerator) HasPullRequest() bool`

HasPullRequest returns a boolean if a field has been set.

### GetScmProvider

`func (o *V1alpha1ApplicationSetGenerator) GetScmProvider() V1alpha1SCMProviderGenerator`

GetScmProvider returns the ScmProvider field if non-nil, zero value otherwise.

### GetScmProviderOk

`func (o *V1alpha1ApplicationSetGenerator) GetScmProviderOk() (*V1alpha1SCMProviderGenerator, bool)`

GetScmProviderOk returns a tuple with the ScmProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmProvider

`func (o *V1alpha1ApplicationSetGenerator) SetScmProvider(v V1alpha1SCMProviderGenerator)`

SetScmProvider sets ScmProvider field to given value.

### HasScmProvider

`func (o *V1alpha1ApplicationSetGenerator) HasScmProvider() bool`

HasScmProvider returns a boolean if a field has been set.

### GetSelector

`func (o *V1alpha1ApplicationSetGenerator) GetSelector() V1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1alpha1ApplicationSetGenerator) GetSelectorOk() (*V1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1alpha1ApplicationSetGenerator) SetSelector(v V1LabelSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *V1alpha1ApplicationSetGenerator) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


