# V1alpha1RevisionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** |  | [optional] 
**Date** | Pointer to [**V1Time**](V1Time.md) |  | [optional] 
**Message** | Pointer to **string** | Message contains the message associated with the revision, most likely the commit message. | [optional] 
**SignatureInfo** | Pointer to **string** | SignatureInfo contains a hint on the signer if the revision was signed with GPG, and signature verification is enabled. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1alpha1RevisionMetadata

`func NewV1alpha1RevisionMetadata() *V1alpha1RevisionMetadata`

NewV1alpha1RevisionMetadata instantiates a new V1alpha1RevisionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1RevisionMetadataWithDefaults

`func NewV1alpha1RevisionMetadataWithDefaults() *V1alpha1RevisionMetadata`

NewV1alpha1RevisionMetadataWithDefaults instantiates a new V1alpha1RevisionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *V1alpha1RevisionMetadata) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *V1alpha1RevisionMetadata) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *V1alpha1RevisionMetadata) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *V1alpha1RevisionMetadata) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetDate

`func (o *V1alpha1RevisionMetadata) GetDate() V1Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *V1alpha1RevisionMetadata) GetDateOk() (*V1Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *V1alpha1RevisionMetadata) SetDate(v V1Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *V1alpha1RevisionMetadata) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMessage

`func (o *V1alpha1RevisionMetadata) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1alpha1RevisionMetadata) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1alpha1RevisionMetadata) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1alpha1RevisionMetadata) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSignatureInfo

`func (o *V1alpha1RevisionMetadata) GetSignatureInfo() string`

GetSignatureInfo returns the SignatureInfo field if non-nil, zero value otherwise.

### GetSignatureInfoOk

`func (o *V1alpha1RevisionMetadata) GetSignatureInfoOk() (*string, bool)`

GetSignatureInfoOk returns a tuple with the SignatureInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureInfo

`func (o *V1alpha1RevisionMetadata) SetSignatureInfo(v string)`

SetSignatureInfo sets SignatureInfo field to given value.

### HasSignatureInfo

`func (o *V1alpha1RevisionMetadata) HasSignatureInfo() bool`

HasSignatureInfo returns a boolean if a field has been set.

### GetTags

`func (o *V1alpha1RevisionMetadata) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *V1alpha1RevisionMetadata) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *V1alpha1RevisionMetadata) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *V1alpha1RevisionMetadata) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


