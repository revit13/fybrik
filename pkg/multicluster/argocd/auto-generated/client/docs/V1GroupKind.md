# V1GroupKind

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 

## Methods

### NewV1GroupKind

`func NewV1GroupKind() *V1GroupKind`

NewV1GroupKind instantiates a new V1GroupKind object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GroupKindWithDefaults

`func NewV1GroupKindWithDefaults() *V1GroupKind`

NewV1GroupKindWithDefaults instantiates a new V1GroupKind object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *V1GroupKind) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *V1GroupKind) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *V1GroupKind) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *V1GroupKind) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetKind

`func (o *V1GroupKind) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1GroupKind) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1GroupKind) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1GroupKind) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


