# OidcClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Essential** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOidcClaim

`func NewOidcClaim() *OidcClaim`

NewOidcClaim instantiates a new OidcClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcClaimWithDefaults

`func NewOidcClaimWithDefaults() *OidcClaim`

NewOidcClaimWithDefaults instantiates a new OidcClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEssential

`func (o *OidcClaim) GetEssential() bool`

GetEssential returns the Essential field if non-nil, zero value otherwise.

### GetEssentialOk

`func (o *OidcClaim) GetEssentialOk() (*bool, bool)`

GetEssentialOk returns a tuple with the Essential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEssential

`func (o *OidcClaim) SetEssential(v bool)`

SetEssential sets Essential field to given value.

### HasEssential

`func (o *OidcClaim) HasEssential() bool`

HasEssential returns a boolean if a field has been set.

### GetValue

`func (o *OidcClaim) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OidcClaim) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OidcClaim) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *OidcClaim) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValues

`func (o *OidcClaim) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *OidcClaim) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *OidcClaim) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *OidcClaim) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


