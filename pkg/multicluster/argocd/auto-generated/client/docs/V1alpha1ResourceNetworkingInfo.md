# V1alpha1ResourceNetworkingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalURLs** | Pointer to **[]string** | ExternalURLs holds list of URLs which should be available externally. List is populated for ingress resources using rules hostnames. | [optional] 
**Ingress** | Pointer to [**[]V1LoadBalancerIngress**](V1LoadBalancerIngress.md) |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**TargetLabels** | Pointer to **map[string]string** |  | [optional] 
**TargetRefs** | Pointer to [**[]V1alpha1ResourceRef**](V1alpha1ResourceRef.md) |  | [optional] 

## Methods

### NewV1alpha1ResourceNetworkingInfo

`func NewV1alpha1ResourceNetworkingInfo() *V1alpha1ResourceNetworkingInfo`

NewV1alpha1ResourceNetworkingInfo instantiates a new V1alpha1ResourceNetworkingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ResourceNetworkingInfoWithDefaults

`func NewV1alpha1ResourceNetworkingInfoWithDefaults() *V1alpha1ResourceNetworkingInfo`

NewV1alpha1ResourceNetworkingInfoWithDefaults instantiates a new V1alpha1ResourceNetworkingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalURLs

`func (o *V1alpha1ResourceNetworkingInfo) GetExternalURLs() []string`

GetExternalURLs returns the ExternalURLs field if non-nil, zero value otherwise.

### GetExternalURLsOk

`func (o *V1alpha1ResourceNetworkingInfo) GetExternalURLsOk() (*[]string, bool)`

GetExternalURLsOk returns a tuple with the ExternalURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalURLs

`func (o *V1alpha1ResourceNetworkingInfo) SetExternalURLs(v []string)`

SetExternalURLs sets ExternalURLs field to given value.

### HasExternalURLs

`func (o *V1alpha1ResourceNetworkingInfo) HasExternalURLs() bool`

HasExternalURLs returns a boolean if a field has been set.

### GetIngress

`func (o *V1alpha1ResourceNetworkingInfo) GetIngress() []V1LoadBalancerIngress`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *V1alpha1ResourceNetworkingInfo) GetIngressOk() (*[]V1LoadBalancerIngress, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *V1alpha1ResourceNetworkingInfo) SetIngress(v []V1LoadBalancerIngress)`

SetIngress sets Ingress field to given value.

### HasIngress

`func (o *V1alpha1ResourceNetworkingInfo) HasIngress() bool`

HasIngress returns a boolean if a field has been set.

### GetLabels

`func (o *V1alpha1ResourceNetworkingInfo) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1alpha1ResourceNetworkingInfo) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1alpha1ResourceNetworkingInfo) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1alpha1ResourceNetworkingInfo) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTargetLabels

`func (o *V1alpha1ResourceNetworkingInfo) GetTargetLabels() map[string]string`

GetTargetLabels returns the TargetLabels field if non-nil, zero value otherwise.

### GetTargetLabelsOk

`func (o *V1alpha1ResourceNetworkingInfo) GetTargetLabelsOk() (*map[string]string, bool)`

GetTargetLabelsOk returns a tuple with the TargetLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLabels

`func (o *V1alpha1ResourceNetworkingInfo) SetTargetLabels(v map[string]string)`

SetTargetLabels sets TargetLabels field to given value.

### HasTargetLabels

`func (o *V1alpha1ResourceNetworkingInfo) HasTargetLabels() bool`

HasTargetLabels returns a boolean if a field has been set.

### GetTargetRefs

`func (o *V1alpha1ResourceNetworkingInfo) GetTargetRefs() []V1alpha1ResourceRef`

GetTargetRefs returns the TargetRefs field if non-nil, zero value otherwise.

### GetTargetRefsOk

`func (o *V1alpha1ResourceNetworkingInfo) GetTargetRefsOk() (*[]V1alpha1ResourceRef, bool)`

GetTargetRefsOk returns a tuple with the TargetRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRefs

`func (o *V1alpha1ResourceNetworkingInfo) SetTargetRefs(v []V1alpha1ResourceRef)`

SetTargetRefs sets TargetRefs field to given value.

### HasTargetRefs

`func (o *V1alpha1ResourceNetworkingInfo) HasTargetRefs() bool`

HasTargetRefs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


