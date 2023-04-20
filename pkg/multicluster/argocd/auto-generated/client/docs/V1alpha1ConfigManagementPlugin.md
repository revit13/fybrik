# V1alpha1ConfigManagementPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Generate** | Pointer to [**V1alpha1Command**](V1alpha1Command.md) |  | [optional] 
**Init** | Pointer to [**V1alpha1Command**](V1alpha1Command.md) |  | [optional] 
**LockRepo** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ConfigManagementPlugin

`func NewV1alpha1ConfigManagementPlugin() *V1alpha1ConfigManagementPlugin`

NewV1alpha1ConfigManagementPlugin instantiates a new V1alpha1ConfigManagementPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ConfigManagementPluginWithDefaults

`func NewV1alpha1ConfigManagementPluginWithDefaults() *V1alpha1ConfigManagementPlugin`

NewV1alpha1ConfigManagementPluginWithDefaults instantiates a new V1alpha1ConfigManagementPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerate

`func (o *V1alpha1ConfigManagementPlugin) GetGenerate() V1alpha1Command`

GetGenerate returns the Generate field if non-nil, zero value otherwise.

### GetGenerateOk

`func (o *V1alpha1ConfigManagementPlugin) GetGenerateOk() (*V1alpha1Command, bool)`

GetGenerateOk returns a tuple with the Generate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerate

`func (o *V1alpha1ConfigManagementPlugin) SetGenerate(v V1alpha1Command)`

SetGenerate sets Generate field to given value.

### HasGenerate

`func (o *V1alpha1ConfigManagementPlugin) HasGenerate() bool`

HasGenerate returns a boolean if a field has been set.

### GetInit

`func (o *V1alpha1ConfigManagementPlugin) GetInit() V1alpha1Command`

GetInit returns the Init field if non-nil, zero value otherwise.

### GetInitOk

`func (o *V1alpha1ConfigManagementPlugin) GetInitOk() (*V1alpha1Command, bool)`

GetInitOk returns a tuple with the Init field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInit

`func (o *V1alpha1ConfigManagementPlugin) SetInit(v V1alpha1Command)`

SetInit sets Init field to given value.

### HasInit

`func (o *V1alpha1ConfigManagementPlugin) HasInit() bool`

HasInit returns a boolean if a field has been set.

### GetLockRepo

`func (o *V1alpha1ConfigManagementPlugin) GetLockRepo() bool`

GetLockRepo returns the LockRepo field if non-nil, zero value otherwise.

### GetLockRepoOk

`func (o *V1alpha1ConfigManagementPlugin) GetLockRepoOk() (*bool, bool)`

GetLockRepoOk returns a tuple with the LockRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockRepo

`func (o *V1alpha1ConfigManagementPlugin) SetLockRepo(v bool)`

SetLockRepo sets LockRepo field to given value.

### HasLockRepo

`func (o *V1alpha1ConfigManagementPlugin) HasLockRepo() bool`

HasLockRepo returns a boolean if a field has been set.

### GetName

`func (o *V1alpha1ConfigManagementPlugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha1ConfigManagementPlugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha1ConfigManagementPlugin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1alpha1ConfigManagementPlugin) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


