# V1alpha1ApplicationMatchExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1alpha1ApplicationMatchExpression

`func NewV1alpha1ApplicationMatchExpression() *V1alpha1ApplicationMatchExpression`

NewV1alpha1ApplicationMatchExpression instantiates a new V1alpha1ApplicationMatchExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ApplicationMatchExpressionWithDefaults

`func NewV1alpha1ApplicationMatchExpressionWithDefaults() *V1alpha1ApplicationMatchExpression`

NewV1alpha1ApplicationMatchExpressionWithDefaults instantiates a new V1alpha1ApplicationMatchExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *V1alpha1ApplicationMatchExpression) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V1alpha1ApplicationMatchExpression) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V1alpha1ApplicationMatchExpression) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *V1alpha1ApplicationMatchExpression) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOperator

`func (o *V1alpha1ApplicationMatchExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *V1alpha1ApplicationMatchExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *V1alpha1ApplicationMatchExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *V1alpha1ApplicationMatchExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *V1alpha1ApplicationMatchExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *V1alpha1ApplicationMatchExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *V1alpha1ApplicationMatchExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *V1alpha1ApplicationMatchExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


