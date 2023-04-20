# V1alpha1SyncPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automated** | Pointer to [**V1alpha1SyncPolicyAutomated**](V1alpha1SyncPolicyAutomated.md) |  | [optional] 
**ManagedNamespaceMetadata** | Pointer to [**V1alpha1ManagedNamespaceMetadata**](V1alpha1ManagedNamespaceMetadata.md) |  | [optional] 
**Retry** | Pointer to [**V1alpha1RetryStrategy**](V1alpha1RetryStrategy.md) |  | [optional] 
**SyncOptions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1alpha1SyncPolicy

`func NewV1alpha1SyncPolicy() *V1alpha1SyncPolicy`

NewV1alpha1SyncPolicy instantiates a new V1alpha1SyncPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1SyncPolicyWithDefaults

`func NewV1alpha1SyncPolicyWithDefaults() *V1alpha1SyncPolicy`

NewV1alpha1SyncPolicyWithDefaults instantiates a new V1alpha1SyncPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomated

`func (o *V1alpha1SyncPolicy) GetAutomated() V1alpha1SyncPolicyAutomated`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *V1alpha1SyncPolicy) GetAutomatedOk() (*V1alpha1SyncPolicyAutomated, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *V1alpha1SyncPolicy) SetAutomated(v V1alpha1SyncPolicyAutomated)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *V1alpha1SyncPolicy) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetManagedNamespaceMetadata

`func (o *V1alpha1SyncPolicy) GetManagedNamespaceMetadata() V1alpha1ManagedNamespaceMetadata`

GetManagedNamespaceMetadata returns the ManagedNamespaceMetadata field if non-nil, zero value otherwise.

### GetManagedNamespaceMetadataOk

`func (o *V1alpha1SyncPolicy) GetManagedNamespaceMetadataOk() (*V1alpha1ManagedNamespaceMetadata, bool)`

GetManagedNamespaceMetadataOk returns a tuple with the ManagedNamespaceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNamespaceMetadata

`func (o *V1alpha1SyncPolicy) SetManagedNamespaceMetadata(v V1alpha1ManagedNamespaceMetadata)`

SetManagedNamespaceMetadata sets ManagedNamespaceMetadata field to given value.

### HasManagedNamespaceMetadata

`func (o *V1alpha1SyncPolicy) HasManagedNamespaceMetadata() bool`

HasManagedNamespaceMetadata returns a boolean if a field has been set.

### GetRetry

`func (o *V1alpha1SyncPolicy) GetRetry() V1alpha1RetryStrategy`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *V1alpha1SyncPolicy) GetRetryOk() (*V1alpha1RetryStrategy, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *V1alpha1SyncPolicy) SetRetry(v V1alpha1RetryStrategy)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *V1alpha1SyncPolicy) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetSyncOptions

`func (o *V1alpha1SyncPolicy) GetSyncOptions() []string`

GetSyncOptions returns the SyncOptions field if non-nil, zero value otherwise.

### GetSyncOptionsOk

`func (o *V1alpha1SyncPolicy) GetSyncOptionsOk() (*[]string, bool)`

GetSyncOptionsOk returns a tuple with the SyncOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncOptions

`func (o *V1alpha1SyncPolicy) SetSyncOptions(v []string)`

SetSyncOptions sets SyncOptions field to given value.

### HasSyncOptions

`func (o *V1alpha1SyncPolicy) HasSyncOptions() bool`

HasSyncOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


