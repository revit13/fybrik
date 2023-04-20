# V1alpha1ApplicationSourceKustomize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonAnnotations** | Pointer to **map[string]string** |  | [optional] 
**CommonLabels** | Pointer to **map[string]string** |  | [optional] 
**ForceCommonAnnotations** | Pointer to **bool** |  | [optional] 
**ForceCommonLabels** | Pointer to **bool** |  | [optional] 
**Images** | Pointer to **[]string** |  | [optional] 
**NamePrefix** | Pointer to **string** |  | [optional] 
**NameSuffix** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ApplicationSourceKustomize

`func NewV1alpha1ApplicationSourceKustomize() *V1alpha1ApplicationSourceKustomize`

NewV1alpha1ApplicationSourceKustomize instantiates a new V1alpha1ApplicationSourceKustomize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationSourceKustomizeWithDefaults

`func NewV1alpha1ApplicationSourceKustomizeWithDefaults() *V1alpha1ApplicationSourceKustomize`

NewV1alpha1ApplicationSourceKustomizeWithDefaults instantiates a new V1alpha1ApplicationSourceKustomize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonAnnotations

`func (o *V1alpha1ApplicationSourceKustomize) GetCommonAnnotations() map[string]string`

GetCommonAnnotations returns the CommonAnnotations field if non-nil, zero value otherwise.

### GetCommonAnnotationsOk

`func (o *V1alpha1ApplicationSourceKustomize) GetCommonAnnotationsOk() (*map[string]string, bool)`

GetCommonAnnotationsOk returns a tuple with the CommonAnnotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonAnnotations

`func (o *V1alpha1ApplicationSourceKustomize) SetCommonAnnotations(v map[string]string)`

SetCommonAnnotations sets CommonAnnotations field to given value.

### HasCommonAnnotations

`func (o *V1alpha1ApplicationSourceKustomize) HasCommonAnnotations() bool`

HasCommonAnnotations returns a boolean if a field has been set.

### GetCommonLabels

`func (o *V1alpha1ApplicationSourceKustomize) GetCommonLabels() map[string]string`

GetCommonLabels returns the CommonLabels field if non-nil, zero value otherwise.

### GetCommonLabelsOk

`func (o *V1alpha1ApplicationSourceKustomize) GetCommonLabelsOk() (*map[string]string, bool)`

GetCommonLabelsOk returns a tuple with the CommonLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonLabels

`func (o *V1alpha1ApplicationSourceKustomize) SetCommonLabels(v map[string]string)`

SetCommonLabels sets CommonLabels field to given value.

### HasCommonLabels

`func (o *V1alpha1ApplicationSourceKustomize) HasCommonLabels() bool`

HasCommonLabels returns a boolean if a field has been set.

### GetForceCommonAnnotations

`func (o *V1alpha1ApplicationSourceKustomize) GetForceCommonAnnotations() bool`

GetForceCommonAnnotations returns the ForceCommonAnnotations field if non-nil, zero value otherwise.

### GetForceCommonAnnotationsOk

`func (o *V1alpha1ApplicationSourceKustomize) GetForceCommonAnnotationsOk() (*bool, bool)`

GetForceCommonAnnotationsOk returns a tuple with the ForceCommonAnnotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceCommonAnnotations

`func (o *V1alpha1ApplicationSourceKustomize) SetForceCommonAnnotations(v bool)`

SetForceCommonAnnotations sets ForceCommonAnnotations field to given value.

### HasForceCommonAnnotations

`func (o *V1alpha1ApplicationSourceKustomize) HasForceCommonAnnotations() bool`

HasForceCommonAnnotations returns a boolean if a field has been set.

### GetForceCommonLabels

`func (o *V1alpha1ApplicationSourceKustomize) GetForceCommonLabels() bool`

GetForceCommonLabels returns the ForceCommonLabels field if non-nil, zero value otherwise.

### GetForceCommonLabelsOk

`func (o *V1alpha1ApplicationSourceKustomize) GetForceCommonLabelsOk() (*bool, bool)`

GetForceCommonLabelsOk returns a tuple with the ForceCommonLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceCommonLabels

`func (o *V1alpha1ApplicationSourceKustomize) SetForceCommonLabels(v bool)`

SetForceCommonLabels sets ForceCommonLabels field to given value.

### HasForceCommonLabels

`func (o *V1alpha1ApplicationSourceKustomize) HasForceCommonLabels() bool`

HasForceCommonLabels returns a boolean if a field has been set.

### GetImages

`func (o *V1alpha1ApplicationSourceKustomize) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *V1alpha1ApplicationSourceKustomize) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *V1alpha1ApplicationSourceKustomize) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *V1alpha1ApplicationSourceKustomize) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetNamePrefix

`func (o *V1alpha1ApplicationSourceKustomize) GetNamePrefix() string`

GetNamePrefix returns the NamePrefix field if non-nil, zero value otherwise.

### GetNamePrefixOk

`func (o *V1alpha1ApplicationSourceKustomize) GetNamePrefixOk() (*string, bool)`

GetNamePrefixOk returns a tuple with the NamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePrefix

`func (o *V1alpha1ApplicationSourceKustomize) SetNamePrefix(v string)`

SetNamePrefix sets NamePrefix field to given value.

### HasNamePrefix

`func (o *V1alpha1ApplicationSourceKustomize) HasNamePrefix() bool`

HasNamePrefix returns a boolean if a field has been set.

### GetNameSuffix

`func (o *V1alpha1ApplicationSourceKustomize) GetNameSuffix() string`

GetNameSuffix returns the NameSuffix field if non-nil, zero value otherwise.

### GetNameSuffixOk

`func (o *V1alpha1ApplicationSourceKustomize) GetNameSuffixOk() (*string, bool)`

GetNameSuffixOk returns a tuple with the NameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSuffix

`func (o *V1alpha1ApplicationSourceKustomize) SetNameSuffix(v string)`

SetNameSuffix sets NameSuffix field to given value.

### HasNameSuffix

`func (o *V1alpha1ApplicationSourceKustomize) HasNameSuffix() bool`

HasNameSuffix returns a boolean if a field has been set.

### GetVersion

`func (o *V1alpha1ApplicationSourceKustomize) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V1alpha1ApplicationSourceKustomize) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V1alpha1ApplicationSourceKustomize) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V1alpha1ApplicationSourceKustomize) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


