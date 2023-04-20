# ApplicationsetApplicationSetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applicationset** | Pointer to [**V1alpha1ApplicationSet**](V1alpha1ApplicationSet.md) |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 

## Methods

### NewApplicationsetApplicationSetResponse

`func NewApplicationsetApplicationSetResponse() *ApplicationsetApplicationSetResponse`

NewApplicationsetApplicationSetResponse instantiates a new ApplicationsetApplicationSetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationsetApplicationSetResponseWithDefaults

`func NewApplicationsetApplicationSetResponseWithDefaults() *ApplicationsetApplicationSetResponse`

NewApplicationsetApplicationSetResponseWithDefaults instantiates a new ApplicationsetApplicationSetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationset

`func (o *ApplicationsetApplicationSetResponse) GetApplicationset() V1alpha1ApplicationSet`

GetApplicationset returns the Applicationset field if non-nil, zero value otherwise.

### GetApplicationsetOk

`func (o *ApplicationsetApplicationSetResponse) GetApplicationsetOk() (*V1alpha1ApplicationSet, bool)`

GetApplicationsetOk returns a tuple with the Applicationset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationset

`func (o *ApplicationsetApplicationSetResponse) SetApplicationset(v V1alpha1ApplicationSet)`

SetApplicationset sets Applicationset field to given value.

### HasApplicationset

`func (o *ApplicationsetApplicationSetResponse) HasApplicationset() bool`

HasApplicationset returns a boolean if a field has been set.

### GetProject

`func (o *ApplicationsetApplicationSetResponse) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ApplicationsetApplicationSetResponse) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ApplicationsetApplicationSetResponse) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ApplicationsetApplicationSetResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


