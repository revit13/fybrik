/*
Consolidate Services

Description of all APIs

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
)

// ClusterOIDCConfig struct for ClusterOIDCConfig
type ClusterOIDCConfig struct {
	CliClientID *string `json:"cliClientID,omitempty"`
	ClientID *string `json:"clientID,omitempty"`
	IdTokenClaims *map[string]OidcClaim `json:"idTokenClaims,omitempty"`
	Issuer *string `json:"issuer,omitempty"`
	Name *string `json:"name,omitempty"`
	Scopes *[]string `json:"scopes,omitempty"`
}

// NewClusterOIDCConfig instantiates a new ClusterOIDCConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterOIDCConfig() *ClusterOIDCConfig {
	this := ClusterOIDCConfig{}
	return &this
}

// NewClusterOIDCConfigWithDefaults instantiates a new ClusterOIDCConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterOIDCConfigWithDefaults() *ClusterOIDCConfig {
	this := ClusterOIDCConfig{}
	return &this
}

// GetCliClientID returns the CliClientID field value if set, zero value otherwise.
func (o *ClusterOIDCConfig) GetCliClientID() string {
	if o == nil || o.CliClientID == nil {
		var ret string
		return ret
	}
	return *o.CliClientID
}

// GetCliClientIDOk returns a tuple with the CliClientID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOIDCConfig) GetCliClientIDOk() (*string, bool) {
	if o == nil || o.CliClientID == nil {
		return nil, false
	}
	return o.CliClientID, true
}

// HasCliClientID returns a boolean if a field has been set.
func (o *ClusterOIDCConfig) HasCliClientID() bool {
	if o != nil && o.CliClientID != nil {
		return true
	}

	return false
}

// SetCliClientID gets a reference to the given string and assigns it to the CliClientID field.
func (o *ClusterOIDCConfig) SetCliClientID(v string) {
	o.CliClientID = &v
}

// GetClientID returns the ClientID field value if set, zero value otherwise.
func (o *ClusterOIDCConfig) GetClientID() string {
	if o == nil || o.ClientID == nil {
		var ret string
		return ret
	}
	return *o.ClientID
}

// GetClientIDOk returns a tuple with the ClientID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOIDCConfig) GetClientIDOk() (*string, bool) {
	if o == nil || o.ClientID == nil {
		return nil, false
	}
	return o.ClientID, true
}

// HasClientID returns a boolean if a field has been set.
func (o *ClusterOIDCConfig) HasClientID() bool {
	if o != nil && o.ClientID != nil {
		return true
	}

	return false
}

// SetClientID gets a reference to the given string and assigns it to the ClientID field.
func (o *ClusterOIDCConfig) SetClientID(v string) {
	o.ClientID = &v
}

// GetIdTokenClaims returns the IdTokenClaims field value if set, zero value otherwise.
func (o *ClusterOIDCConfig) GetIdTokenClaims() map[string]OidcClaim {
	if o == nil || o.IdTokenClaims == nil {
		var ret map[string]OidcClaim
		return ret
	}
	return *o.IdTokenClaims
}

// GetIdTokenClaimsOk returns a tuple with the IdTokenClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOIDCConfig) GetIdTokenClaimsOk() (*map[string]OidcClaim, bool) {
	if o == nil || o.IdTokenClaims == nil {
		return nil, false
	}
	return o.IdTokenClaims, true
}

// HasIdTokenClaims returns a boolean if a field has been set.
func (o *ClusterOIDCConfig) HasIdTokenClaims() bool {
	if o != nil && o.IdTokenClaims != nil {
		return true
	}

	return false
}

// SetIdTokenClaims gets a reference to the given map[string]OidcClaim and assigns it to the IdTokenClaims field.
func (o *ClusterOIDCConfig) SetIdTokenClaims(v map[string]OidcClaim) {
	o.IdTokenClaims = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *ClusterOIDCConfig) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOIDCConfig) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *ClusterOIDCConfig) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *ClusterOIDCConfig) SetIssuer(v string) {
	o.Issuer = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterOIDCConfig) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOIDCConfig) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterOIDCConfig) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterOIDCConfig) SetName(v string) {
	o.Name = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ClusterOIDCConfig) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOIDCConfig) GetScopesOk() (*[]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ClusterOIDCConfig) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ClusterOIDCConfig) SetScopes(v []string) {
	o.Scopes = &v
}

func (o ClusterOIDCConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CliClientID != nil {
		toSerialize["cliClientID"] = o.CliClientID
	}
	if o.ClientID != nil {
		toSerialize["clientID"] = o.ClientID
	}
	if o.IdTokenClaims != nil {
		toSerialize["idTokenClaims"] = o.IdTokenClaims
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	return json.Marshal(toSerialize)
}

type NullableClusterOIDCConfig struct {
	value *ClusterOIDCConfig
	isSet bool
}

func (v NullableClusterOIDCConfig) Get() *ClusterOIDCConfig {
	return v.value
}

func (v *NullableClusterOIDCConfig) Set(val *ClusterOIDCConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterOIDCConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterOIDCConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterOIDCConfig(val *ClusterOIDCConfig) *NullableClusterOIDCConfig {
	return &NullableClusterOIDCConfig{value: val, isSet: true}
}

func (v NullableClusterOIDCConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterOIDCConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


