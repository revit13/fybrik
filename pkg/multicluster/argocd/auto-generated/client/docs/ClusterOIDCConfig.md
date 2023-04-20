# ClusterOIDCConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliClientID** | Pointer to **string** |  | [optional] 
**ClientID** | Pointer to **string** |  | [optional] 
**IdTokenClaims** | Pointer to [**map[string]OidcClaim**](OidcClaim.md) |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewClusterOIDCConfig

`func NewClusterOIDCConfig() *ClusterOIDCConfig`

NewClusterOIDCConfig instantiates a new ClusterOIDCConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterOIDCConfigWithDefaults

`func NewClusterOIDCConfigWithDefaults() *ClusterOIDCConfig`

NewClusterOIDCConfigWithDefaults instantiates a new ClusterOIDCConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliClientID

`func (o *ClusterOIDCConfig) GetCliClientID() string`

GetCliClientID returns the CliClientID field if non-nil, zero value otherwise.

### GetCliClientIDOk

`func (o *ClusterOIDCConfig) GetCliClientIDOk() (*string, bool)`

GetCliClientIDOk returns a tuple with the CliClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliClientID

`func (o *ClusterOIDCConfig) SetCliClientID(v string)`

SetCliClientID sets CliClientID field to given value.

### HasCliClientID

`func (o *ClusterOIDCConfig) HasCliClientID() bool`

HasCliClientID returns a boolean if a field has been set.

### GetClientID

`func (o *ClusterOIDCConfig) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *ClusterOIDCConfig) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *ClusterOIDCConfig) SetClientID(v string)`

SetClientID sets ClientID field to given value.

### HasClientID

`func (o *ClusterOIDCConfig) HasClientID() bool`

HasClientID returns a boolean if a field has been set.

### GetIdTokenClaims

`func (o *ClusterOIDCConfig) GetIdTokenClaims() map[string]OidcClaim`

GetIdTokenClaims returns the IdTokenClaims field if non-nil, zero value otherwise.

### GetIdTokenClaimsOk

`func (o *ClusterOIDCConfig) GetIdTokenClaimsOk() (*map[string]OidcClaim, bool)`

GetIdTokenClaimsOk returns a tuple with the IdTokenClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenClaims

`func (o *ClusterOIDCConfig) SetIdTokenClaims(v map[string]OidcClaim)`

SetIdTokenClaims sets IdTokenClaims field to given value.

### HasIdTokenClaims

`func (o *ClusterOIDCConfig) HasIdTokenClaims() bool`

HasIdTokenClaims returns a boolean if a field has been set.

### GetIssuer

`func (o *ClusterOIDCConfig) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ClusterOIDCConfig) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ClusterOIDCConfig) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ClusterOIDCConfig) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetName

`func (o *ClusterOIDCConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterOIDCConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterOIDCConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterOIDCConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopes

`func (o *ClusterOIDCConfig) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ClusterOIDCConfig) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ClusterOIDCConfig) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ClusterOIDCConfig) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


