# V1alpha1TLSClientConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaData** | Pointer to **string** |  | [optional] 
**CertData** | Pointer to **string** |  | [optional] 
**Insecure** | Pointer to **bool** | Insecure specifies that the server should be accessed without verifying the TLS certificate. For testing only. | [optional] 
**KeyData** | Pointer to **string** |  | [optional] 
**ServerName** | Pointer to **string** | ServerName is passed to the server for SNI and is used in the client to check server certificates against. If ServerName is empty, the hostname used to contact the server is used. | [optional] 

## Methods

### NewV1alpha1TLSClientConfig

`func NewV1alpha1TLSClientConfig() *V1alpha1TLSClientConfig`

NewV1alpha1TLSClientConfig instantiates a new V1alpha1TLSClientConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1TLSClientConfigWithDefaults

`func NewV1alpha1TLSClientConfigWithDefaults() *V1alpha1TLSClientConfig`

NewV1alpha1TLSClientConfigWithDefaults instantiates a new V1alpha1TLSClientConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaData

`func (o *V1alpha1TLSClientConfig) GetCaData() string`

GetCaData returns the CaData field if non-nil, zero value otherwise.

### GetCaDataOk

`func (o *V1alpha1TLSClientConfig) GetCaDataOk() (*string, bool)`

GetCaDataOk returns a tuple with the CaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaData

`func (o *V1alpha1TLSClientConfig) SetCaData(v string)`

SetCaData sets CaData field to given value.

### HasCaData

`func (o *V1alpha1TLSClientConfig) HasCaData() bool`

HasCaData returns a boolean if a field has been set.

### GetCertData

`func (o *V1alpha1TLSClientConfig) GetCertData() string`

GetCertData returns the CertData field if non-nil, zero value otherwise.

### GetCertDataOk

`func (o *V1alpha1TLSClientConfig) GetCertDataOk() (*string, bool)`

GetCertDataOk returns a tuple with the CertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertData

`func (o *V1alpha1TLSClientConfig) SetCertData(v string)`

SetCertData sets CertData field to given value.

### HasCertData

`func (o *V1alpha1TLSClientConfig) HasCertData() bool`

HasCertData returns a boolean if a field has been set.

### GetInsecure

`func (o *V1alpha1TLSClientConfig) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *V1alpha1TLSClientConfig) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *V1alpha1TLSClientConfig) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *V1alpha1TLSClientConfig) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetKeyData

`func (o *V1alpha1TLSClientConfig) GetKeyData() string`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *V1alpha1TLSClientConfig) GetKeyDataOk() (*string, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *V1alpha1TLSClientConfig) SetKeyData(v string)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *V1alpha1TLSClientConfig) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.

### GetServerName

`func (o *V1alpha1TLSClientConfig) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *V1alpha1TLSClientConfig) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *V1alpha1TLSClientConfig) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *V1alpha1TLSClientConfig) HasServerName() bool`

HasServerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


