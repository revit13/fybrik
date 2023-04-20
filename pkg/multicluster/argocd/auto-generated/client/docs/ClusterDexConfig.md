# ClusterDexConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectors** | Pointer to [**[]ClusterConnector**](ClusterConnector.md) |  | [optional] 

## Methods

### NewClusterDexConfig

`func NewClusterDexConfig() *ClusterDexConfig`

NewClusterDexConfig instantiates a new ClusterDexConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDexConfigWithDefaults

`func NewClusterDexConfigWithDefaults() *ClusterDexConfig`

NewClusterDexConfigWithDefaults instantiates a new ClusterDexConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectors

`func (o *ClusterDexConfig) GetConnectors() []ClusterConnector`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *ClusterDexConfig) GetConnectorsOk() (*[]ClusterConnector, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *ClusterDexConfig) SetConnectors(v []ClusterConnector)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *ClusterDexConfig) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


