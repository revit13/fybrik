/*
Consolidate Services

Description of all APIs

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
)

// ProjectDetailedProjectsResponse struct for ProjectDetailedProjectsResponse
type ProjectDetailedProjectsResponse struct {
	Clusters *[]V1alpha1Cluster `json:"clusters,omitempty"`
	GlobalProjects *[]V1alpha1AppProject `json:"globalProjects,omitempty"`
	Project *V1alpha1AppProject `json:"project,omitempty"`
	Repositories *[]V1alpha1Repository `json:"repositories,omitempty"`
}

// NewProjectDetailedProjectsResponse instantiates a new ProjectDetailedProjectsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDetailedProjectsResponse() *ProjectDetailedProjectsResponse {
	this := ProjectDetailedProjectsResponse{}
	return &this
}

// NewProjectDetailedProjectsResponseWithDefaults instantiates a new ProjectDetailedProjectsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDetailedProjectsResponseWithDefaults() *ProjectDetailedProjectsResponse {
	this := ProjectDetailedProjectsResponse{}
	return &this
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *ProjectDetailedProjectsResponse) GetClusters() []V1alpha1Cluster {
	if o == nil || o.Clusters == nil {
		var ret []V1alpha1Cluster
		return ret
	}
	return *o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetailedProjectsResponse) GetClustersOk() (*[]V1alpha1Cluster, bool) {
	if o == nil || o.Clusters == nil {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *ProjectDetailedProjectsResponse) HasClusters() bool {
	if o != nil && o.Clusters != nil {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []V1alpha1Cluster and assigns it to the Clusters field.
func (o *ProjectDetailedProjectsResponse) SetClusters(v []V1alpha1Cluster) {
	o.Clusters = &v
}

// GetGlobalProjects returns the GlobalProjects field value if set, zero value otherwise.
func (o *ProjectDetailedProjectsResponse) GetGlobalProjects() []V1alpha1AppProject {
	if o == nil || o.GlobalProjects == nil {
		var ret []V1alpha1AppProject
		return ret
	}
	return *o.GlobalProjects
}

// GetGlobalProjectsOk returns a tuple with the GlobalProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetailedProjectsResponse) GetGlobalProjectsOk() (*[]V1alpha1AppProject, bool) {
	if o == nil || o.GlobalProjects == nil {
		return nil, false
	}
	return o.GlobalProjects, true
}

// HasGlobalProjects returns a boolean if a field has been set.
func (o *ProjectDetailedProjectsResponse) HasGlobalProjects() bool {
	if o != nil && o.GlobalProjects != nil {
		return true
	}

	return false
}

// SetGlobalProjects gets a reference to the given []V1alpha1AppProject and assigns it to the GlobalProjects field.
func (o *ProjectDetailedProjectsResponse) SetGlobalProjects(v []V1alpha1AppProject) {
	o.GlobalProjects = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ProjectDetailedProjectsResponse) GetProject() V1alpha1AppProject {
	if o == nil || o.Project == nil {
		var ret V1alpha1AppProject
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetailedProjectsResponse) GetProjectOk() (*V1alpha1AppProject, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProjectDetailedProjectsResponse) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given V1alpha1AppProject and assigns it to the Project field.
func (o *ProjectDetailedProjectsResponse) SetProject(v V1alpha1AppProject) {
	o.Project = &v
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *ProjectDetailedProjectsResponse) GetRepositories() []V1alpha1Repository {
	if o == nil || o.Repositories == nil {
		var ret []V1alpha1Repository
		return ret
	}
	return *o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetailedProjectsResponse) GetRepositoriesOk() (*[]V1alpha1Repository, bool) {
	if o == nil || o.Repositories == nil {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *ProjectDetailedProjectsResponse) HasRepositories() bool {
	if o != nil && o.Repositories != nil {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given []V1alpha1Repository and assigns it to the Repositories field.
func (o *ProjectDetailedProjectsResponse) SetRepositories(v []V1alpha1Repository) {
	o.Repositories = &v
}

func (o ProjectDetailedProjectsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Clusters != nil {
		toSerialize["clusters"] = o.Clusters
	}
	if o.GlobalProjects != nil {
		toSerialize["globalProjects"] = o.GlobalProjects
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Repositories != nil {
		toSerialize["repositories"] = o.Repositories
	}
	return json.Marshal(toSerialize)
}

type NullableProjectDetailedProjectsResponse struct {
	value *ProjectDetailedProjectsResponse
	isSet bool
}

func (v NullableProjectDetailedProjectsResponse) Get() *ProjectDetailedProjectsResponse {
	return v.value
}

func (v *NullableProjectDetailedProjectsResponse) Set(val *ProjectDetailedProjectsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDetailedProjectsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDetailedProjectsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDetailedProjectsResponse(val *ProjectDetailedProjectsResponse) *NullableProjectDetailedProjectsResponse {
	return &NullableProjectDetailedProjectsResponse{value: val, isSet: true}
}

func (v NullableProjectDetailedProjectsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDetailedProjectsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


