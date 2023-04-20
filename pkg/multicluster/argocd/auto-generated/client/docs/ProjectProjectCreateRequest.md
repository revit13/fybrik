# ProjectProjectCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to [**V1alpha1AppProject**](V1alpha1AppProject.md) |  | [optional] 
**Upsert** | Pointer to **bool** |  | [optional] 

## Methods

### NewProjectProjectCreateRequest

`func NewProjectProjectCreateRequest() *ProjectProjectCreateRequest`

NewProjectProjectCreateRequest instantiates a new ProjectProjectCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectCreateRequestWithDefaults

`func NewProjectProjectCreateRequestWithDefaults() *ProjectProjectCreateRequest`

NewProjectProjectCreateRequestWithDefaults instantiates a new ProjectProjectCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ProjectProjectCreateRequest) GetProject() V1alpha1AppProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectProjectCreateRequest) GetProjectOk() (*V1alpha1AppProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectProjectCreateRequest) SetProject(v V1alpha1AppProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectProjectCreateRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetUpsert

`func (o *ProjectProjectCreateRequest) GetUpsert() bool`

GetUpsert returns the Upsert field if non-nil, zero value otherwise.

### GetUpsertOk

`func (o *ProjectProjectCreateRequest) GetUpsertOk() (*bool, bool)`

GetUpsertOk returns a tuple with the Upsert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsert

`func (o *ProjectProjectCreateRequest) SetUpsert(v bool)`

SetUpsert sets Upsert field to given value.

### HasUpsert

`func (o *ProjectProjectCreateRequest) HasUpsert() bool`

HasUpsert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


