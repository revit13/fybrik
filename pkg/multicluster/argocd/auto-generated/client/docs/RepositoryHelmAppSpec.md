# RepositoryHelmAppSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileParameters** | Pointer to [**[]V1alpha1HelmFileParameter**](V1alpha1HelmFileParameter.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]V1alpha1HelmParameter**](V1alpha1HelmParameter.md) |  | [optional] 
**ValueFiles** | Pointer to **[]string** |  | [optional] 
**Values** | Pointer to **string** |  | [optional] 

## Methods

### NewRepositoryHelmAppSpec

`func NewRepositoryHelmAppSpec() *RepositoryHelmAppSpec`

NewRepositoryHelmAppSpec instantiates a new RepositoryHelmAppSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryHelmAppSpecWithDefaults

`func NewRepositoryHelmAppSpecWithDefaults() *RepositoryHelmAppSpec`

NewRepositoryHelmAppSpecWithDefaults instantiates a new RepositoryHelmAppSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileParameters

`func (o *RepositoryHelmAppSpec) GetFileParameters() []V1alpha1HelmFileParameter`

GetFileParameters returns the FileParameters field if non-nil, zero value otherwise.

### GetFileParametersOk

`func (o *RepositoryHelmAppSpec) GetFileParametersOk() (*[]V1alpha1HelmFileParameter, bool)`

GetFileParametersOk returns a tuple with the FileParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileParameters

`func (o *RepositoryHelmAppSpec) SetFileParameters(v []V1alpha1HelmFileParameter)`

SetFileParameters sets FileParameters field to given value.

### HasFileParameters

`func (o *RepositoryHelmAppSpec) HasFileParameters() bool`

HasFileParameters returns a boolean if a field has been set.

### GetName

`func (o *RepositoryHelmAppSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryHelmAppSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryHelmAppSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RepositoryHelmAppSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *RepositoryHelmAppSpec) GetParameters() []V1alpha1HelmParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *RepositoryHelmAppSpec) GetParametersOk() (*[]V1alpha1HelmParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *RepositoryHelmAppSpec) SetParameters(v []V1alpha1HelmParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *RepositoryHelmAppSpec) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetValueFiles

`func (o *RepositoryHelmAppSpec) GetValueFiles() []string`

GetValueFiles returns the ValueFiles field if non-nil, zero value otherwise.

### GetValueFilesOk

`func (o *RepositoryHelmAppSpec) GetValueFilesOk() (*[]string, bool)`

GetValueFilesOk returns a tuple with the ValueFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFiles

`func (o *RepositoryHelmAppSpec) SetValueFiles(v []string)`

SetValueFiles sets ValueFiles field to given value.

### HasValueFiles

`func (o *RepositoryHelmAppSpec) HasValueFiles() bool`

HasValueFiles returns a boolean if a field has been set.

### GetValues

`func (o *RepositoryHelmAppSpec) GetValues() string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RepositoryHelmAppSpec) GetValuesOk() (*string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RepositoryHelmAppSpec) SetValues(v string)`

SetValues sets Values field to given value.

### HasValues

`func (o *RepositoryHelmAppSpec) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


