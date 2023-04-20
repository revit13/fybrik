# V1alpha1AWSAuthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** |  | [optional] 
**RoleARN** | Pointer to **string** | RoleARN contains optional role ARN. If set then AWS IAM Authenticator assume a role to perform cluster operations instead of the default AWS credential provider chain. | [optional] 

## Methods

### NewV1alpha1AWSAuthConfig

`func NewV1alpha1AWSAuthConfig() *V1alpha1AWSAuthConfig`

NewV1alpha1AWSAuthConfig instantiates a new V1alpha1AWSAuthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1AWSAuthConfigWithDefaults

`func NewV1alpha1AWSAuthConfigWithDefaults() *V1alpha1AWSAuthConfig`

NewV1alpha1AWSAuthConfigWithDefaults instantiates a new V1alpha1AWSAuthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *V1alpha1AWSAuthConfig) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *V1alpha1AWSAuthConfig) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *V1alpha1AWSAuthConfig) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *V1alpha1AWSAuthConfig) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetRoleARN

`func (o *V1alpha1AWSAuthConfig) GetRoleARN() string`

GetRoleARN returns the RoleARN field if non-nil, zero value otherwise.

### GetRoleARNOk

`func (o *V1alpha1AWSAuthConfig) GetRoleARNOk() (*string, bool)`

GetRoleARNOk returns a tuple with the RoleARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleARN

`func (o *V1alpha1AWSAuthConfig) SetRoleARN(v string)`

SetRoleARN sets RoleARN field to given value.

### HasRoleARN

`func (o *V1alpha1AWSAuthConfig) HasRoleARN() bool`

HasRoleARN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


