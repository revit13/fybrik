# RepositoryRepoAppDetailsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**AppProject** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**V1alpha1ApplicationSource**](V1alpha1ApplicationSource.md) |  | [optional] 

## Methods

### NewRepositoryRepoAppDetailsQuery

`func NewRepositoryRepoAppDetailsQuery() *RepositoryRepoAppDetailsQuery`

NewRepositoryRepoAppDetailsQuery instantiates a new RepositoryRepoAppDetailsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryRepoAppDetailsQueryWithDefaults

`func NewRepositoryRepoAppDetailsQueryWithDefaults() *RepositoryRepoAppDetailsQuery`

NewRepositoryRepoAppDetailsQueryWithDefaults instantiates a new RepositoryRepoAppDetailsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *RepositoryRepoAppDetailsQuery) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *RepositoryRepoAppDetailsQuery) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *RepositoryRepoAppDetailsQuery) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *RepositoryRepoAppDetailsQuery) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAppProject

`func (o *RepositoryRepoAppDetailsQuery) GetAppProject() string`

GetAppProject returns the AppProject field if non-nil, zero value otherwise.

### GetAppProjectOk

`func (o *RepositoryRepoAppDetailsQuery) GetAppProjectOk() (*string, bool)`

GetAppProjectOk returns a tuple with the AppProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProject

`func (o *RepositoryRepoAppDetailsQuery) SetAppProject(v string)`

SetAppProject sets AppProject field to given value.

### HasAppProject

`func (o *RepositoryRepoAppDetailsQuery) HasAppProject() bool`

HasAppProject returns a boolean if a field has been set.

### GetSource

`func (o *RepositoryRepoAppDetailsQuery) GetSource() V1alpha1ApplicationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RepositoryRepoAppDetailsQuery) GetSourceOk() (*V1alpha1ApplicationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RepositoryRepoAppDetailsQuery) SetSource(v V1alpha1ApplicationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *RepositoryRepoAppDetailsQuery) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


