# V1alpha1ApplicationSourceHelm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileParameters** | Pointer to [**[]V1alpha1HelmFileParameter**](V1alpha1HelmFileParameter.md) |  | [optional] 
**IgnoreMissingValueFiles** | Pointer to **bool** |  | [optional] 
**Parameters** | Pointer to [**[]V1alpha1HelmParameter**](V1alpha1HelmParameter.md) |  | [optional] 
**PassCredentials** | Pointer to **bool** |  | [optional] 
**ReleaseName** | Pointer to **string** |  | [optional] 
**SkipCrds** | Pointer to **bool** |  | [optional] 
**ValueFiles** | Pointer to **[]string** |  | [optional] 
**Values** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ApplicationSourceHelm

`func NewV1alpha1ApplicationSourceHelm() *V1alpha1ApplicationSourceHelm`

NewV1alpha1ApplicationSourceHelm instantiates a new V1alpha1ApplicationSourceHelm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSourceHelmWithDefaults

`func NewV1alpha1ApplicationSourceHelmWithDefaults() *V1alpha1ApplicationSourceHelm`

NewV1alpha1ApplicationSourceHelmWithDefaults instantiates a new V1alpha1ApplicationSourceHelm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileParameters

`func (o *V1alpha1ApplicationSourceHelm) GetFileParameters() []V1alpha1HelmFileParameter`

GetFileParameters returns the FileParameters field if non-nil, zero value otherwise.

### GetFileParametersOk

`func (o *V1alpha1ApplicationSourceHelm) GetFileParametersOk() (*[]V1alpha1HelmFileParameter, bool)`

GetFileParametersOk returns a tuple with the FileParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileParameters

`func (o *V1alpha1ApplicationSourceHelm) SetFileParameters(v []V1alpha1HelmFileParameter)`

SetFileParameters sets FileParameters field to given value.

### HasFileParameters

`func (o *V1alpha1ApplicationSourceHelm) HasFileParameters() bool`

HasFileParameters returns a boolean if a field has been set.

### GetIgnoreMissingValueFiles

`func (o *V1alpha1ApplicationSourceHelm) GetIgnoreMissingValueFiles() bool`

GetIgnoreMissingValueFiles returns the IgnoreMissingValueFiles field if non-nil, zero value otherwise.

### GetIgnoreMissingValueFilesOk

`func (o *V1alpha1ApplicationSourceHelm) GetIgnoreMissingValueFilesOk() (*bool, bool)`

GetIgnoreMissingValueFilesOk returns a tuple with the IgnoreMissingValueFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMissingValueFiles

`func (o *V1alpha1ApplicationSourceHelm) SetIgnoreMissingValueFiles(v bool)`

SetIgnoreMissingValueFiles sets IgnoreMissingValueFiles field to given value.

### HasIgnoreMissingValueFiles

`func (o *V1alpha1ApplicationSourceHelm) HasIgnoreMissingValueFiles() bool`

HasIgnoreMissingValueFiles returns a boolean if a field has been set.

### GetParameters

`func (o *V1alpha1ApplicationSourceHelm) GetParameters() []V1alpha1HelmParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *V1alpha1ApplicationSourceHelm) GetParametersOk() (*[]V1alpha1HelmParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *V1alpha1ApplicationSourceHelm) SetParameters(v []V1alpha1HelmParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *V1alpha1ApplicationSourceHelm) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPassCredentials

`func (o *V1alpha1ApplicationSourceHelm) GetPassCredentials() bool`

GetPassCredentials returns the PassCredentials field if non-nil, zero value otherwise.

### GetPassCredentialsOk

`func (o *V1alpha1ApplicationSourceHelm) GetPassCredentialsOk() (*bool, bool)`

GetPassCredentialsOk returns a tuple with the PassCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCredentials

`func (o *V1alpha1ApplicationSourceHelm) SetPassCredentials(v bool)`

SetPassCredentials sets PassCredentials field to given value.

### HasPassCredentials

`func (o *V1alpha1ApplicationSourceHelm) HasPassCredentials() bool`

HasPassCredentials returns a boolean if a field has been set.

### GetReleaseName

`func (o *V1alpha1ApplicationSourceHelm) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *V1alpha1ApplicationSourceHelm) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *V1alpha1ApplicationSourceHelm) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *V1alpha1ApplicationSourceHelm) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetSkipCrds

`func (o *V1alpha1ApplicationSourceHelm) GetSkipCrds() bool`

GetSkipCrds returns the SkipCrds field if non-nil, zero value otherwise.

### GetSkipCrdsOk

`func (o *V1alpha1ApplicationSourceHelm) GetSkipCrdsOk() (*bool, bool)`

GetSkipCrdsOk returns a tuple with the SkipCrds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCrds

`func (o *V1alpha1ApplicationSourceHelm) SetSkipCrds(v bool)`

SetSkipCrds sets SkipCrds field to given value.

### HasSkipCrds

`func (o *V1alpha1ApplicationSourceHelm) HasSkipCrds() bool`

HasSkipCrds returns a boolean if a field has been set.

### GetValueFiles

`func (o *V1alpha1ApplicationSourceHelm) GetValueFiles() []string`

GetValueFiles returns the ValueFiles field if non-nil, zero value otherwise.

### GetValueFilesOk

`func (o *V1alpha1ApplicationSourceHelm) GetValueFilesOk() (*[]string, bool)`

GetValueFilesOk returns a tuple with the ValueFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFiles

`func (o *V1alpha1ApplicationSourceHelm) SetValueFiles(v []string)`

SetValueFiles sets ValueFiles field to given value.

### HasValueFiles

`func (o *V1alpha1ApplicationSourceHelm) HasValueFiles() bool`

HasValueFiles returns a boolean if a field has been set.

### GetValues

`func (o *V1alpha1ApplicationSourceHelm) GetValues() string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *V1alpha1ApplicationSourceHelm) GetValuesOk() (*string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *V1alpha1ApplicationSourceHelm) SetValues(v string)`

SetValues sets Values field to given value.

### HasValues

`func (o *V1alpha1ApplicationSourceHelm) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetVersion

`func (o *V1alpha1ApplicationSourceHelm) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V1alpha1ApplicationSourceHelm) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V1alpha1ApplicationSourceHelm) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V1alpha1ApplicationSourceHelm) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


