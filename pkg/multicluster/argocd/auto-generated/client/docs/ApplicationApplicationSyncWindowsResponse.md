# ApplicationApplicationSyncWindowsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveWindows** | Pointer to [**[]ApplicationApplicationSyncWindow**](ApplicationApplicationSyncWindow.md) |  | [optional] 
**AssignedWindows** | Pointer to [**[]ApplicationApplicationSyncWindow**](ApplicationApplicationSyncWindow.md) |  | [optional] 
**CanSync** | Pointer to **bool** |  | [optional] 

## Methods

### NewApplicationApplicationSyncWindowsResponse

`func NewApplicationApplicationSyncWindowsResponse() *ApplicationApplicationSyncWindowsResponse`

NewApplicationApplicationSyncWindowsResponse instantiates a new ApplicationApplicationSyncWindowsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationApplicationSyncWindowsResponseWithDefaults

`func NewApplicationApplicationSyncWindowsResponseWithDefaults() *ApplicationApplicationSyncWindowsResponse`

NewApplicationApplicationSyncWindowsResponseWithDefaults instantiates a new ApplicationApplicationSyncWindowsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveWindows

`func (o *ApplicationApplicationSyncWindowsResponse) GetActiveWindows() []ApplicationApplicationSyncWindow`

GetActiveWindows returns the ActiveWindows field if non-nil, zero value otherwise.

### GetActiveWindowsOk

`func (o *ApplicationApplicationSyncWindowsResponse) GetActiveWindowsOk() (*[]ApplicationApplicationSyncWindow, bool)`

GetActiveWindowsOk returns a tuple with the ActiveWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveWindows

`func (o *ApplicationApplicationSyncWindowsResponse) SetActiveWindows(v []ApplicationApplicationSyncWindow)`

SetActiveWindows sets ActiveWindows field to given value.

### HasActiveWindows

`func (o *ApplicationApplicationSyncWindowsResponse) HasActiveWindows() bool`

HasActiveWindows returns a boolean if a field has been set.

### GetAssignedWindows

`func (o *ApplicationApplicationSyncWindowsResponse) GetAssignedWindows() []ApplicationApplicationSyncWindow`

GetAssignedWindows returns the AssignedWindows field if non-nil, zero value otherwise.

### GetAssignedWindowsOk

`func (o *ApplicationApplicationSyncWindowsResponse) GetAssignedWindowsOk() (*[]ApplicationApplicationSyncWindow, bool)`

GetAssignedWindowsOk returns a tuple with the AssignedWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedWindows

`func (o *ApplicationApplicationSyncWindowsResponse) SetAssignedWindows(v []ApplicationApplicationSyncWindow)`

SetAssignedWindows sets AssignedWindows field to given value.

### HasAssignedWindows

`func (o *ApplicationApplicationSyncWindowsResponse) HasAssignedWindows() bool`

HasAssignedWindows returns a boolean if a field has been set.

### GetCanSync

`func (o *ApplicationApplicationSyncWindowsResponse) GetCanSync() bool`

GetCanSync returns the CanSync field if non-nil, zero value otherwise.

### GetCanSyncOk

`func (o *ApplicationApplicationSyncWindowsResponse) GetCanSyncOk() (*bool, bool)`

GetCanSyncOk returns a tuple with the CanSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSync

`func (o *ApplicationApplicationSyncWindowsResponse) SetCanSync(v bool)`

SetCanSync sets CanSync field to given value.

### HasCanSync

`func (o *ApplicationApplicationSyncWindowsResponse) HasCanSync() bool`

HasCanSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


