# V1alpha1BasicAuthBitbucketServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordRef** | Pointer to [**V1alpha1SecretRef**](V1alpha1SecretRef.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1BasicAuthBitbucketServer

`func NewV1alpha1BasicAuthBitbucketServer() *V1alpha1BasicAuthBitbucketServer`

NewV1alpha1BasicAuthBitbucketServer instantiates a new V1alpha1BasicAuthBitbucketServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1BasicAuthBitbucketServerWithDefaults

`func NewV1alpha1BasicAuthBitbucketServerWithDefaults() *V1alpha1BasicAuthBitbucketServer`

NewV1alpha1BasicAuthBitbucketServerWithDefaults instantiates a new V1alpha1BasicAuthBitbucketServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasswordRef

`func (o *V1alpha1BasicAuthBitbucketServer) GetPasswordRef() V1alpha1SecretRef`

GetPasswordRef returns the PasswordRef field if non-nil, zero value otherwise.

### GetPasswordRefOk

`func (o *V1alpha1BasicAuthBitbucketServer) GetPasswordRefOk() (*V1alpha1SecretRef, bool)`

GetPasswordRefOk returns a tuple with the PasswordRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRef

`func (o *V1alpha1BasicAuthBitbucketServer) SetPasswordRef(v V1alpha1SecretRef)`

SetPasswordRef sets PasswordRef field to given value.

### HasPasswordRef

`func (o *V1alpha1BasicAuthBitbucketServer) HasPasswordRef() bool`

HasPasswordRef returns a boolean if a field has been set.

### GetUsername

`func (o *V1alpha1BasicAuthBitbucketServer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *V1alpha1BasicAuthBitbucketServer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *V1alpha1BasicAuthBitbucketServer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *V1alpha1BasicAuthBitbucketServer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


