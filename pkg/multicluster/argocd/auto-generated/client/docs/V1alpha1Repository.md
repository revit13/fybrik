# V1alpha1Repository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionState** | Pointer to [**V1alpha1ConnectionState**](V1alpha1ConnectionState.md) |  | [optional] 
**EnableLfs** | Pointer to **bool** | EnableLFS specifies whether git-lfs support should be enabled for this repo. Only valid for Git repositories. | [optional] 
**EnableOCI** | Pointer to **bool** |  | [optional] 
**ForceHttpBasicAuth** | Pointer to **bool** |  | [optional] 
**GcpServiceAccountKey** | Pointer to **string** |  | [optional] 
**GithubAppEnterpriseBaseUrl** | Pointer to **string** |  | [optional] 
**GithubAppID** | Pointer to **string** |  | [optional] 
**GithubAppInstallationID** | Pointer to **string** |  | [optional] 
**GithubAppPrivateKey** | Pointer to **string** |  | [optional] 
**InheritedCreds** | Pointer to **bool** |  | [optional] 
**Insecure** | Pointer to **bool** |  | [optional] 
**InsecureIgnoreHostKey** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Proxy** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**SshPrivateKey** | Pointer to **string** | SSHPrivateKey contains the PEM data for authenticating at the repo server. Only used with Git repos. | [optional] 
**TlsClientCertData** | Pointer to **string** |  | [optional] 
**TlsClientCertKey** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Type specifies the type of the repo. Can be either \&quot;git\&quot; or \&quot;helm. \&quot;git\&quot; is assumed if empty or absent. | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1Repository

`func NewV1alpha1Repository() *V1alpha1Repository`

NewV1alpha1Repository instantiates a new V1alpha1Repository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1RepositoryWithDefaults

`func NewV1alpha1RepositoryWithDefaults() *V1alpha1Repository`

NewV1alpha1RepositoryWithDefaults instantiates a new V1alpha1Repository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionState

`func (o *V1alpha1Repository) GetConnectionState() V1alpha1ConnectionState`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *V1alpha1Repository) GetConnectionStateOk() (*V1alpha1ConnectionState, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *V1alpha1Repository) SetConnectionState(v V1alpha1ConnectionState)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *V1alpha1Repository) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetEnableLfs

`func (o *V1alpha1Repository) GetEnableLfs() bool`

GetEnableLfs returns the EnableLfs field if non-nil, zero value otherwise.

### GetEnableLfsOk

`func (o *V1alpha1Repository) GetEnableLfsOk() (*bool, bool)`

GetEnableLfsOk returns a tuple with the EnableLfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLfs

`func (o *V1alpha1Repository) SetEnableLfs(v bool)`

SetEnableLfs sets EnableLfs field to given value.

### HasEnableLfs

`func (o *V1alpha1Repository) HasEnableLfs() bool`

HasEnableLfs returns a boolean if a field has been set.

### GetEnableOCI

`func (o *V1alpha1Repository) GetEnableOCI() bool`

GetEnableOCI returns the EnableOCI field if non-nil, zero value otherwise.

### GetEnableOCIOk

`func (o *V1alpha1Repository) GetEnableOCIOk() (*bool, bool)`

GetEnableOCIOk returns a tuple with the EnableOCI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOCI

`func (o *V1alpha1Repository) SetEnableOCI(v bool)`

SetEnableOCI sets EnableOCI field to given value.

### HasEnableOCI

`func (o *V1alpha1Repository) HasEnableOCI() bool`

HasEnableOCI returns a boolean if a field has been set.

### GetForceHttpBasicAuth

`func (o *V1alpha1Repository) GetForceHttpBasicAuth() bool`

GetForceHttpBasicAuth returns the ForceHttpBasicAuth field if non-nil, zero value otherwise.

### GetForceHttpBasicAuthOk

`func (o *V1alpha1Repository) GetForceHttpBasicAuthOk() (*bool, bool)`

GetForceHttpBasicAuthOk returns a tuple with the ForceHttpBasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceHttpBasicAuth

`func (o *V1alpha1Repository) SetForceHttpBasicAuth(v bool)`

SetForceHttpBasicAuth sets ForceHttpBasicAuth field to given value.

### HasForceHttpBasicAuth

`func (o *V1alpha1Repository) HasForceHttpBasicAuth() bool`

HasForceHttpBasicAuth returns a boolean if a field has been set.

### GetGcpServiceAccountKey

