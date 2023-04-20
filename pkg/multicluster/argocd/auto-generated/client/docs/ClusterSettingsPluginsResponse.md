# ClusterSettingsPluginsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugins** | Pointer to [**[]ClusterPlugin**](ClusterPlugin.md) |  | [optional] 

## Methods

### NewClusterSettingsPluginsResponse

`func NewClusterSettingsPluginsResponse() *ClusterSettingsPluginsResponse`

NewClusterSettingsPluginsResponse instantiates a new ClusterSettingsPluginsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSettingsPluginsResponseWithDefaults

`func NewClusterSettingsPluginsResponseWithDefaults() *ClusterSettingsPluginsResponse`

NewClusterSettingsPluginsResponseWithDefaults instantiates a new ClusterSettingsPluginsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugins

`func (o *ClusterSettingsPluginsResponse) GetPlugins() []ClusterPlugin`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *ClusterSettingsPluginsResponse) GetPluginsOk() (*[]ClusterPlugin, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *ClusterSettingsPluginsResponse) SetPlugins(v []ClusterPlugin)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *ClusterSettingsPluginsResponse) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

