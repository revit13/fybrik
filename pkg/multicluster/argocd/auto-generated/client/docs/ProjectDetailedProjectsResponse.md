# ProjectDetailedProjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]V1alpha1Cluster**](V1alpha1Cluster.md) |  | [optional] 
**GlobalProjects** | Pointer to [**[]V1alpha1AppProject**](V1alpha1AppProject.md) |  | [optional] 
**Project** | Pointer to [**V1alpha1AppProject**](V1alpha1AppProject.md) |  | [optional] 
**Repositories** | Pointer to [**[]V1alpha1Repository**](V1alpha1Repository.md) |  | [optional] 

## Methods

### NewProjectDetailedProjectsResponse

`func NewProjectDetailedProjectsResponse() *ProjectDetailedProjectsResponse`

NewProjectDetailedProjectsResponse instantiates a new ProjectDetailedProjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDetailedProjectsResponseWithDefaults

`func NewProjectDetailedProjectsResponseWithDefaults() *ProjectDetailedProjectsResponse`

NewProjectDetailedProjectsResponseWithDefaults instantiates a new ProjectDetailedProjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *ProjectDetailedProjectsResponse) GetClusters() []V1alpha1Cluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *ProjectDetailedProjectsResponse) GetClustersOk() (*[]V1alpha1Cluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *ProjectDetailedProjectsResponse) SetClusters(v []V1alpha1Cluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *ProjectDetailedProjectsResponse) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetGlobalProjects

`func (o *ProjectDetailedProjectsResponse) GetGlobalProjects() []V1alpha1AppProject`

GetGlobalProjects returns the GlobalProjects field if non-nil, zero value otherwise.

### GetGlobalProjectsOk

`func (o *ProjectDetailedProjectsResponse) GetGlobalProjectsOk() (*[]V1alpha1AppProject, bool)`

GetGlobalProjectsOk returns a tuple with the GlobalProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalProjects

`func (o *ProjectDetailedProjectsResponse) SetGlobalProjects(v []V1alpha1AppProject)`

SetGlobalProjects sets GlobalProjects field to given value.

### HasGlobalProjects

`func (o *ProjectDetailedProjectsResponse) HasGlobalProjects() bool`

HasGlobalProjects returns a boolean if a field has been set.

### GetProject

`func (o *ProjectDetailedProjectsResponse) GetProject() V1alpha1AppProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectDetailedProjectsResponse) GetProjectOk() (*V1alpha1AppProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectDetailedProjectsResponse) SetProject(v V1alpha1AppProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectDetailedProjectsResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRepositories

`func (o *ProjectDetailedProjectsResponse) GetRepositories() []V1alpha1Repository`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *ProjectDetailedProjectsResponse) GetRepositoriesOk() (*[]V1alpha1Repository, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *ProjectDetailedProjectsResponse) SetRepositories(v []V1alpha1Repository)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *ProjectDetailedProjectsResponse) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


