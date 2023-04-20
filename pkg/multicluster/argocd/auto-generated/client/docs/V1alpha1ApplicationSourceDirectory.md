# V1alpha1ApplicationSourceDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclude** | Pointer to **string** |  | [optional] 
**Include** | Pointer to **string** |  | [optional] 
**Jsonnet** | Pointer to [**V1alpha1ApplicationSourceJsonnet**](V1alpha1ApplicationSourceJsonnet.md) |  | [optional] 
**Recurse** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1alpha1ApplicationSourceDirectory

`func NewV1alpha1ApplicationSourceDirectory() *V1alpha1ApplicationSourceDirectory`

NewV1alpha1ApplicationSourceDirectory instantiates a new V1alpha1ApplicationSourceDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSourceDirectoryWithDefaults

`func NewV1alpha1ApplicationSourceDirectoryWithDefaults() *V1alpha1ApplicationSourceDirectory`

NewV1alpha1ApplicationSourceDirectoryWithDefaults instantiates a new V1alpha1ApplicationSourceDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclude

`func (o *V1alpha1ApplicationSourceDirectory) GetExclude() string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *V1alpha1ApplicationSourceDirectory) GetExcludeOk() (*string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *V1alpha1ApplicationSourceDirectory) SetExclude(v string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *V1alpha1ApplicationSourceDirectory) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetInclude

`func (o *V1alpha1ApplicationSourceDirectory) GetInclude() string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *V1alpha1ApplicationSourceDirectory) GetIncludeOk() (*string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *V1alpha1ApplicationSourceDirectory) SetInclude(v string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *V1alpha1ApplicationSourceDirectory) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### GetJsonnet

`func (o *V1alpha1ApplicationSourceDirectory) GetJsonnet() V1alpha1ApplicationSourceJsonnet`

GetJsonnet returns the Jsonnet field if non-nil, zero value otherwise.

### GetJsonnetOk

`func (o *V1alpha1ApplicationSourceDirectory) GetJsonnetOk() (*V1alpha1ApplicationSourceJsonnet, bool)`

GetJsonnetOk returns a tuple with the Jsonnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonnet

`func (o *V1alpha1ApplicationSourceDirectory) SetJsonnet(v V1alpha1ApplicationSourceJsonnet)`

SetJsonnet sets Jsonnet field to given value.

### HasJsonnet

`func (o *V1alpha1ApplicationSourceDirectory) HasJsonnet() bool`

HasJsonnet returns a boolean if a field has been set.

### GetRecurse

`func (o *V1alpha1ApplicationSourceDirectory) GetRecurse() bool`

GetRecurse returns the Recurse field if non-nil, zero value otherwise.

### GetRecurseOk

`func (o *V1alpha1ApplicationSourceDirectory) GetRecurseOk() (*bool, bool)`

GetRecurseOk returns a tuple with the Recurse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurse

`func (o *V1alpha1ApplicationSourceDirectory) SetRecurse(v bool)`

SetRecurse sets Recurse field to given value.

### HasRecurse

`func (o *V1alpha1ApplicationSourceDirectory) HasRecurse() bool`

HasRecurse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


