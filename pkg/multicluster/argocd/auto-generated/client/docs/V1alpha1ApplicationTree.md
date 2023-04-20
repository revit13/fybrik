# V1alpha1ApplicationTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to [**[]V1alpha1HostInfo**](V1alpha1HostInfo.md) |  | [optional] 
**Nodes** | Pointer to [**[]V1alpha1ResourceNode**](V1alpha1ResourceNode.md) | Nodes contains list of nodes which either directly managed by the application and children of directly managed nodes. | [optional] 
**OrphanedNodes** | Pointer to [**[]V1alpha1ResourceNode**](V1alpha1ResourceNode.md) | OrphanedNodes contains if or orphaned nodes: nodes which are not managed by the app but in the same namespace. List is populated only if orphaned resources enabled in app project. | [optional] 

## Methods

### NewV1alpha1ApplicationTree

`func NewV1alpha1ApplicationTree() *V1alpha1ApplicationTree`

NewV1alpha1ApplicationTree instantiates a new V1alpha1ApplicationTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationTreeWithDefaults

`func NewV1alpha1ApplicationTreeWithDefaults() *V1alpha1ApplicationTree`

NewV1alpha1ApplicationTreeWithDefaults instantiates a new V1alpha1ApplicationTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *V1alpha1ApplicationTree) GetHosts() []V1alpha1HostInfo`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *V1alpha1ApplicationTree) GetHostsOk() (*[]V1alpha1HostInfo, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *V1alpha1ApplicationTree) SetHosts(v []V1alpha1HostInfo)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *V1alpha1ApplicationTree) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetNodes

`func (o *V1alpha1ApplicationTree) GetNodes() []V1alpha1ResourceNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V1alpha1ApplicationTree) GetNodesOk() (*[]V1alpha1ResourceNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V1alpha1ApplicationTree) SetNodes(v []V1alpha1ResourceNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V1alpha1ApplicationTree) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetOrphanedNodes

`func (o *V1alpha1ApplicationTree) GetOrphanedNodes() []V1alpha1ResourceNode`

GetOrphanedNodes returns the OrphanedNodes field if non-nil, zero value otherwise.

### GetOrphanedNodesOk

`func (o *V1alpha1ApplicationTree) GetOrphanedNodesOk() (*[]V1alpha1ResourceNode, bool)`

GetOrphanedNodesOk returns a tuple with the OrphanedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphanedNodes

`func (o *V1alpha1ApplicationTree) SetOrphanedNodes(v []V1alpha1ResourceNode)`

SetOrphanedNodes sets OrphanedNodes field to given value.

### HasOrphanedNodes

`func (o *V1alpha1ApplicationTree) HasOrphanedNodes() bool`

HasOrphanedNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


