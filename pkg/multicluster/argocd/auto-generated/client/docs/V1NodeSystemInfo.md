# V1NodeSystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** |  | [optional] 
**BootID** | Pointer to **string** | Boot ID reported by the node. | [optional] 
**ContainerRuntimeVersion** | Pointer to **string** | ContainerRuntime Version reported by the node through runtime remote API (e.g. containerd://1.4.2). | [optional] 
**KernelVersion** | Pointer to **string** | Kernel Version reported by the node from &#39;uname -r&#39; (e.g. 3.16.0-0.bpo.4-amd64). | [optional] 
**KubeProxyVersion** | Pointer to **string** | KubeProxy Version reported by the node. | [optional] 
**KubeletVersion** | Pointer to **string** | Kubelet Version reported by the node. | [optional] 
**MachineID** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**OsImage** | Pointer to **string** | OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)). | [optional] 
**SystemUUID** | Pointer to **string** |  | [optional] 

## Methods

### NewV1NodeSystemInfo

`func NewV1NodeSystemInfo() *V1NodeSystemInfo`

NewV1NodeSystemInfo instantiates a new V1NodeSystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NodeSystemInfoWithDefaults

`func NewV1NodeSystemInfoWithDefaults() *V1NodeSystemInfo`

NewV1NodeSystemInfoWithDefaults instantiates a new V1NodeSystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *V1NodeSystemInfo) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *V1NodeSystemInfo) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *V1NodeSystemInfo) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *V1NodeSystemInfo) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetBootID

`func (o *V1NodeSystemInfo) GetBootID() string`

GetBootID returns the BootID field if non-nil, zero value otherwise.

### GetBootIDOk

`func (o *V1NodeSystemInfo) GetBootIDOk() (*string, bool)`

GetBootIDOk returns a tuple with the BootID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootID

`func (o *V1NodeSystemInfo) SetBootID(v string)`

SetBootID sets BootID field to given value.

### HasBootID

`func (o *V1NodeSystemInfo) HasBootID() bool`

HasBootID returns a boolean if a field has been set.

### GetContainerRuntimeVersion

`func (o *V1NodeSystemInfo) GetContainerRuntimeVersion() string`

GetContainerRuntimeVersion returns the ContainerRuntimeVersion field if non-nil, zero value otherwise.

### GetContainerRuntimeVersionOk

`func (o *V1NodeSystemInfo) GetContainerRuntimeVersionOk() (*string, bool)`

GetContainerRuntimeVersionOk returns a tuple with the ContainerRuntimeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerRuntimeVersion

`func (o *V1NodeSystemInfo) SetContainerRuntimeVersion(v string)`

SetContainerRuntimeVersion sets ContainerRuntimeVersion field to given value.

### HasContainerRuntimeVersion

`func (o *V1NodeSystemInfo) HasContainerRuntimeVersion() bool`

HasContainerRuntimeVersion returns a boolean if a field has been set.

### GetKernelVersion

`func (o *V1NodeSystemInfo) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *V1NodeSystemInfo) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *V1NodeSystemInfo) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.

### HasKernelVersion

`func (o *V1NodeSystemInfo) HasKernelVersion() bool`

HasKernelVersion returns a boolean if a field has been set.

### GetKubeProxyVersion

`func (o *V1NodeSystemInfo) GetKubeProxyVersion() string`

GetKubeProxyVersion returns the KubeProxyVersion field if non-nil, zero value otherwise.

### GetKubeProxyVersionOk

`func (o *V1NodeSystemInfo) GetKubeProxyVersionOk() (*string, bool)`

GetKubeProxyVersionOk returns a tuple with the KubeProxyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeProxyVersion

`func (o *V1NodeSystemInfo) SetKubeProxyVersion(v string)`

SetKubeProxyVersion sets KubeProxyVersion field to given value.

### HasKubeProxyVersion

`func (o *V1NodeSystemInfo) HasKubeProxyVersion() bool`

HasKubeProxyVersion returns a boolean if a field has been set.

### GetKubeletVersion

`func (o *V1NodeSystemInfo) GetKubeletVersion() string`

GetKubeletVersion returns the KubeletVersion field if non-nil, zero value otherwise.

### GetKubeletVersionOk

`func (o *V1NodeSystemInfo) GetKubeletVersionOk() (*string, bool)`

GetKubeletVersionOk returns a tuple with the KubeletVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletVersion

`func (o *V1NodeSystemInfo) SetKubeletVersion(v string)`

SetKubeletVersion sets KubeletVersion field to given value.

### HasKubeletVersion

`func (o *V1NodeSystemInfo) HasKubeletVersion() bool`

HasKubeletVersion returns a boolean if a field has been set.

### GetMachineID

`func (o *V1NodeSystemInfo) GetMachineID() string`

GetMachineID returns the MachineID field if non-nil, zero value otherwise.

### GetMachineIDOk

`func (o *V1NodeSystemInfo) GetMachineIDOk() (*string, bool)`

GetMachineIDOk returns a tuple with the MachineID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineID

`func (o *V1NodeSystemInfo) SetMachineID(v string)`

SetMachineID sets MachineID field to given value.

### HasMachineID

`func (o *V1NodeSystemInfo) HasMachineID() bool`

HasMachineID returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *V1NodeSystemInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *V1NodeSystemInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *V1NodeSystemInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *V1NodeSystemInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOsImage

`func (o *V1NodeSystemInfo) GetOsImage() string`

GetOsImage returns the OsImage field if non-nil, zero value otherwise.

### GetOsImageOk

`func (o *V1NodeSystemInfo) GetOsImageOk() (*string, bool)`

GetOsImageOk returns a tuple with the OsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsImage

`func (o *V1NodeSystemInfo) SetOsImage(v string)`

SetOsImage sets OsImage field to given value.

### HasOsImage

`func (o *V1NodeSystemInfo) HasOsImage() bool`

HasOsImage returns a boolean if a field has been set.

### GetSystemUUID

`func (o *V1NodeSystemInfo) GetSystemUUID() string`

GetSystemUUID returns the SystemUUID field if non-nil, zero value otherwise.

### GetSystemUUIDOk

`func (o *V1NodeSystemInfo) GetSystemUUIDOk() (*string, bool)`

GetSystemUUIDOk returns a tuple with the SystemUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUUID

`func (o *V1NodeSystemInfo) SetSystemUUID(v string)`

SetSystemUUID sets SystemUUID field to given value.

### HasSystemUUID

`func (o *V1NodeSystemInfo) HasSystemUUID() bool`

HasSystemUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


