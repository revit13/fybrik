# V1PortStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 

## Methods

### NewV1PortStatus

`func NewV1PortStatus() *V1PortStatus`

NewV1PortStatus instantiates a new V1PortStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PortStatusWithDefaults

`func NewV1PortStatusWithDefaults() *V1PortStatus`

NewV1PortStatusWithDefaults instantiates a new V1PortStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *V1PortStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V1PortStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V1PortStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V1PortStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPort

`func (o *V1PortStatus) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V1PortStatus) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V1PortStatus) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *V1PortStatus) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *V1PortStatus) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *V1PortStatus) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *V1PortStatus) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *V1PortStatus) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


