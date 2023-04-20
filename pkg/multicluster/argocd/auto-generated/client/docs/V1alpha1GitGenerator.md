# V1alpha1GitGenerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directories** | Pointer to [**[]V1alpha1GitDirectoryGeneratorItem**](V1alpha1GitDirectoryGeneratorItem.md) |  | [optional] 
**Files** | Pointer to [**[]V1alpha1GitFileGeneratorItem**](V1alpha1GitFileGeneratorItem.md) |  | [optional] 
**PathParamPrefix** | Pointer to **string** |  | [optional] 
**RepoURL** | Pointer to **string** |  | [optional] 
**RequeueAfterSeconds** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Template** | Pointer to [**V1alpha1ApplicationSetTemplate**](V1alpha1ApplicationSetTemplate.md) |  | [optional] 

## Methods

### NewV1alpha1GitGenerator

`func NewV1alpha1GitGenerator() *V1alpha1GitGenerator`

NewV1alpha1GitGenerator instantiates a new V1alpha1GitGenerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1GitGeneratorWithDefaults

`func NewV1alpha1GitGeneratorWithDefaults() *V1alpha1GitGenerator`

NewV1alpha1GitGeneratorWithDefaults instantiates a new V1alpha1GitGenerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectories

`func (o *V1alpha1GitGenerator) GetDirectories() []V1alpha1GitDirectoryGeneratorItem`

GetDirectories returns the Directories field if non-nil, zero value otherwise.

### GetDirectoriesOk

`func (o *V1alpha1GitGenerator) GetDirectoriesOk() (*[]V1alpha1GitDirectoryGeneratorItem, bool)`

GetDirectoriesOk returns a tuple with the Directories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectories

`func (o *V1alpha1GitGenerator) SetDirectories(v []V1alpha1GitDirectoryGeneratorItem)`

SetDirectories sets Directories field to given value.

### HasDirectories

`func (o *V1alpha1GitGenerator) HasDirectories() bool`

HasDirectories returns a boolean if a field has been set.

### GetFiles

`func (o *V1alpha1GitGenerator) GetFiles() []V1alpha1GitFileGeneratorItem`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *V1alpha1GitGenerator) GetFilesOk() (*[]V1alpha1GitFileGeneratorItem, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *V1alpha1GitGenerator) SetFiles(v []V1alpha1GitFileGeneratorItem)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *V1alpha1GitGenerator) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetPathParamPrefix

`func (o *V1alpha1GitGenerator) GetPathParamPrefix() string`

GetPathParamPrefix returns the PathParamPrefix field if non-nil, zero value otherwise.

### GetPathParamPrefixOk

`func (o *V1alpha1GitGenerator) GetPathParamPrefixOk() (*string, bool)`

GetPathParamPrefixOk returns a tuple with the PathParamPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathParamPrefix

`func (o *V1alpha1GitGenerator) SetPathParamPrefix(v string)`

SetPathParamPrefix sets PathParamPrefix field to given value.

### HasPathParamPrefix

`func (o *V1alpha1GitGenerator) HasPathParamPrefix() bool`

HasPathParamPrefix returns a boolean if a field has been set.

### GetRepoURL

`func (o *V1alpha1GitGenerator) GetRepoURL() string`

GetRepoURL returns the RepoURL field if non-nil, zero value otherwise.

### GetRepoURLOk

`func (o *V1alpha1GitGenerator) GetRepoURLOk() (*string, bool)`

GetRepoURLOk returns a tuple with the RepoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoURL

`func (o *V1alpha1GitGenerator) SetRepoURL(v string)`

SetRepoURL sets RepoURL field to given value.

### HasRepoURL

`func (o *V1alpha1GitGenerator) HasRepoURL() bool`

HasRepoURL returns a boolean if a field has been set.

### GetRequeueAfterSeconds

`func (o *V1alpha1GitGenerator) GetRequeueAfterSeconds() string`

GetRequeueAfterSeconds returns the RequeueAfterSeconds field if non-nil, zero value otherwise.

### GetRequeueAfterSecondsOk

`func (o *V1alpha1GitGenerator) GetRequeueAfterSecondsOk() (*string, bool)`

GetRequeueAfterSecondsOk returns a tuple with the RequeueAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequeueAfterSeconds

`func (o *V1alpha1GitGenerator) SetRequeueAfterSeconds(v string)`

SetRequeueAfterSeconds sets RequeueAfterSeconds field to given value.

### HasRequeueAfterSeconds

`func (o *V1alpha1GitGenerator) HasRequeueAfterSeconds() bool`

HasRequeueAfterSeconds returns a boolean if a field has been set.

### GetRevision

`func (o *V1alpha1GitGenerator) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1alpha1GitGenerator) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1alpha1GitGenerator) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1alpha1GitGenerator) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetTemplate

`func (o *V1alpha1GitGenerator) GetTemplate() V1alpha1ApplicationSetTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1alpha1GitGenerator) GetTemplateOk() (*V1alpha1ApplicationSetTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1alpha1GitGenerator) SetTemplate(v V1alpha1ApplicationSetTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *V1alpha1GitGenerator) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