`func (o *V1alpha1Repository) GetGcpServiceAccountKey() string`

GetGcpServiceAccountKey returns the GcpServiceAccountKey field if non-nil, zero value otherwise.

### GetGcpServiceAccountKeyOk

`func (o *V1alpha1Repository) GetGcpServiceAccountKeyOk() (*string, bool)`

GetGcpServiceAccountKeyOk returns a tuple with the GcpServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountKey

`func (o *V1alpha1Repository) SetGcpServiceAccountKey(v string)`

SetGcpServiceAccountKey sets GcpServiceAccountKey field to given value.

### HasGcpServiceAccountKey

`func (o *V1alpha1Repository) HasGcpServiceAccountKey() bool`

HasGcpServiceAccountKey returns a boolean if a field has been set.

### GetGithubAppEnterpriseBaseUrl

`func (o *V1alpha1Repository) GetGithubAppEnterpriseBaseUrl() string`

GetGithubAppEnterpriseBaseUrl returns the GithubAppEnterpriseBaseUrl field if non-nil, zero value otherwise.

### GetGithubAppEnterpriseBaseUrlOk

`func (o *V1alpha1Repository) GetGithubAppEnterpriseBaseUrlOk() (*string, bool)`

GetGithubAppEnterpriseBaseUrlOk returns a tuple with the GithubAppEnterpriseBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppEnterpriseBaseUrl

`func (o *V1alpha1Repository) SetGithubAppEnterpriseBaseUrl(v string)`

SetGithubAppEnterpriseBaseUrl sets GithubAppEnterpriseBaseUrl field to given value.

### HasGithubAppEnterpriseBaseUrl

`func (o *V1alpha1Repository) HasGithubAppEnterpriseBaseUrl() bool`

HasGithubAppEnterpriseBaseUrl returns a boolean if a field has been set.

### GetGithubAppID

`func (o *V1alpha1Repository) GetGithubAppID() string`

GetGithubAppID returns the GithubAppID field if non-nil, zero value otherwise.

### GetGithubAppIDOk

`func (o *V1alpha1Repository) GetGithubAppIDOk() (*string, bool)`

GetGithubAppIDOk returns a tuple with the GithubAppID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppID

`func (o *V1alpha1Repository) SetGithubAppID(v string)`

SetGithubAppID sets GithubAppID field to given value.

### HasGithubAppID

`func (o *V1alpha1Repository) HasGithubAppID() bool`

HasGithubAppID returns a boolean if a field has been set.

### GetGithubAppInstallationID

`func (o *V1alpha1Repository) GetGithubAppInstallationID() string`

GetGithubAppInstallationID returns the GithubAppInstallationID field if non-nil, zero value otherwise.

### GetGithubAppInstallationIDOk

`func (o *V1alpha1Repository) GetGithubAppInstallationIDOk() (*string, bool)`

GetGithubAppInstallationIDOk returns a tuple with the GithubAppInstallationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppInstallationID

`func (o *V1alpha1Repository) SetGithubAppInstallationID(v string)`

SetGithubAppInstallationID sets GithubAppInstallationID field to given value.

### HasGithubAppInstallationID

`func (o *V1alpha1Repository) HasGithubAppInstallationID() bool`

HasGithubAppInstallationID returns a boolean if a field has been set.

### GetGithubAppPrivateKey

`func (o *V1alpha1Repository) GetGithubAppPrivateKey() string`

GetGithubAppPrivateKey returns the GithubAppPrivateKey field if non-nil, zero value otherwise.

### GetGithubAppPrivateKeyOk

`func (o *V1alpha1Repository) GetGithubAppPrivateKeyOk() (*string, bool)`

GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppPrivateKey

`func (o *V1alpha1Repository) SetGithubAppPrivateKey(v string)`

SetGithubAppPrivateKey sets GithubAppPrivateKey field to given value.

### HasGithubAppPrivateKey

`func (o *V1alpha1Repository) HasGithubAppPrivateKey() bool`

HasGithubAppPrivateKey returns a boolean if a field has been set.

### GetInheritedCreds

`func (o *V1alpha1Repository) GetInheritedCreds() bool`

GetInheritedCreds returns the InheritedCreds field if non-nil, zero value otherwise.

### GetInheritedCredsOk

`func (o *V1alpha1Repository) GetInheritedCredsOk() (*bool, bool)`

GetInheritedCredsOk returns a tuple with the InheritedCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedCreds

