# ApplicationResourceActionsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]V1alpha1ResourceAction**](V1alpha1ResourceAction.md) |  | [optional] 

## Methods

### NewApplicationResourceActionsListResponse

`func NewApplicationResourceActionsListResponse() *ApplicationResourceActionsListResponse`

NewApplicationResourceActionsListResponse instantiates a new ApplicationResourceActionsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResourceActionsListResponseWithDefaults

`func NewApplicationResourceActionsListResponseWithDefaults() *ApplicationResourceActionsListResponse`

NewApplicationResourceActionsListResponseWithDefaults instantiates a new ApplicationResourceActionsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ApplicationResourceActionsListResponse) GetActions() []V1alpha1ResourceAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApplicationResourceActionsListResponse) GetActionsOk() (*[]V1alpha1ResourceAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApplicationResourceActionsListResponse) SetActions(v []V1alpha1ResourceAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ApplicationResourceActionsListResponse) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


