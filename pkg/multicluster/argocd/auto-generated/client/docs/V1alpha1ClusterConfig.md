# V1alpha1ClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAuthConfig** | Pointer to [**V1alpha1AWSAuthConfig**](V1alpha1AWSAuthConfig.md) |  | [optional] 
**BearerToken** | Pointer to **string** | Server requires Bearer authentication. This client will not attempt to use refresh tokens for an OAuth2 flow. TODO: demonstrate an OAuth2 compatible client. | [optional] 
**ExecProviderConfig** | Pointer to [**V1alpha1ExecProviderConfig**](V1alpha1ExecProviderConfig.md) |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**TlsClientConfig** | Pointer to [**V1alpha1TLSClientConfig**](V1alpha1TLSClientConfig.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ClusterConfig

`func NewV1alpha1ClusterConfig() *V1alpha1ClusterConfig`

NewV1alpha1ClusterConfig instantiates a new V1alpha1ClusterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ClusterConfigWithDefaults

`func NewV1alpha1ClusterConfigWithDefaults() *V1alpha1ClusterConfig`

NewV1alpha1ClusterConfigWithDefaults instantiates a new V1alpha1ClusterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAuthConfig

`func (o *V1alpha1ClusterConfig) GetAwsAuthConfig() V1alpha1AWSAuthConfig`

GetAwsAuthConfig returns the AwsAuthConfig field if non-nil, zero value otherwise.

### GetAwsAuthConfigOk

`func (o *V1alpha1ClusterConfig) GetAwsAuthConfigOk() (*V1alpha1AWSAuthConfig, bool)`

GetAwsAuthConfigOk returns a tuple with the AwsAuthConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAuthConfig

`func (o *V1alpha1ClusterConfig) SetAwsAuthConfig(v V1alpha1AWSAuthConfig)`

SetAwsAuthConfig sets AwsAuthConfig field to given value.

### HasAwsAuthConfig

`func (o *V1alpha1ClusterConfig) HasAwsAuthConfig() bool`

HasAwsAuthConfig returns a boolean if a field has been set.

### GetBearerToken

`func (o *V1alpha1ClusterConfig) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *V1alpha1ClusterConfig) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *V1alpha1ClusterConfig) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *V1alpha1ClusterConfig) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetExecProviderConfig

`func (o *V1alpha1ClusterConfig) GetExecProviderConfig() V1alpha1ExecProviderConfig`

GetExecProviderConfig returns the ExecProviderConfig field if non-nil, zero value otherwise.

### GetExecProviderConfigOk

`func (o *V1alpha1ClusterConfig) GetExecProviderConfigOk() (*V1alpha1ExecProviderConfig, bool)`

GetExecProviderConfigOk returns a tuple with the ExecProviderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecProviderConfig

`func (o *V1alpha1ClusterConfig) SetExecProviderConfig(v V1alpha1ExecProviderConfig)`

SetExecProviderConfig sets ExecProviderConfig field to given value.

### HasExecProviderConfig

`func (o *V1alpha1ClusterConfig) HasExecProviderConfig() bool`

HasExecProviderConfig returns a boolean if a field has been set.

### GetPassword

`func (o *V1alpha1ClusterConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V1alpha1ClusterConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V1alpha1ClusterConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V1alpha1ClusterConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTlsClientConfig

`func (o *V1alpha1ClusterConfig) GetTlsClientConfig() V1alpha1TLSClientConfig`

GetTlsClientConfig returns the TlsClientConfig field if non-nil, zero value otherwise.

### GetTlsClientConfigOk

`func (o *V1alpha1ClusterConfig) GetTlsClientConfigOk() (*V1alpha1TLSClientConfig, bool)`

GetTlsClientConfigOk returns a tuple with the TlsClientConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientConfig

`func (o *V1alpha1ClusterConfig) SetTlsClientConfig(v V1alpha1TLSClientConfig)`

SetTlsClientConfig sets TlsClientConfig field to given value.

### HasTlsClientConfig

`func (o *V1alpha1ClusterConfig) HasTlsClientConfig() bool`

HasTlsClientConfig returns a boolean if a field has been set.

### GetUsername

`func (o *V1alpha1ClusterConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *V1alpha1ClusterConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *V1alpha1ClusterConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *V1alpha1ClusterConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