`func (o *V1alpha1Repository) SetInheritedCreds(v bool)`

SetInheritedCreds sets InheritedCreds field to given value.

### HasInheritedCreds

`func (o *V1alpha1Repository) HasInheritedCreds() bool`

HasInheritedCreds returns a boolean if a field has been set.

### GetInsecure

`func (o *V1alpha1Repository) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *V1alpha1Repository) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *V1alpha1Repository) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *V1alpha1Repository) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetInsecureIgnoreHostKey

`func (o *V1alpha1Repository) GetInsecureIgnoreHostKey() bool`

GetInsecureIgnoreHostKey returns the InsecureIgnoreHostKey field if non-nil, zero value otherwise.

### GetInsecureIgnoreHostKeyOk

`func (o *V1alpha1Repository) GetInsecureIgnoreHostKeyOk() (*bool, bool)`

GetInsecureIgnoreHostKeyOk returns a tuple with the InsecureIgnoreHostKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureIgnoreHostKey

`func (o *V1alpha1Repository) SetInsecureIgnoreHostKey(v bool)`

SetInsecureIgnoreHostKey sets InsecureIgnoreHostKey field to given value.

### HasInsecureIgnoreHostKey

`func (o *V1alpha1Repository) HasInsecureIgnoreHostKey() bool`

HasInsecureIgnoreHostKey returns a boolean if a field has been set.

### GetName

`func (o *V1alpha1Repository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha1Repository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha1Repository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1alpha1Repository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *V1alpha1Repository) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V1alpha1Repository) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V1alpha1Repository) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V1alpha1Repository) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProject

`func (o *V1alpha1Repository) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *V1alpha1Repository) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *V1alpha1Repository) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *V1alpha1Repository) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProxy

`func (o *V1alpha1Repository) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *V1alpha1Repository) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *V1alpha1Repository) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *V1alpha1Repository) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetRepo

`func (o *V1alpha1Repository) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *V1alpha1Repository) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *V1alpha1Repository) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *V1alpha1Repository) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSshPrivateKey

`func (o *V1alpha1Repository) GetSshPrivateKey() string`

GetSshPrivateKey returns the SshPrivateKey field if non-nil, zero value otherwise.

### GetSshPrivateKeyOk

`func (o *V1alpha1Repository) GetSshPrivateKeyOk() (*string, bool)`

GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKey

`func (o *V1alpha1Repository) SetSshPrivateKey(v string)`

SetSshPrivateKey sets SshPrivateKey field to given value.

### HasSshPrivateKey

`func (o *V1alpha1Repository) HasSshPrivateKey() bool`

HasSshPrivateKey returns a boolean if a field has been set.

### GetTlsClientCertData

`func (o *V1alpha1Repository) GetTlsClientCertData() string`

GetTlsClientCertData returns the TlsClientCertData field if non-nil, zero value otherwise.

### GetTlsClientCertDataOk

`func (o *V1alpha1Repository) GetTlsClientCertDataOk() (*string, bool)`

GetTlsClientCertDataOk returns a tuple with the TlsClientCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCertData

`func (o *V1alpha1Repository) SetTlsClientCertData(v string)`

SetTlsClientCertData sets TlsClientCertData field to given value.

### HasTlsClientCertData

`func (o *V1alpha1Repository) HasTlsClientCertData() bool`

HasTlsClientCertData returns a boolean if a field has been set.

### GetTlsClientCertKey

`func (o *V1alpha1Repository) GetTlsClientCertKey() string`

GetTlsClientCertKey returns the TlsClientCertKey field if non-nil, zero value otherwise.

### GetTlsClientCertKeyOk

`func (o *V1alpha1Repository) GetTlsClientCertKeyOk() (*string, bool)`

GetTlsClientCertKeyOk returns a tuple with the TlsClientCertKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCertKey

`func (o *V1alpha1Repository) SetTlsClientCertKey(v string)`

SetTlsClientCertKey sets TlsClientCertKey field to given value.

### HasTlsClientCertKey

`func (o *V1alpha1Repository) HasTlsClientCertKey() bool`

HasTlsClientCertKey returns a boolean if a field has been set.

### GetType

`func (o *V1alpha1Repository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1alpha1Repository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1alpha1Repository) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1alpha1Repository) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *V1alpha1Repository) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *V1alpha1Repository) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *V1alpha1Repository) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *V1alpha1Repository) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


