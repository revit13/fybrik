# RuntimeStreamError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]ProtobufAny**](ProtobufAny.md) |  | [optional] 
**GrpcCode** | Pointer to **int32** |  | [optional] 
**HttpCode** | Pointer to **int32** |  | [optional] 
**HttpStatus** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewRuntimeStreamError

`func NewRuntimeStreamError() *RuntimeStreamError`

NewRuntimeStreamError instantiates a new RuntimeStreamError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeStreamErrorWithDefaults

`func NewRuntimeStreamErrorWithDefaults() *RuntimeStreamError`

NewRuntimeStreamErrorWithDefaults instantiates a new RuntimeStreamError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *RuntimeStreamError) GetDetails() []ProtobufAny`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RuntimeStreamError) GetDetailsOk() (*[]ProtobufAny, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RuntimeStreamError) SetDetails(v []ProtobufAny)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RuntimeStreamError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetGrpcCode

`func (o *RuntimeStreamError) GetGrpcCode() int32`

GetGrpcCode returns the GrpcCode field if non-nil, zero value otherwise.

### GetGrpcCodeOk

`func (o *RuntimeStreamError) GetGrpcCodeOk() (*int32, bool)`

GetGrpcCodeOk returns a tuple with the GrpcCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpcCode

`func (o *RuntimeStreamError) SetGrpcCode(v int32)`

SetGrpcCode sets GrpcCode field to given value.

### HasGrpcCode

`func (o *RuntimeStreamError) HasGrpcCode() bool`

HasGrpcCode returns a boolean if a field has been set.

### GetHttpCode

`func (o *RuntimeStreamError) GetHttpCode() int32`

GetHttpCode returns the HttpCode field if non-nil, zero value otherwise.

### GetHttpCodeOk

`func (o *RuntimeStreamError) GetHttpCodeOk() (*int32, bool)`

GetHttpCodeOk returns a tuple with the HttpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCode

`func (o *RuntimeStreamError) SetHttpCode(v int32)`

SetHttpCode sets HttpCode field to given value.

### HasHttpCode

`func (o *RuntimeStreamError) HasHttpCode() bool`

HasHttpCode returns a boolean if a field has been set.

### GetHttpStatus

`func (o *RuntimeStreamError) GetHttpStatus() string`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *RuntimeStreamError) GetHttpStatusOk() (*string, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *RuntimeStreamError) SetHttpStatus(v string)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *RuntimeStreamError) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetMessage

`func (o *RuntimeStreamError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RuntimeStreamError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RuntimeStreamError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RuntimeStreamError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


