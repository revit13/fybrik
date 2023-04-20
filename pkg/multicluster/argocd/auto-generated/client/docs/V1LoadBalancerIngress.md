# V1LoadBalancerIngress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to [**[]V1PortStatus**](V1PortStatus.md) |  | [optional] 

## Methods

### NewV1LoadBalancerIngress

`func NewV1LoadBalancerIngress() *V1LoadBalancerIngress`

NewV1LoadBalancerIngress instantiates a new V1LoadBalancerIngress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LoadBalancerIngressWithDefaults

`func NewV1LoadBalancerIngressWithDefaults() *V1LoadBalancerIngress`

NewV1LoadBalancerIngressWithDefaults instantiates a new V1LoadBalancerIngress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *V1LoadBalancerIngress) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V1LoadBalancerIngress) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V1LoadBalancerIngress) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V1LoadBalancerIngress) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *V1LoadBalancerIngress) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V1LoadBalancerIngress) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V1LoadBalancerIngress) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V1LoadBalancerIngress) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPorts

`func (o *V1LoadBalancerIngress) GetPorts() []V1PortStatus`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *V1LoadBalancerIngress) GetPortsOk() (*[]V1PortStatus, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *V1LoadBalancerIngress) SetPorts(v []V1PortStatus)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *V1LoadBalancerIngress) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


