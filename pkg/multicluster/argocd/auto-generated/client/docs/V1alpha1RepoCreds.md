# V1alpha1RepoCreds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableOCI** | Pointer to **bool** |  | [optional] 
**ForceHttpBasicAuth** | Pointer to **bool** |  | [optional] 
**GcpServiceAccountKey** | Pointer to **string** |  | [optional] 
**GithubAppEnterpriseBaseUrl** | Pointer to **string** |  | [optional] 
**GithubAppID** | Pointer to **string** |  | [optional] 
**GithubAppInstallationID** | Pointer to **string** |  | [optional] 
**GithubAppPrivateKey** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Proxy** | Pointer to **string** |  | [optional] 
**SshPrivateKey** | Pointer to **string** |  | [optional] 
**TlsClientCertData** | Pointer to **string** |  | [optional] 
**TlsClientCertKey** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Type specifies the type of the repoCreds. Can be either \&quot;git\&quot; or \&quot;helm. \&quot;git\&quot; is assumed if empty or absent. | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1RepoCreds

`func NewV1alpha1RepoCreds() *V1alpha1RepoCreds`

NewV1alpha1RepoCreds instantiates a new V1alpha1RepoCreds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1RepoCredsWithDefaults

`func NewV1alpha1RepoCredsWithDefaults() *V1alpha1RepoCreds`

NewV1alpha1RepoCredsWithDefaults instantiates a new V1alpha1RepoCreds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableOCI

`func (o *V1alpha1RepoCreds) GetEnableOCI() bool`

GetEnableOCI returns the EnableOCI field if non-nil, zero value otherwise.

### GetEnableOCIOk

`func (o *V1alpha1RepoCreds) GetEnableOCIOk() (*bool, bool)`

GetEnableOCIOk returns a tuple with the EnableOCI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOCI

`func (o *V1alpha1RepoCreds) SetEnableOCI(v bool)`

SetEnableOCI sets EnableOCI field to given value.

### HasEnableOCI

`func (o *V1alpha1RepoCreds) HasEnableOCI() bool`

HasEnableOCI returns a boolean if a field has been set.

### GetForceHttpBasicAuth

`func (o *V1alpha1RepoCreds) GetForceHttpBasicAuth() bool`

GetForceHttpBasicAuth returns the ForceHttpBasicAuth field if non-nil, zero value otherwise.

### GetForceHttpBasicAuthOk

`func (o *V1alpha1RepoCreds) GetForceHttpBasicAuthOk() (*bool, bool)`

GetForceHttpBasicAuthOk returns a tuple with the ForceHttpBasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceHttpBasicAuth

`func (o *V1alpha1RepoCreds) SetForceHttpBasicAuth(v bool)`

SetForceHttpBasicAuth sets ForceHttpBasicAuth field to given value.

### HasForceHttpBasicAuth

`func (o *V1alpha1RepoCreds) HasForceHttpBasicAuth() bool`

HasForceHttpBasicAuth returns a boolean if a field has been set.

### GetGcpServiceAccountKey

`func (o *V1alpha1RepoCreds) GetGcpServiceAccountKey() string`

GetGcpServiceAccountKey returns the GcpServiceAccountKey field if non-nil, zero value otherwise.

### GetGcpServiceAccountKeyOk

`func (o *V1alpha1RepoCreds) GetGcpServiceAccountKeyOk() (*string, bool)`

GetGcpServiceAccountKeyOk returns a tuple with the GcpServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountKey

`func (o *V1alpha1RepoCreds) SetGcpServiceAccountKey(v string)`

SetGcpServiceAccountKey sets GcpServiceAccountKey field to given value.

### HasGcpServiceAccountKey

`func (o *V1alpha1RepoCreds) HasGcpServiceAccountKey() bool`

HasGcpServiceAccountKey returns a boolean if a field has been set.

### GetGithubAppEnterpriseBaseUrl

`func (o *V1alpha1RepoCreds) GetGithubAppEnterpriseBaseUrl() string`

GetGithubAppEnterpriseBaseUrl returns the GithubAppEnterpriseBaseUrl field if non-nil, zero value otherwise.

### GetGithubAppEnterpriseBaseUrlOk

`func (o *V1alpha1RepoCreds) GetGithubAppEnterpriseBaseUrlOk() (*string, bool)`

GetGithubAppEnterpriseBaseUrlOk returns a tuple with the GithubAppEnterpriseBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppEnterpriseBaseUrl

`func (o *V1alpha1RepoCreds) SetGithubAppEnterpriseBaseUrl(v string)`

SetGithubAppEnterpriseBaseUrl sets GithubAppEnterpriseBaseUrl field to given value.

### HasGithubAppEnterpriseBaseUrl

`func (o *V1alpha1RepoCreds) HasGithubAppEnterpriseBaseUrl() bool`

HasGithubAppEnterpriseBaseUrl returns a boolean if a field has been set.

### GetGithubAppID

`func (o *V1alpha1RepoCreds) GetGithubAppID() string`

GetGithubAppID returns the GithubAppID field if non-nil, zero value otherwise.

### GetGithubAppIDOk

