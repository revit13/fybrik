# V1alpha1PullRequestGeneratorGitLab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Api** | Pointer to **string** | The GitLab API URL to talk to. If blank, uses https://gitlab.com/. | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Project** | Pointer to **string** | GitLab project to scan. Required. | [optional] 
**PullRequestState** | Pointer to **string** |  | [optional] 
**TokenRef** | Pointer to [**V1alpha1SecretRef**](V1alpha1SecretRef.md) |  | [optional] 

## Methods

### NewV1alpha1PullRequestGeneratorGitLab

`func NewV1alpha1PullRequestGeneratorGitLab() *V1alpha1PullRequestGeneratorGitLab`

NewV1alpha1PullRequestGeneratorGitLab instantiates a new V1alpha1PullRequestGeneratorGitLab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1PullRequestGeneratorGitLabWithDefaults

`func NewV1alpha1PullRequestGeneratorGitLabWithDefaults() *V1alpha1PullRequestGeneratorGitLab`

NewV1alpha1PullRequestGeneratorGitLabWithDefaults instantiates a new V1alpha1PullRequestGeneratorGitLab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApi

`func (o *V1alpha1PullRequestGeneratorGitLab) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *V1alpha1PullRequestGeneratorGitLab) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *V1alpha1PullRequestGeneratorGitLab) SetApi(v string)`

SetApi sets Api field to given value.

### HasApi

`func (o *V1alpha1PullRequestGeneratorGitLab) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetLabels

`func (o *V1alpha1PullRequestGeneratorGitLab) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1alpha1PullRequestGeneratorGitLab) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1alpha1PullRequestGeneratorGitLab) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1alpha1PullRequestGeneratorGitLab) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProject

`func (o *V1alpha1PullRequestGeneratorGitLab) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *V1alpha1PullRequestGeneratorGitLab) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *V1alpha1PullRequestGeneratorGitLab) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *V1alpha1PullRequestGeneratorGitLab) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPullRequestState

`func (o *V1alpha1PullRequestGeneratorGitLab) GetPullRequestState() string`

GetPullRequestState returns the PullRequestState field if non-nil, zero value otherwise.

### GetPullRequestStateOk

`func (o *V1alpha1PullRequestGeneratorGitLab) GetPullRequestStateOk() (*string, bool)`

GetPullRequestStateOk returns a tuple with the PullRequestState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestState

`func (o *V1alpha1PullRequestGeneratorGitLab) SetPullRequestState(v string)`

SetPullRequestState sets PullRequestState field to given value.

### HasPullRequestState

`func (o *V1alpha1PullRequestGeneratorGitLab) HasPullRequestState() bool`

HasPullRequestState returns a boolean if a field has been set.

### GetTokenRef

`func (o *V1alpha1PullRequestGeneratorGitLab) GetTokenRef() V1alpha1SecretRef`

GetTokenRef returns the TokenRef field if non-nil, zero value otherwise.

### GetTokenRefOk

`func (o *V1alpha1PullRequestGeneratorGitLab) GetTokenRefOk() (*V1alpha1SecretRef, bool)`

GetTokenRefOk returns a tuple with the TokenRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRef

`func (o *V1alpha1PullRequestGeneratorGitLab) SetTokenRef(v V1alpha1SecretRef)`

SetTokenRef sets TokenRef field to given value.

### HasTokenRef

`func (o *V1alpha1PullRequestGeneratorGitLab) HasTokenRef() bool`

HasTokenRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


