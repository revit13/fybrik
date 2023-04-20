# V1alpha1ApplicationSourcePlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to [**[]Applicationv1alpha1EnvEntry**](Applicationv1alpha1EnvEntry.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]V1alpha1ApplicationSourcePluginParameter**](V1alpha1ApplicationSourcePluginParameter.md) |  | [optional] 

## Methods

### NewV1alpha1ApplicationSourcePlugin

`func NewV1alpha1ApplicationSourcePlugin() *V1alpha1ApplicationSourcePlugin`

NewV1alpha1ApplicationSourcePlugin instantiates a new V1alpha1ApplicationSourcePlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSourcePluginWithDefaults

`func NewV1alpha1ApplicationSourcePluginWithDefaults() *V1alpha1ApplicationSourcePlugin`

NewV1alpha1ApplicationSourcePluginWithDefaults instantiates a new V1alpha1ApplicationSourcePlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *V1alpha1ApplicationSourcePlugin) GetEnv() []Applicationv1alpha1EnvEntry`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *V1alpha1ApplicationSourcePlugin) GetEnvOk() (*[]Applicationv1alpha1EnvEntry, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *V1alpha1ApplicationSourcePlugin) SetEnv(v []Applicationv1alpha1EnvEntry)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *V1alpha1ApplicationSourcePlugin) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetName

`func (o *V1alpha1ApplicationSourcePlugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha1ApplicationSourcePlugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha1ApplicationSourcePlugin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1alpha1ApplicationSourcePlugin) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *V1alpha1ApplicationSourcePlugin) GetParameters() []V1alpha1ApplicationSourcePluginParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *V1alpha1ApplicationSourcePlugin) GetParametersOk() (*[]V1alpha1ApplicationSourcePluginParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *V1alpha1ApplicationSourcePlugin) SetParameters(v []V1alpha1ApplicationSourcePluginParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *V1alpha1ApplicationSourcePlugin) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


