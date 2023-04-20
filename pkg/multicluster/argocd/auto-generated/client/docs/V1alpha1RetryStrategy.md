# V1alpha1RetryStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backoff** | Pointer to [**V1alpha1Backoff**](V1alpha1Backoff.md) |  | [optional] 
**Limit** | Pointer to **string** | Limit is the maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed. | [optional] 

## Methods

### NewV1alpha1RetryStrategy

`func NewV1alpha1RetryStrategy() *V1alpha1RetryStrategy`

NewV1alpha1RetryStrategy instantiates a new V1alpha1RetryStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1RetryStrategyWithDefaults

`func NewV1alpha1RetryStrategyWithDefaults() *V1alpha1RetryStrategy`

NewV1alpha1RetryStrategyWithDefaults instantiates a new V1alpha1RetryStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackoff

`func (o *V1alpha1RetryStrategy) GetBackoff() V1alpha1Backoff`

GetBackoff returns the Backoff field if non-nil, zero value otherwise.

### GetBackoffOk

`func (o *V1alpha1RetryStrategy) GetBackoffOk() (*V1alpha1Backoff, bool)`

GetBackoffOk returns a tuple with the Backoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoff

`func (o *V1alpha1RetryStrategy) SetBackoff(v V1alpha1Backoff)`

SetBackoff sets Backoff field to given value.

### HasBackoff

`func (o *V1alpha1RetryStrategy) HasBackoff() bool`

HasBackoff returns a boolean if a field has been set.

### GetLimit

`func (o *V1alpha1RetryStrategy) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V1alpha1RetryStrategy) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V1alpha1RetryStrategy) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *V1alpha1RetryStrategy) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


