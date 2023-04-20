# V1alpha1Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Operation** | Pointer to [**V1alpha1Operation**](V1alpha1Operation.md) |  | [optional] 
**Spec** | Pointer to [**V1alpha1ApplicationSpec**](V1alpha1ApplicationSpec.md) |  | [optional] 
**Status** | Pointer to [**V1alpha1ApplicationStatus**](V1alpha1ApplicationStatus.md) |  | [optional] 

## Methods

### NewV1alpha1Application

`func NewV1alpha1Application() *V1alpha1Application`

NewV1alpha1Application instantiates a new V1alpha1Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationWithDefaults

`func NewV1alpha1ApplicationWithDefaults() *V1alpha1Application`

NewV1alpha1ApplicationWithDefaults instantiates a new V1alpha1Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *V1alpha1Application) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1alpha1Application) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1alpha1Application) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1alpha1Application) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOperation

`func (o *V1alpha1Application) GetOperation() V1alpha1Operation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *V1alpha1Application) GetOperationOk() (*V1alpha1Operation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *V1alpha1Application) SetOperation(v V1alpha1Operation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *V1alpha1Application) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetSpec

`func (o *V1alpha1Application) GetSpec() V1alpha1ApplicationSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1alpha1Application) GetSpecOk() (*V1alpha1ApplicationSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1alpha1Application) SetSpec(v V1alpha1ApplicationSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *V1alpha1Application) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *V1alpha1Application) GetStatus() V1alpha1ApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1alpha1Application) GetStatusOk() (*V1alpha1ApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1alpha1Application) SetStatus(v V1alpha1ApplicationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1alpha1Application) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