`func (o *V1alpha1RepoCreds) GetGithubAppIDOk() (*string, bool)`

GetGithubAppIDOk returns a tuple with the GithubAppID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppID

`func (o *V1alpha1RepoCreds) SetGithubAppID(v string)`

SetGithubAppID sets GithubAppID field to given value.

### HasGithubAppID

`func (o *V1alpha1RepoCreds) HasGithubAppID() bool`

HasGithubAppID returns a boolean if a field has been set.

### GetGithubAppInstallationID

`func (o *V1alpha1RepoCreds) GetGithubAppInstallationID() string`

GetGithubAppInstallationID returns the GithubAppInstallationID field if non-nil, zero value otherwise.

### GetGithubAppInstallationIDOk

`func (o *V1alpha1RepoCreds) GetGithubAppInstallationIDOk() (*string, bool)`

GetGithubAppInstallationIDOk returns a tuple with the GithubAppInstallationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppInstallationID

`func (o *V1alpha1RepoCreds) SetGithubAppInstallationID(v string)`

SetGithubAppInstallationID sets GithubAppInstallationID field to given value.

### HasGithubAppInstallationID

`func (o *V1alpha1RepoCreds) HasGithubAppInstallationID() bool`

HasGithubAppInstallationID returns a boolean if a field has been set.

### GetGithubAppPrivateKey

`func (o *V1alpha1RepoCreds) GetGithubAppPrivateKey() string`

GetGithubAppPrivateKey returns the GithubAppPrivateKey field if non-nil, zero value otherwise.

### GetGithubAppPrivateKeyOk

`func (o *V1alpha1RepoCreds) GetGithubAppPrivateKeyOk() (*string, bool)`

GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppPrivateKey

`func (o *V1alpha1RepoCreds) SetGithubAppPrivateKey(v string)`

SetGithubAppPrivateKey sets GithubAppPrivateKey field to given value.

### HasGithubAppPrivateKey

`func (o *V1alpha1RepoCreds) HasGithubAppPrivateKey() bool`

HasGithubAppPrivateKey returns a boolean if a field has been set.

### GetPassword

`func (o *V1alpha1RepoCreds) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V1alpha1RepoCreds) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V1alpha1RepoCreds) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V1alpha1RepoCreds) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProxy

`func (o *V1alpha1RepoCreds) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *V1alpha1RepoCreds) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *V1alpha1RepoCreds) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *V1alpha1RepoCreds) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetSshPrivateKey

`func (o *V1alpha1RepoCreds) GetSshPrivateKey() string`

GetSshPrivateKey returns the SshPrivateKey field if non-nil, zero value otherwise.

### GetSshPrivateKeyOk

`func (o *V1alpha1RepoCreds) GetSshPrivateKeyOk() (*string, bool)`

GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKey

`func (o *V1alpha1RepoCreds) SetSshPrivateKey(v string)`

SetSshPrivateKey sets SshPrivateKey field to given value.

### HasSshPrivateKey

`func (o *V1alpha1RepoCreds) HasSshPrivateKey() bool`

HasSshPrivateKey returns a boolean if a field has been set.

### GetTlsClientCertData

`func (o *V1alpha1RepoCreds) GetTlsClientCertData() string`

GetTlsClientCertData returns the TlsClientCertData field if non-nil, zero value otherwise.

### GetTlsClientCertDataOk

`func (o *V1alpha1RepoCreds) GetTlsClientCertDataOk() (*string, bool)`

GetTlsClientCertDataOk returns a tuple with the TlsClientCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCertData

`func (o *V1alpha1RepoCreds) SetTlsClientCertData(v string)`

SetTlsClientCertData sets TlsClientCertData field to given value.

### HasTlsClientCertData

`func (o *V1alpha1RepoCreds) HasTlsClientCertData() bool`

HasTlsClientCertData returns a boolean if a field has been set.

### GetTlsClientCertKey

`func (o *V1alpha1RepoCreds) GetTlsClientCertKey() string`

GetTlsClientCertKey returns the TlsClientCertKey field if non-nil, zero value otherwise.

### GetTlsClientCertKeyOk

`func (o *V1alpha1RepoCreds) GetTlsClientCertKeyOk() (*string, bool)`

GetTlsClientCertKeyOk returns a tuple with the TlsClientCertKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCertKey

`func (o *V1alpha1RepoCreds) SetTlsClientCertKey(v string)`

SetTlsClientCertKey sets TlsClientCertKey field to given value.

### HasTlsClientCertKey

`func (o *V1alpha1RepoCreds) HasTlsClientCertKey() bool`

HasTlsClientCertKey returns a boolean if a field has been set.

### GetType

`func (o *V1alpha1RepoCreds) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1alpha1RepoCreds) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1alpha1RepoCreds) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1alpha1RepoCreds) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *V1alpha1RepoCreds) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1alpha1RepoCreds) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1alpha1RepoCreds) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V1alpha1RepoCreds) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *V1alpha1RepoCreds) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *V1alpha1RepoCreds) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *V1alpha1RepoCreds) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *V1alpha1RepoCreds) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


